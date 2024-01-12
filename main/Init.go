package main

import (
	"GeeCache/src/conf"
	"fmt"
	"log"
	"strconv"
)

var Online_servers []conf.Server
var Config_path = "src/conf/config.yaml"
var ApiAddr string
var AddrMap = make(map[int]string)
var Addrs []string
var DB = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func Init_before() {
	//配置front-server
	confdata := conf.NewConfigData(Config_path)
	ApiAddr = confdata.GetFrontServer()
	Online_servers = confdata.GetOnlineServers()
	log.Println("ApiAddr:", ApiAddr)
	// 配置 peers
	// test
	var peers []string
	//var as []string
	for _, v := range confdata.GetOnlineServers() {
		str := "http://" + v.IP + ":" + strconv.Itoa(v.Port)
		println(str)
		err := append(peers, str)
		//println("http://" + v.IP + ":" + strconv.Itoa(v.Port))
		if err != nil {
			fmt.Errorf("ip init append error")
			//panic("append error")
		}
		AddrMap[v.Port] = str
		//append(as, )
	}
	Addrs = peers

}
