package main

import (
	"fmt"
	"time"
)

func main() {
	StringToTime()
	toString()
	AddTimeMonth()
}

func toString() {
	tm2, _ := time.Parse("01/02/2006", "00/00/0000")
	fmt.Println(tm2)
	//1970-01-01 12:17:46
}

func StringToTime() {
	t := time.Unix(15466, 0)
	fmt.Println(t.Format("2006-01-02 15:04:05"))
	//0001-01-01 00:00:00 +0000 UTC
}

func AddTimeMonth() {
	LoanDate := time.Now()
	//2018-7-11 15:47:54.4729745 +0800 CST
	fmt.Println(LoanDate.AddDate(0, 5, 0))
	//2018-12-11 15:47:54.4729745 +0800 CST
}
