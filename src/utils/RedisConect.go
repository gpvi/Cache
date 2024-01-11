package utils

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient() *RedisClient {
	var cli = NewClient()
	client := &RedisClient{
		client: cli,
	}
	return client
}

func NewClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // redis地址
		Password: "",               // 密码
		DB:       0,                // 使用默认数据库
	})
	return client
}

func (*RedisClient) GetKey(key string) (string, bool) {
	//var cli = NewClient()
	client := NewClient()
	defer client.Close()

	val, err := client.Get(key).Result()
	if err != nil {
		fmt.Errorf("key not exist in DB \n you can use 'set key' before get ")
		return "", false
	}
	log.Println(val)
	log.Println("------------------")
	return val, true
}

func (*RedisClient) Setkey(key string, value string) {
	cli := NewClient()
	err := cli.Set(key, value, 0).Err()
	if err != nil {
		fmt.Errorf("SET ERROR")
	}
}
