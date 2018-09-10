package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"mygolearn/gocodeanddoc/errhanding/fib"
	"os"
)

func tryDefer() {
	//退出时打印
	//defer里相当于有一个栈，按照定义顺序来讲，先进后出
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)

	for i := 0; i < 100; i++ {
		//此时会倒序打印出从30到0
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many")
		}
	}
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//建立缓冲
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := fib.Fibonacci()
	for i := 0; i <= 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func writeFile1(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE|os.O_TRUNC, 0666)
	err = errors.New("this is a custom error")
	if err != nil {
		if PathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s,%s,%s\n", PathError.Op, PathError.Path, PathError.Err)
		}
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	//tryDefer()
	//fib.txt在mygolearn目录下
	writeFile("fib.txt")
}
