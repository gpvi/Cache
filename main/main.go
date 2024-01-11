package main

import (
	"GeeCache/src"
	"GeeCache/src/conf"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

var config_path = "src/conf/config.yaml"
var apiAddr string
var addrMap = make(map[int]string)
var addrs []string
var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

// 创建组
func createGroup() *src.Group {
	return src.NewGroup("scores", 2<<10, src.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))
}

// addr 本地 ，addrs peers启动缓存服务器：创建 HTTPPool，添加节点信息，注册到 group中，启动 HTTP 服务
func startCacheServer(addr string, addrs []string, gee *src.Group) {
	peers := src.NewHTTPPool(addr)
	peers.Set(addrs...)
	gee.RegisterPeers(peers)
	log.Println("geecache is running at", addr)
	log.Fatal(http.ListenAndServe(addr[7:], peers))
}

// 用来启动一个 API 服务（端口 9999），与用户进行交互，用户感知。
// 主要是为了获取 key
func startAPIServer(apiAddr string, gee *src.Group) {
	http.Handle("/api", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			key := r.URL.Query().Get("key")
			view, err := gee.Get(key)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/octet-stream")
			w.Write(view.ByteSlice())
		}))
	log.Println("fontend server is running at", apiAddr)
	log.Fatal(http.ListenAndServe(apiAddr[7:], nil))
}
func init() {
	// 配置front-server
	confdata := conf.NewConfigData(config_path)
	apiAddr = confdata.GetFrontServer()
	log.Println(apiAddr)
	// 配置 peers
	peers := confdata.GetPeers()
	for _, peer := range peers {
		parsedURL, err := url.Parse(peer)
		if err != nil {
			fmt.Println("Error parsing URL:", err)
			return
		}

		// 获取端口值
		port := parsedURL.Port()
		if port == "" {
			// 如果URL中没有指定端口，则使用默认端口
			port = "80"
		}
		portnum, err := strconv.Atoi(port)
		if err != nil {
			fmt.Println("转换失败:", err)
			return
		}
		addrMap[portnum] = peer
	}
	for _, v := range addrMap {
		addrs = append(addrs, v)
	}

	for key, val := range addrMap {
		log.Println(key, "-->", val)
	}

}
func main() {
	var port int
	var api bool

	flag.IntVar(&port, "port", 8001, "Geecache server port")
	flag.BoolVar(&api, "api", false, "Start a api server?")
	flag.Parse()
	gee := createGroup()
	if api {
		log.Println("start API")
		go startAPIServer(apiAddr, gee)
	}
	startCacheServer(addrMap[port], []string(addrs), gee)
}
