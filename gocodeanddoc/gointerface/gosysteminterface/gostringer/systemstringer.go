package gostringer

import "fmt"

type StringIntemplete struct {
	Contents string
}

func (s StringIntemplete) String() string {
	return fmt.Sprintf("Retriever:{Contents=%s}", s.Contents)
}
