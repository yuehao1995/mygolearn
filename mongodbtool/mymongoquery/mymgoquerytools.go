package query

import (
	tools "mygolearn/mongodbtool/mymongotools"
	"ry_dock/utils/mongodb"

	"github.com/astaxie/beego/logs"

	"gopkg.in/mgo.v2/bson"
)

var collectionName = "mgo_stu"

func QueryMgo(idCard string) interface{} {
	s := tools.GetSession()
	defer s.Close()
	var result interface{}
	c := s.DB(mongodb.DbName).C(collectionName)
	err := c.Find(bson.M{"idCard": idCard}).One(result)

	err = c.Find(bson.M{"idCard": idCard}).Sort("-createTime").Limit(10).Skip(5).One(&result)
	logs.Error(err)
	return result
}
