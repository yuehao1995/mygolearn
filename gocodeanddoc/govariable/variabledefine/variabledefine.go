package variabledefine

import "fmt"

var (
	//包内变量
	aa = 3
	ss = "zyh"
	bb = true
	//函数外定义变量时,下边的简写是错误的
	//bb:=3
)

//未手动初始化的变量
func variableZeroValue() {
	//go语言有默认值,int默认为0,string默认为 "",但是在要定义变量,因此使用结构体时最好先定义一个空结构体,这样子会有一个初始化值
	var a int
	var s string
	//输出值为0,""打不出
	fmt.Println(a, s)
	//%q带""的双引号引起来的 Go 语法字符串,%s普通字符串
	fmt.Printf("%d %q %s\n", a, s, s)
}

//进行了初始化的变量定义
func variableInitialValue() {
	//go 语言变量名在类型之前,且可以定义多个变量在类型之前,并对多个变量进行赋值,赋值时变量与值数量应该相等
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

//赋值时自动匹配类型,编译器自动决定类型
func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

//更简略的变量定义
func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}
func main() {
	fmt.Println("Hello world")

}
