package main

import (
	"fmt"
	"time"
)

func main() {
	StringToTime()
	toString()
}

func toString() {
	tm2, _ := time.Parse("01/02/2006", "00/00/0000")
	fmt.Println(tm2)
}

func StringToTime() {
	t := time.Unix(15466, 0)
	fmt.Println(t.Format("2006-01-02 15:04:05"))
}
