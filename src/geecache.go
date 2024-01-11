package src

import (
	pb "GeeCache/geecachepb"
	"GeeCache/src/singleFlight"
	"fmt"
	"log"
	"sync"
)

var (
	mu sync.RWMutex
	// Group 和 Group name 对应
	groups = make(map[string]*Group)
)

// A Getter loads data for a key.
type Getter interface {
	Get(key string) ([]byte, error)
}

// A GetterFunc implements Getter with a function.
type GetterFunc func(key string) ([]byte, error)

// Get implements Getter interface function
func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

type Group struct {
	name      string
	getter    Getter
	mainCache cache
	peers     PeerPicker
	// 确保短时间内相同的key向远端请求仅一次
	loader *singleFlight.Group
}

func NewGroup(name string, cacheBytes int64, getter Getter) *Group {
	if getter == nil {
		panic("nil Getter")
	}
	mu.Lock()
	defer mu.Unlock()
	g := &Group{
		name:      name,
		getter:    getter,
		mainCache: cache{cacheBytes: cacheBytes},
		loader:    &singleFlight.Group{},
	}
	groups[name] = g
	return g
}

// RegisterPeers 将 实现了 PeerPicker 接口的 HTTPPool 注入到 Group 中
func (g *Group) RegisterPeers(peers PeerPicker) {
	if g.peers != nil {
		panic("RegisterPeerPicker called more than once")
	}
	g.peers = peers
}

func GetGroup(name string) *Group {
	mu.RLock()
	g := groups[name]
	mu.RUnlock()
	return g
}

// Get value for a key from cache
func (g *Group) Get(key string) (ByteView, error) {
	if key == "" {
		return ByteView{}, fmt.Errorf("key is required")
	}
	if v, ok := g.mainCache.get(key); ok {
		log.Println("[LocalCache] hit")
		return v, nil
	}
	// 不在本地的cache中执行加载
	return g.load(key)
}

func (g *Group) load(key string) (value ByteView, err error) {
	// 确保短时间内相同的key向远端请求仅一次
	viewi, err := g.loader.Do(key, func() (interface{}, error) {
		// 向远程节点请求
		if g.peers != nil {
			if peer, ok := g.peers.PickPeer(key); ok {
				value, err = g.getFromPeer(peer, key)
				if err == nil {
					return value, nil
				}
				log.Println("[RemoteCache] Failed to get from peer", err)
			}
		}
		// 本机节点或者远程节点获取失败
		return g.getLocally(key)
	})
	// Do 执行没有出错，返回
	if err == nil {
		return viewi.(ByteView), nil
	}
	return

}

func (g *Group) getFromPeer(peer PeerGetter, key string) (ByteView, error) {
	req := &pb.Request{
		Group: g.name,
		Key:   key,
	}
	res := &pb.Response{}
	err := peer.Get(req, res)
	if err != nil {
		return ByteView{}, err
	}
	return ByteView{b: res.Value}, nil
}

func (g *Group) getLocally(key string) (ByteView, error) {
	// 在数据中查看数据字段是否存在 byte 为值的字节码
	bytes, err := g.getter.Get(key)
	if err != nil {
		return ByteView{}, err
	}
	value := ByteView{cloneBytes(bytes)}
	g.populateCache(key, value)
	return value, nil
}

func (g *Group) populateCache(key string, value ByteView) {
	g.mainCache.add(key, value)
}
