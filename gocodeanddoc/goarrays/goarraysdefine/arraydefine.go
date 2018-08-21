package main

import "fmt"

//数组是值类型,并且go语言认为不同的长度数组是不同类型,传递时用的拷贝
func main() {
	//定义一个未初始化的数组
	var arr1 [5]int
	//定义数组并赋值
	arr2 := [3]int{1, 3, 5}
	//定义数组并由编译器自动读出数量
	arr3 := [...]int{2, 4, 6, 8, 10}
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

	//遍历数组第一种方式
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	//遍历数组第二种方式，i是下标，v是值
	for i, v := range arr3 {
		fmt.Println(i, v)
	}
	change(&arr3)
	//定义一个二维数组
	var grid [4][5]int
	fmt.Println(grid)
}

func change(arr *[5]int) {
	//go语言指针使用灵活，可以不加*直接进行使用
	arr[0] = 100
	for _, v := range arr {
		fmt.Println(v)
	}
}
