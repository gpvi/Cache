package utils

import "testing"

func TestRedis(t *testing.T) {
	cli := NewRedisClient()
	cli.Setkey("tom", "12")
	val, ok := cli.GetKey("tom")
	if ok {
		println(val)
	}
}
