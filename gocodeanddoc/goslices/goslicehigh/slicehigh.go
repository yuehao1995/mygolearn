package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	s1 := arr[2:6]
	s2 := s1[3:5]
	fmt.Println(s2)
}
