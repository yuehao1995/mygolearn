package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	s := `abc"d"
	kkkk
123
	`
	//io中的Reader与Writer接口
	printFileContents(strings.NewReader(s))
	//bytes.NewReader()
}
