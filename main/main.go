package main

import (
	"GeeCache/src/conf"
	"log"
	"net/http"
	"sync"
)

func init() {
	Init_before()
}
func main() {
	//var port int
	//var api bool
	var wg sync.WaitGroup
	//log.Println(Addrs)
	for _, server := range Online_servers {
		wg.Add(1)
		go func(server conf.Server, muxCopy *http.ServeMux) {
			defer wg.Done()
			//println(AddrMap[server.Port])
			gee := createGroup()
			if server.API {
				log.Println("start API")
				go startAPIServer(ApiAddr, gee, muxCopy)
			}
			startCacheServer(AddrMap[server.Port], Addrs, gee)

		}(server, http.NewServeMux())
	}
	wg.Wait()
	//flag.IntVar(&port, "port", 8001, "Geecache server port")
	//flag.BoolVar(&api, "api", false, "Start a api server?")
	//flag.Parse()
	//gee := createGroup()
	//if api {
	//	log.Println("start API")
	//	go startAPIServer(ApiAddr, gee)
	//}
	//startCacheServer(AddrMap[port], []string(Addrs), gee)
}
