package test

import (
	"GeeCache/src"
	"fmt"
	"log"
	"net/http"
	"testing"
)

//type server string
//
//func (h *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	log.Println(r.URL.Path)
//	w.Write([]byte("Hello World!"))
//}

func TestHttp(t *testing.T) {
	var db = map[string]string{
		"Tom":  "630",
		"Jack": "589",
		"Sam":  "dfaf",
	}
	src.NewGroup("scores", 2<<10, src.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))

	addr := "localhost:9999"
	peers := src.NewHTTPPool(addr)
	log.Println("geecache is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}
