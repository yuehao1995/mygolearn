package znsq

import (
	"encoding/json"
	"fmt"
	"log"
	"mygolearn/mynsq/znsq/models"
	"os"
	"sync"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/nsqio/go-nsq"
)

type Producer struct {
	P  *nsq.Producer
	ip string
}

// 初始化生产者 最好是在项目初始化的时候就调用,否则可能会出现orm注册model冲突的情况
func NewProducer(addr string, mysql string) (wr *Producer) {
	// 注册数据库
	_, err := orm.GetDB("default_nsq")
	if err != nil {
		fmt.Println("[ry_core-znsq-producer]orm.RegisterDataBase=", mysql)
		err = orm.RegisterDataBase("default_nsq", "mysql", mysql, 10, 10)
	}
	if err != nil {
		panic("nsq NewProducer注册数据库失败:" + err.Error())
	}
	// 生产者配置,这里使用基本配置
	cfg := nsq.NewConfig()
	cfg.AuthSecret = "hey,zcm" // nsq认证密钥,暂时不需要
	p, err := nsq.NewProducer(addr, cfg)
	if err != nil {
		panic(err)
	}
	// 设置日志级别
	p.SetLogger(log.New(os.Stdout, "nsq producer:", 0), nsq.LogLevelInfo)
	wr = &Producer{P: p}
	// 获取生产者ip
	ip, err := getLocalIp()
	if err != nil {
		panic("nsq NewProducer:" + err.Error())
	}
	wr.ip = ip
	return wr
}

func (p *Producer) Publish(topic string, body []byte) error {
	body, uuid := appendMessageId(body)
	go p.PublishLog(topic, uuid, body) // 添加日志
	return p.P.Publish(topic, body)
}

func (p *Producer) PublishLog(topic, uuid string, body []byte) (int64, error) {
	log := &models.NsqPublishLog{}
	// log.MessageId = "" // 因为messageId由nsqd生成,所以这里还无法获取messageId
	log.Message = string(body)
	log.NsqdUrl = p.P.String()
	log.Topic = topic
	log.ProducerIp = p.ip
	log.MessageId = uuid
	return models.AddNsqPublishLog(log)
}

func (p *Producer) Stop() {
	p.P.Stop()
}

// 为生产者推送的消息添加message_id字段,供消费者使用
func appendMessageId(body []byte) ([]byte, string) {
	m := map[string]interface{}{}
	err := json.Unmarshal(body, &m)
	if err != nil {
		// 消息内容不是json格式
		return body, ""
	}
	uuid := NewUUID().Hex()
	m["nsq_msg_uuid"] = uuid
	body, _ = json.Marshal(m)
	return body, uuid
}

var bytesPool = sync.Pool{New: newBytes}

func newBytes() interface{} {
	bytes := make([]byte, 0, 500)
	return bytes
}
