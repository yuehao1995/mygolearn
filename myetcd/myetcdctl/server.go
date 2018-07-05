package main

import (
	"context"
	"time"
)

var (
	dialTimeout    = 5 * time.Second
	requestTimeout = 2 * time.Second
	endpoints      = []string{"192.168.1.105:2379"} // 测试服
)

func EtcdPut() {
	cli, err := NewClientv3Tls(endpoints, dialTimeout)
	if err != nil {
		panic(err)
	}
	defer CloseClientv3(cli)

	// mysql
	cli.Put(context.TODO(), "/zyh/test/ry_mysql_url", "user:pwd@tcp(192.168.1.655:3306)/zyh_database?charset=utf8&loc=Asia%2FShanghai")
	cli.Put(context.TODO(), "/zyh/release/ry_mysql_url", "user:pwd@tcp(192.168.1.456:3306)/zyh_database?charset=utf8&loc=Asia%2FShanghai")
}
