package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	const filename = "abc.txt"
	//读文件内容
	/*
	*		abcde
	*	    qwert
	*	    asdfg
	*	    zxcvb
	 */
	contents, err := ioutil.ReadFile(filename)
	//if的判断是不能带()的
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	//go if,else 简化版
	//if条件里可以定义变量
	if contents1, err1 := ioutil.ReadFile(filename); err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Printf("%s\n", contents1)
	}
}
