package main

import (
	"mygolearn/myredis/myredisutils"
	"time"
)

var (
	dialTimeout    = 5 * time.Second
	requestTimeout = 2 * time.Second
	endpoints      = []string{"192.168.1.105:2379"} // 测试服
	REDIS_URI      string
	REDIS_PWD      string
)

func init() {
	// boltdb文件读取配置项
	cli, err := cache.NewClientv3Tls(endpoints, 5*time.Second)
	if err != nil {
		panic("etcd配置项读取出错" + err.Error())
	}
	defer cache.CloseClientv3(cli)
	keyPrefix := ""
	REDIS_URI, _ := cache.EtcdGet(cli, keyPrefix+"redis_url")
	REDIS_PWD, _ := cache.EtcdGet(cli, keyPrefix+"redis_pwd")
	cache.InitRedis(REDIS_URI, REDIS_PWD)
}

func main() {
	cache.InitRedis(REDIS_URI, REDIS_PWD)
}
