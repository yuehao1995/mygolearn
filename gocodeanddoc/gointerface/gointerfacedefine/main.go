package main

import (
	"fmt"
	"mygolearn/gocodeanddoc/gointerface/gointerfacedefine/mock"
	"mygolearn/gocodeanddoc/gointerface/gointerfacedefine/real"
)

//1.实现者不需要说明自己实现了哪个接口，只要实现接口中的方法
//2.使用者规定Retriever必须有get方法
//接口的实现是隐式的，不需要说明自己实现了哪个接口，只需要实现接口中的方法就行
type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func main() {
	var r Retriever
	//接口调用时使用实现接口的结构体
	r = mock.Retriever{"this is a fack imooc.com"}
	fmt.Println(download(r))
	//使用的第二种方法
	fmt.Println(download(
		mock.Retriever{"this is a fack imooc.com"}))
	//真实的retriever
	fmt.Println(download(
		real.Retriever{}))

}
