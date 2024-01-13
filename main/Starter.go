package main

import (
	"GeeCache/src"
	"fmt"
	"log"
	"net/http"
)

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

func startAPIServer(apiAddr string, gee *src.Group, mux *http.ServeMux) {
	mux.Handle("/api", http.HandlerFunc(
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
	//log.Println(apiAddr[7:])
	log.Fatal(http.ListenAndServe(apiAddr[7:], mux))
}

//func startAPIServer(apiAddr string, gee *src.Group) {
//	http.Handle("/api", http.HandlerFunc(
//		func(w http.ResponseWriter, r *http.Request) {
//			key := r.URL.Query().Get("key")
//			view, err := gee.Get(key)
//			if err != nil {
//				http.Error(w, err.Error(), http.StatusInternalServerError)
//				return
//			}
//			w.Header().Set("Content-Type", "application/octet-stream")
//			w.Write(view.ByteSlice())
//		}))
//	log.Println("fontend server is running at", apiAddr)
//	log.Fatal(http.ListenAndServe(apiAddr[7:], nil))
//}
