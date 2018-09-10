package main

import (
	"fmt"
	"mygolearn/gocodeanddoc/gointerface/complex/mock"
)

//接口的组合，定义两个小接口

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

const url = "http://www.imooc.com"

func post(poster Poster) {
	poster.Post("http://www.imooc.com",
		map[string]string{
			"name":   "ccmouse",
			"course": "golang",
		})
}

//定义一个大接口内部放入两个小接口
type RetrieverPoster interface {
	Retriever
	Poster
}

func seesion(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}

func main() {
	var r Retriever
	retriever := mock.Retriever{"this is a fake imooc.com"}
	//r只有Retriever接口能力
	r = &retriever
	fmt.Println(r)
	//retriever由于mock.Retriever{}结构体实现了两个接口的定义，因此具有两个接口的能力
	//修改retriever中的内容时，要使用指针
	fmt.Println(seesion(&retriever))
}
