package main

import (
	"fmt"
	"mygolearn/gocodeanddoc/gointerface/gosysteminterface/gostringer"
)

type StringInterface interface {
	String() string
}

func main() {
	s := gostringer.StringIntemplete{"this is zyh"}
	inspect(s)
}

func inspect(s StringInterface) {
	switch v := s.(type) {
	case gostringer.StringIntemplete:
		fmt.Println(v.Contents)
	}
}
