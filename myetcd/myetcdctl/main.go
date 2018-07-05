package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"
	"yd_key/kit"

	"github.com/astaxie/beego"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/pkg/transport"
)

var (
	RunEnv         string = "RYDEV" // 运行环境
	RunEnvName     string
	KeyOnline      = "/zyh/workers_online/"  // 在线
	KeyOutline     = "/zyh/workers_outline/" // 离线
	MYSQL_URI      string                    //主库
	MYSQL_LOG_URI  string                    //日志
	MYSQL_URI_READ string                    //只读库

	REDIS_URI string //redis URI
	REDIS_PWD string //redis 密码

)

func main() {
	RunEnv, RunEnvName = RunEnvi()
	endpoints := []string{}
	keyPrefix := ""
	if RunEnv == "ZYHRELEASE" {
		beego.Warn("当前运行生产模式请注意")
		endpoints = []string{"", "", ""} //ip
		keyPrefix = "/yd/release/"
	} else {
		endpoints = []string{""} //ip
		keyPrefix = "/zyh/test/"
	}
	beego.Warn("RunEnv", RunEnv, "Debug", endpoints)

	// boltdb文件读取配置项
	cli, err := NewClientv3Tls(endpoints, 5*time.Second)
	if err != nil {
		panic("etcd配置项读取出错" + err.Error())
	}
	defer CloseClientv3(cli)

	MYSQL_URI, _ = kit.EtcdGet(cli, keyPrefix+"zyh_mysql_url")
	MYSQL_URI_READ, _ = kit.EtcdGet(cli, keyPrefix+"zyh_mysql_url_read")

	REDIS_URI, _ = kit.EtcdGet(cli, keyPrefix+"redis_url")
	REDIS_PWD, _ = kit.EtcdGet(cli, keyPrefix+"redis_pwd")

}

// 创建clientv3
func NewClientv3(endpoints []string, dialTimeout time.Duration) (*clientv3.Client, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: dialTimeout,
	})
	if err != nil {
		log.Fatal("cliNew", err)
		return nil, err
	}
	return cli, err
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

func CloseClientv3(cli *clientv3.Client) {
	cli.Close()
}

// Worker 工作节点
type Worker struct {
	Name     string    `json:"name,omitempty"`      // 项目名
	IP       string    `json:"ip,omitempty"`        // 服务器ip
	Status   string    `json:"status,omitempty"`    // online outline
	CreateAt time.Time `json:"create_at,omitempty"` // 创建时间
	ModifyAt time.Time `json:"modify_at,omitempty"` // 修改时间
}

// NewWorker 创建工作节点实例
func NewWorker(endpoints []string, worker Worker, aliveTime int64) error {
	cli, err := NewClientv3Tls(endpoints, 5*time.Second)
	if err != nil {
		panic("etcd配置项读取出错" + err.Error())
	}
	// defer kit.CloseClientv3(cli)
	fmt.Println("NewWorker=", worker)
	worker.KeepAlive(cli, aliveTime)
	return nil
}

// KeepAlive 工作节点保持存活
func (m *Worker) KeepAlive(cli *clientv3.Client, aliveTime int64) error {
	fmt.Println("AliveWorker=", m)
	grant, err := cli.Grant(context.TODO(), aliveTime)
	if err != nil {
		log.Fatalln("KeepAliveErr=", err)
		return err
	}
	key := KeyOnline + m.Name
	val, err := json.Marshal(m)
	if err != nil {
		log.Fatalln("KeepAliveMarshalErr=", err)
		return err
	}
	fmt.Println(key, string(val), grant.ID, grant.TTL)
	_, err = cli.Put(context.TODO(), key, string(val), clientv3.WithLease(grant.ID))
	if err != nil {
		log.Fatalln("KeepAlivePutErr=", err)
		return err
	}
	live, err := cli.KeepAlive(context.TODO(), grant.ID)
	if err != nil {
		log.Fatalln("KeepAliveErr=", err)
		return err
	}
	ka := <-live
	fmt.Println("ttl:", ka.TTL, grant.ID, grant.TTL)

	return err
}
func RunEnvi() (env string, name string) {
	host := map[string]string{
		"rhqlkjewlkkjqwel": "(PROD)", // 正式服
		"rqwikjqwqwe":      "(TEST)", // 测试服
		"qwhlrrjqwoi":      "ZYHTEST",
		"werqjopweok":      "(DEV)", // 本地开发服
		"ZYH_werqjopweok":  "ZYHDEV",
		"yd_crm":           "(ALPHA)", // 预发服
		"ZYh_crm":          "ZYHALPHA",
	}
	n, err := os.Hostname()
	if err != nil {
		env, name = host["ENV_dev"], host["dev"]
		return
	}
	_, ok := host[n]
	if ok {
		env, name = host["ENV_"+n], host[n]
		return
	}
	fmt.Println("----RunEnv", n, env, name)
	env, name = host["ENV_dev"], host["dev"]
	return
}
