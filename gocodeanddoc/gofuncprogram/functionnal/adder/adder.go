package main

import "fmt"

//返回值为函数
//这个返回函数参数是int，返回值是int
func adder() func(int) int {
	//自由变量
	sum := 0
	return func(v int) int {
		//v是局部变量
		sum += v
		return sum
	}
}

//闭包，函数体含有局部变量，函数体外含有自由变量,是它的环境中的变量
//闭包，函数作为返回值，且外部调用这个函数。
//这个函数中返回值，且这个返回值为自由变量

func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(a(i))
	}
}
