package main

import (
	"encoding/json"
	nsql "mygolearn/mynsq/znsq/nsq"
)

func main() {
	msg := nsql.RyMsgDataReq{Message: "信息"}
	bytes, _ := json.Marshal(msg)
	nsql.Wr.Publish("MsgHandle", bytes)
}
