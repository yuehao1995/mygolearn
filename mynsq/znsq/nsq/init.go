package nsq

import (
	"ry_core/znsq"
)

var (
	Wr        *znsq.Producer
	NsqPoint  = "127.0.0.1:14150"
	MYSQL_URI = "127.0.0.1:3306"
)

func init() {
	Wr = znsq.NewProducer(NsqPoint, MYSQL_URI)
}

type RyMsgDataReq struct {
	Message string
}
