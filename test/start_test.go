package test

import (
	"GeeCache/src/conf"
	"strconv"
	"sync"
	"testing"
)

func TestStart(t *testing.T) {
	confdata := conf.NewConfigData("../src/conf/config.yaml")
	servers := confdata.GetOnlineServers()
	var wg sync.WaitGroup

	for _, v := range servers {
		// 捕获循环变量的局部副本
		server := v

		go func() {
			defer wg.Done()
			println(server.IP + strconv.Itoa(server.Port))
			wg.Add(1)
			// 执行您的测试逻辑
		}()
	}

	wg.Wait()
}
