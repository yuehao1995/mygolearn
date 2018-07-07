package main

import (
	"mygolearn/myredis/myredisutils"
)

func init() {
	keyPrefix := ""
	REDIS_URI, _ := cache.EtcdGet(cli, keyPrefix+"redis_url")
	REDIS_PWD, _ := cache.EtcdGet(cli, keyPrefix+"redis_pwd")
	cache.InitRedis(REDIS_URI, REDIS_PWD)
}

func main() {

}
