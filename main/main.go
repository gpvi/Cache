package main

import (
	"GeeCache/src"
	"flag"
	"fmt"
	"log"
	"net/http"
)

// var Online_servers []conf.Server
// var Config_path = "src/conf/config.yaml"
// var ApiAddr string
// var AddrMap = make(map[int]string)
// var Addrs []string
//
//	var DB = map[string]string{
//		"Tom":  "630",
//		"Jack": "589",
//		"Sam":  "567",
//	}
func init() {
	// 配置front-server
	//confdata := conf.NewConfigData(Config_path)
	//ApiAddr = confdata.GetFrontServer()
	//Online_servers = confdata.GetOnlineServers()
	//log.Println("ApiAddr:", ApiAddr)
	//// 配置 peers
	//// test
	//var peers []string
	////var as []string
	//for _, v := range confdata.GetOnlineServers() {
	//	str := "http://" + v.IP + ":" + strconv.Itoa(v.Port)
	//	println(str)
	//	err := append(peers, str)
	//	//println("http://" + v.IP + ":" + strconv.Itoa(v.Port))
	//	if err != nil {
	//		fmt.Errorf("ip init append error")
	//		//panic("append error")
	//	}
	//	AddrMap[v.Port] = str
	//	//append(as, )
	//}
	//Addrs = peers
	Init_before()
}

// 创建组
func createGroup() *src.Group {
	return src.NewGroup("scores", 2<<10, src.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := DB[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))
}

// addr 本地 ，Addrs peers启动缓存服务器：创建 HTTPPool，添加节点信息，注册到 group中，启动 HTTP 服务
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

func main() {
	var port int
	var api bool
	flag.IntVar(&port, "port", 8001, "Geecache server port")
	flag.BoolVar(&api, "api", false, "Start a api server?")
	flag.Parse()
	gee := createGroup()
	if api {
		log.Println("start API")
		go startAPIServer(ApiAddr, gee)
	}
	startCacheServer(AddrMap[port], []string(Addrs), gee)
}
