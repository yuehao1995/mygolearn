package main

import (
	"fmt"
	"strings"
)

/*
*Join  os.StdOut strings.Join(s,"|")------------- hello|world|hhh
 */
func main() {
	s := []string{"hello", "world", "hhh"}
	fmt.Println(`strings.Join(s,"|")-------------`, strings.Join(s, "|"))
}
