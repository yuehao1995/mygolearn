package main

import "fmt"

func updateslice(s []int) {
	s[0] = 100
}
func main() {
	//1.从数组中截出切片
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	//2,3,4,5，包含数组下标为2的，不包含数组下标为6的
	//[2 3 4 5]
	s := arr[2:6]
	fmt.Println(s)
	//数组下标从0到6的数组值，不包括6
	//[0 1 2 3 4 5]
	fmt.Println("arr[:6] =", arr[:6])
	//从2开始之后的全部数组值
	//[2 3 4 5 6 7]
	s1 := arr[2:]
	fmt.Println("arr[2:] =", s1)
	//全部
	//[0 1 2 3 4 5 6 7]
	s2 := arr[:]
	fmt.Println("arr[:] =", s2)

	//2.切片是指针类型,该表时不仅改变了s1，而且改变了arr
	//[100 3 4 5 6 7]
	updateslice(s1)
	fmt.Println(s1)
	//[0 1 100 3 4 5 6 7]
	fmt.Println(arr)

	//slice是对原本arr的一个review，改变会改变本身及其切出来的所有
	//[100 1 100 3 4 5 6 7]
	updateslice(s2)
	fmt.Println(s2)
	//[100 1 100 3 4 5 6 7]
	fmt.Println(arr)

	//3.re slice
	//re slice的改变也是对开始arr的改变
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

}
