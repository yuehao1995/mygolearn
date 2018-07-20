package main

import "fmt"

func grade(score int) string {
	g := ""
	//switch语句是自上往下走,遇到匹配就执行,执行完就跳出,fallthrough可以强制执行
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score:%d", score))
	case score < 60:
		g = "F"
	case score < 80, score < 90:
		g = "C"
	case score < 100:
		g = "A"
	}
	return g
}

func grade1() {
	//switch无break,但可以自动跳出,switch可以无表达式,case后边可以跟表达式
	switch {
	case false:
		fmt.Println("The integer was <= 4")
		fallthrough
	case true:
		fmt.Println("The integer was <= 5")
		fallthrough
	case false:
		fmt.Println("The integer was <= 6")
		fallthrough
	case true:
		fmt.Println("The integer was <= 7")
		fallthrough
	case false:
		fmt.Println("The integer was <= 8")
	default:
		fmt.Println("default case")
	}
}
func main() {
	fmt.Println(
		grade(0),
		grade(59),
		grade(69),
		grade(79),
		grade(89),
		grade(99),
	)
	grade1()
}
