package test

import (
	"GeeCache/src/conf"
	"testing"
)

//	func TestConfig(t *testing.T) {
//		confdata := conf.NewConfigData("../src/conf/config.yaml")
//		for _, v := range confdata.GetOnlineServers() {
//			str := "http://" + v.IP + ":"
//			println(str)
//		}
//	}
func TestConfig(t *testing.T) {
	confdata := conf.NewConfigData("../src/conf/config.yaml")
	println(confdata.GetFrontServer())
}
