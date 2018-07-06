package tools

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
)

var (
	MAX_POOL_LIMIT = 50
	MgoSession     *mgo.Session
)

const MgoUri = "mongodb://zyhwrite:zyhpwd@192.168.1.8954:27017/zyh_database"

/**
 * 公共方法，获取session，如果存在则拷贝一份
 */

type Student struct {
	Age    string
	Name   string
	CardNo string
}

func GetSession() *mgo.Session {
	if MgoSession == nil {
		var err error
		MgoSession, err = mgo.Dial(MgoUri)
		if err != nil {
			panic(err) //直接终止程序运行
		}
	} else {
		err := MgoSession.Ping()
		if err != nil {
			MgoSession, err = mgo.Dial(MgoUri)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	MgoSession.SetPoolLimit(MAX_POOL_LIMIT)
	//最大连接池默认为4096
	return MgoSession.Clone()
}

func init() {
	MgoSession = GetSession()
}

func saveToMongo(StructTool string) {
	session := MgoSession.Copy()
	defer session.Close()
	session.DB("tools_mongo").C("tools_student").Insert()
}
