package cache

import (
	"context"
	"log"
	"path"
	"path/filepath"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/pkg/transport"
	"github.com/gomodule/redigo/redis"
)

var (
	dialTimeout    = 5 * time.Second
	requestTimeout = 2 * time.Second
	endpoints      = []string{"192.168.1.105:2379"} // 测试服
)

var RedisPool *redis.Pool

// 初始化连接池
func InitRedis(host, pwd string, conns ...int) {
	maxIdle := 2
	maxActive := 10
	if len(conns) > 0 {
		maxIdle = conns[0]
	}
	if len(conns) > 1 {
		maxActive = conns[1]
	}
	RedisPool = &redis.Pool{
		MaxIdle:     maxIdle,           // 最大的空闲连接数
		MaxActive:   maxActive,         // 最大的激活连接数
		IdleTimeout: 180 * time.Second, // 最大的空闲连接等待时间
		// 建立连接
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host, redis.DialPassword(pwd))
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
}

// 创建clientv3-tls
func NewClientv3Tls(endpoints []string, dialTimeout time.Duration) (*clientv3.Client, error) {
	src, err := filepath.Abs("../")
	log.Println("[NewClientv3Tls]", src, err, path.Join(src, "zyh_core/conf/ssl/ca.pem"))
	tlsInfo := transport.TLSInfo{
		CertFile:      path.Join(src, "zyh_core/conf/ssl/etcd.pem"),
		KeyFile:       path.Join(src, "zyh_core/conf/ssl/etcd-key.pem"),
		TrustedCAFile: path.Join(src, "zyh_core/conf/ssl/ca.pem"),
	}
	tlsConfig, err := tlsInfo.ClientConfig()
	if err != nil {
		panic("配置文件tls读取错误" + err.Error())
	}
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: dialTimeout,
		TLS:         tlsConfig,
	})
	if err != nil {
		log.Fatal("cliTlsNew=", err)
		return nil, err
	}
	return cli, err
}

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

func CloseClientv3(cli *clientv3.Client) {
	cli.Close()
}

// 获取etcd指定key的值
func EtcdGet(cli *clientv3.Client, key string) (string, error) {
	resp, err := cli.Get(context.TODO(), key)
	if err != nil {
		return "", err
	}
	if resp.Count == 0 {
		return "", err
	}
	return string(resp.Kvs[0].Value), err
}
