package znsq

import (
	"encoding/json"
	"fmt"
	"log"
	"mygolearn/mynsq/znsq/models"
	"mygolearn/mynsq/znsq/nsq"
	"os"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	nsqd "github.com/nsqio/go-nsq"
)

// 初始化消费者 最好是在项目初始化的时候就调用,否则可能会出现orm注册model冲突的情况
func NewNsqConsumer(topic, channel, address, mysql string, handle nsq.Handler) *nsq.Consumer {
	// 注册数据库
	_, err := orm.GetDB("default_nsq")
	if err != nil {
		fmt.Println("[ry_core-znsq-consumer]orm.RegisterDataBase=", mysql)
		err = orm.RegisterDataBase("default_nsq", "mysql", mysql, 10, 10)
	}
	if err != nil {
		panic("nsq NewNsqConsumer注册数据库失败:" + err.Error())
	}
	cfg := nsqd.NewConfig()
	cfg.AuthSecret = "hey,zcm"                      // nsq认证密钥,暂时不需要
	c, err := nsqd.NewConsumer(topic, channel, cfg) // 新建一个消费者
	if err != nil {
		panic(err)
	}
	c.SetLogger(log.New(os.Stdout, "nsq consumer:", 0), nsq.LogLevelInfo)

	rg := &handlerRegist{
		h:       handle,
		topic:   topic,
		channel: channel,
		nsqd:    address,
	}
	// 获取消费者ip
	ip, err := getLocalIp()
	if err != nil {
		panic("nsq NewNsqConsumer:" + err.Error())
	}
	rg.ip = ip
	c.AddHandler(rg) // 添加消费者接口
	if err := c.ConnectToNSQD(address); err != nil {
		panic(err)
	}
	return c
}

// dbregister isn't my job
func NewConsumerV2(topic, channel, address string, handle nsq.Handler) *nsq.Consumer {
	_, err := orm.GetDB("default")
	if err != nil {
		panic(topic + "/" + channel + "的nsq orm.数据库default未注册 " + err.Error())
	}
	cfg := nsqd.NewConfig()
	// cfg.AuthSecret = "hey,zcm" // nsq认证密钥,暂时不需要
	c, err := nsqd.NewConsumer(topic, channel, cfg) // 新建一个消费者
	if err != nil {
		panic(" " + err.Error())
	}
	c.SetLogger(log.New(os.Stdout, "nsq consumer2: ", 0), nsq.LogLevelInfo)
	rg := &handlerRegist{
		h:       handle,
		topic:   topic,
		channel: channel,
		nsqd:    address,
	}
	// 获取消费者ip
	ip, err := getLocalIp()
	if err != nil {
		panic("nsq NewNsqConsumer:" + err.Error())
	}
	rg.ip = ip
	c.AddHandler(rg) // 添加消费者接口
	if err := c.ConnectToNSQD(address); err != nil {
		panic(err)
	}
	return c
	// don't forget to close it if not use
}

type handlerRegist struct {
	h       nsqd.Handler
	topic   string
	channel string
	nsqd    string
	ip      string
}

// 对调用方的handel封装
func (rg *handlerRegist) HandleMessage(message *nsqd.Message) error {
	go rg.ConsumeLog(message) // 不知道会不会造成gc压力
	return rg.h.HandleMessage(message)
}

// 添加消费日志到mysql
func (rg *handlerRegist) ConsumeLog(message *nsqd.Message) (int64, error) {
	log := &models.NsqConsumeLog{}
	log.ConsumerIp = rg.ip
	log.NsqdUrl = rg.nsqd
	log.Topic = rg.topic
	log.Channel = rg.channel
	log.Message = string(message.Body)
	log.MessageId = getMessageId(message.Body)
	return models.AddNsqConsumeLog(log)
}

func getMessageId(body []byte) string {
	m := map[string]interface{}{}
	json.Unmarshal(body, &m)
	if v, ok := m["nsq_msg_uuid"].(string); ok {
		return v
	}
	return ""
}
