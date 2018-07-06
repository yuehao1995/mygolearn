package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Unix(0, 0)
	fmt.Println(t.Format("2006-01-02 15:04:05"))
}
