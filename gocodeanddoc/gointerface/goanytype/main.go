package main

import (
	"fmt"
	"mygolearn/gocodeanddoc/gointerface/goanytype/queue"
)

func main() {
	//go语言中可以用interface来代表不确定的类型
	//

	q := queue.Queue{1}

	q.Push(2)
	q.Push(3)

	//最开始的q与最后的q不是一个，因为q是指针接收者，可以改变值
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
}
