package main

import (
	"fmt"
	"math"
)

func consts() {
	const filename = "abc.txt"
	//a,b不规定类型(做文本替换),则类型不确定,可以为int或者float
	const a, b = 3, 4
	const (
		cfilename = "abc.txt"
		//a,b不规定类型(做文本替换),则类型不确定,可以为int或者float
		d, e = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(c)
}

func enums() {
	//普通枚举类型
	const (
		cpp    = 0
		java   = 1
		python = 2
		golang = 4
	)
	//简化(自增枚举类型)
	const (
		cpp1    = iota
		_       //省略1
		java1   //2
		python1 //3
		golang1 //4
	)
	//1 1024 1048576 1073741824 1099511627776 1125899906842624
	const (
		b  = 1 << (10 * iota) //1左移0
		kb                    //1左移 10
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp1, java1, python1, golang1)
	fmt.Println(b, kb, mb, gb, tb, pb)
}
func main() {
	enums()
}
