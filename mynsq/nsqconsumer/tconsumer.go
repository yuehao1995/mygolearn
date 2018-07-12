package main

import (
	"log"
	"time"

	"github.com/nsqio/go-nsq"
)

func main() {
	config := nsq.NewConfig()
	q, err := nsq.NewConsumer("write_test", "ch", config)
	if err != nil {
		log.Panic("Could not create consumer.")
	}
	defer q.Stop()
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Printf("Got a message: %v", string(message.Body))
		time.Sleep(5 * time.Second)
		return nil
	}))
	//err = q.ConnectToNSQD("192.168.9.111:32771");
	err = q.ConnectToNSQD("192.168.9.111:4150")
	if err != nil {
		log.Panic("Could not connect")
	}
	time.Sleep(3600 * time.Second)

	/*
			2017/08/29 15:29:45 INF    1 [write_test/ch] (192.168.9.111:4150) connecting to nsqd
		2017/08/29 15:29:45 Got a message: {"Id":0,"Name":"Jack0","Age":0,"NickName":"Luo0"}
		2017/08/29 15:29:50 Got a message: {"Id":1,"Name":"Jack1","Age":1,"NickName":"Luo1"}
		2017/08/29 15:29:55 Got a message: {"Id":2,"Name":"Jack2","Age":2,"NickName":"Luo2"}
		2017/08/29 15:30:00 Got a message: {"Id":3,"Name":"Jack3","Age":3,"NickName":"Luo3"}
		2017/08/29 15:30:05 Got a message: {"Id":4,"Name":"Jack4","Age":4,"NickName":"Luo4"}
		2017/08/29 15:30:10 Got a message: {"Id":5,"Name":"Jack5","Age":5,"NickName":"Luo5"}
		2017/08/29 15:30:15 Got a message: {"Id":6,"Name":"Jack6","Age":6,"NickName":"Luo6"}
		2017/08/29 15:30:20 Got a message: {"Id":7,"Name":"Jack7","Age":7,"NickName":"Luo7"}
		2017/08/29 15:30:25 Got a message: {"Id":8,"Name":"Jack8","Age":8,"NickName":"Luo8"}
		2017/08/29 15:30:30 Got a message: {"Id":9,"Name":"Jack9","Age":9,"NickName":"Luo9"}
		2017/08/29 15:30:35 Got a message: {"Id":10,"Name":"Jack10","Age":10,"NickName":"Luo10"}
		2017/08/29 15:30:40 Got a message: {"Id":11,"Name":"Jack11","Age":11,"NickName":"Luo11"}
		2017/08/29 15:30:45 Got a message: {"Id":12,"Name":"Jack12","Age":12,"NickName":"Luo12"}*/
}
