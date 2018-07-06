package main

import (
	"fmt"
	tools "mygolearn/mongodbtool/mymongotools"
)

func main() {
	session := tools.MgoSession.Copy()
	defer session.Close()
	st1 := &tools.Student{}
	c1 := session.DB("tool_mongo").C("tool_mongo")
	err := c1.Insert(st1)
	if err != nil {
		fmt.Println(err)
	}
}
