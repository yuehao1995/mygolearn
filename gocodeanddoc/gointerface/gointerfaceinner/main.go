package main

import (
	"fmt"
	"mygolearn/gocodeanddoc/gointerface/gointerfacedefine/mock"
	"mygolearn/gocodeanddoc/gointerface/gointerfacedefine/real"
	"time"
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

	r = mock.Retriever{"this is a fack imooc.com"}
	fmt.Println(download(r))
	//1.类型mock.Retriever {this is a fack imooc.com}
	fmt.Printf("%T %v \n", r, r)
	inspect(r)
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		Timeout:   time.Minute,
	}
	//2. 类型real.Retriever {Mozilla/5.0 1m0s}
	fmt.Printf("%T %v \n", r, r)

	//3.通过指针访问get时，只需要在定义结构体方法时，使用指针定义即可，使用时取结构体地址即可以
	//指针用于大结构体

	//4.获取接口类型的第一种方法switch
	inspect(r)

	//获取接口类型的第二种方法,type assertion
	//可以通过点.加具体类型从而获取interface里的具体类型
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.Timeout)

	//类型写错时会导致panic,因此可以这样判断
	if mockRetriever, ok := r.(mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	//接口变量里有实现者的类型，以及实现者的指针或值
	//接口变量自带指针
	//接口变量同样采用值传递，几乎不需要使用接口的指针
	//指针接收者只能用指针使用，值接收者两者都可以
	//表示任何类型interface{}
}

func inspect(r Retriever) {
	//获取接口类型
	//1. switch与r.type结合
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("Contents:", v.UserAgent)
	}
}
