package test

import (
	"GeeCache/src/conf"
	"strconv"
	"testing"
)

func TestConfig(t *testing.T) {
	confdata := conf.NewConfigData("../src/conf/config.yaml")
	//apiAddr := confdata.GetFrontServer()
	//// 配置 peers
	//peers := confdata.GetPeers()
	//var addrMap = make(map[int]string)
	//var addrs []string
	//for _, peer := range peers {
	//	parsedURL, err := url.Parse(peer)
	//	if err != nil {
	//		fmt.Println("Error parsing URL:", err)
	//		return
	//	}
	//
	//	// 获取端口值
	//	port := parsedURL.Port()
	//	if port == "" {
	//		// 如果URL中没有指定端口，则使用默认端口
	//		port = "80"
	//	}
	//	portnum, err := strconv.Atoi(port)
	//	if err != nil {
	//		fmt.Println("转换失败:", err)
	//		return
	//	}
	//	addrMap[portnum] = peer
	//}
	//println("addrs: ")
	//for _, v := range addrMap {
	//	addrs = append(addrs, v)
	//	print(v, " ")
	//}
	//print("\n")
	//
	//println("apiaddr: ", apiAddr)
	//
	//for k, v := range addrMap {
	//	println(k, "-->", v)
	//}
	for _, v := range confdata.GetOnlineServers() {
		println("http://" + v.IP + ":" + strconv.Itoa(v.Port))
	}

}
