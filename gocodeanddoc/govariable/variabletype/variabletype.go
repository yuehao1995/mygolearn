package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func euler() {
	//(0+1.2246467991473515e-16i)
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
}

func triangle() {
	//5 go语言无隐式转换,必须强制转换
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func main() {
	euler()
	triangle()
}
