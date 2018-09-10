package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4, 5}
	s1 := arr[1:4]
	//第一次append会对原arr最后一个进行改变
	s2 := append(s1, 10)
	//第二次append会创建一个新的数组,把arr拷过去，新的arr长度会开的长一些,原来数组会回收掉
	s3 := append(s2, 11)
	s4 := append(s3, 12)
	fmt.Println(s2, s3, s4)
	fmt.Println("************************")
	fmt.Println(arr)

	//由于值传递的原因，必须将用一个变量去接append后的值
	s4 = append(s4, 13)
}

func createSlice() {
	//1. 定义一个切片变量
	var s []int //slice控制时为nil

	for i := 0; i < 100; i++ {
		//创建时len每次加一个长度，cap每次不够就扩展成两倍
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	//2.建立一个数组，然后用slice维护数组
	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	//3.只知道长度，不知道内容，建立一个slice
	s2 := make([]int, 16)

	//4.预估到slice涨的块，开辟32长度空间
	s3 := make([]int, 10, 32)
	printSlice(s1)
	printSlice(s2)
	printSlice(s3)
}

func printSlice(s []int) {
	fmt.Println("value=%v,len=%d,cap=%d", s, len(s), cap(s))
}

func copySlice() {
	s := []int{1, 2, 3, 4, 5, 6}
	var s1 []int
	//官方函数，复制,复制len，cap
	copy(s1, s)
	printSlice(s1)
}

func deleteSlice() {
	s := []int{1, 2, 3, 4, 5, 6}
	//s[4:]...可以将一个数组按可变参数内容传入相应的函数

	//删除下标为3，即第四个元素
	s = append(s[:3], s[4:]...)
	fmt.Println(s)

	//删除头
	fmt.Println("Poping from front")
	front := s[0]
	s2 := s[1:]
	fmt.Println(front)

	//删除尾
	fmt.Println("Poping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(tail)
}
