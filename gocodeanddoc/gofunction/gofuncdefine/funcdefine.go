package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

//函数名在前,参数在函数名后传入,返回值在最后,多个返回值记得加括号
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		//其中返回值如果有不用的话可以用_省略
		q, _ := div(a, b)
		return q, nil
	default:
		//格式化输入error
		return 0, fmt.Errorf("unsupported operation:%s", op)
	}
}

//go 函数返回值可以为多个,多个返回值一般为返回值与error
func div(a, b int) (int, int) {
	return a / b, a % b
}

//go 返回值可以命名
func div1(a, b int) (q int, r int) {
	//给q和r赋值,当函数返回时,q,r会自动到返回值里,函数体长时,不好明确在哪赋的值,因此,不推荐
	q = a / b
	r = a % b
	return
}

//go语言为函数式编程，函数的参数，返回值，函数体里都可以包括函数
func apply(op func(int, int) int, a, b int) int {
	//反射获取函名
	//反射获取op的值,获取函数指针
	p := reflect.ValueOf(op).Pointer()
	//获取函数名
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args "+"(%d,%d)", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	//pow,a的b次方
	return int(math.Pow(float64(a), float64(b)))
}

//可变参数,可以传多个number,像数组一样用
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}
func main() {
	fmt.Println(eval(3, 4, "*"))
	q, r := div1(13, 3)
	fmt.Println(div(q, r))
	fmt.Println(div(13, 3))
	//81
	//Calling function main.pow with args (3,4)81
	fmt.Println(apply(pow, 3, 4))
	//匿名函数
	//Calling function main.main.func1 with args (3,4)81
	fmt.Println(apply(
		func(a, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 4))
	//15
	fmt.Println(sum(1, 2, 3, 4, 5))
}
