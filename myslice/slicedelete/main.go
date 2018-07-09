package main

import (
	"fmt"
)

func main() {
	//1. 声明字符串类型的切片变量，go自动初始化为nil，长度：0，地址：0
	var ss []string
	fmt.Printf("length:%v \taddr:%p \tisnil:%v", len(ss), ss, ss == nil)
	//2. 切片尾部追加元素
	for i := 0; i < 10; i++ {
		ss = append(ss, fmt.Sprintf("key%d", i))
	}
	fmt.Printf("\n切片尾部追加元素:%s", ss)
	fmt.Printf("\nlength:%v \taddr:%p \tisnil:%v", len(ss), ss, ss == nil)
	//3. 删除切片指定索引的元素
	index := 3
	ss = append(ss[:index], ss[index+1:]...)
	fmt.Printf("\n删除切片指定索引的元素:%s", ss)
	fmt.Printf("\nlength:%v \taddr:%p \tisnil:%v", len(ss), ss, ss == nil)
	//4. 在切片中间插入元素,注意保存后面一部分的元素，必须新建一个临时temp切片
	index = 5
	temp := append([]string{}, ss[index:]...)
	ss = append(ss[:index], "key999")
	ss = append(ss, temp...)
	fmt.Printf("\n在切片中间插入元素:%s", ss)
	fmt.Printf("\nlength:%v \taddr:%p \tisnil:%v", len(ss), ss, ss == nil)
}
