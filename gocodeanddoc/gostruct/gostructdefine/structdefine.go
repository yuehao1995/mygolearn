package main

import (
	"fmt"
)

//go中无构造函数,需要使用时可以用工厂函数
//要改变内容必须使用指针接受者
//结构过大也要考虑指针接受者，拷贝性能问题
//一致性，如果有指针接受者，最好都使用指针接收者
//只有go语言有值接收者，且go中值接受者与指针接受者都可以直接使用
type treeNode struct {
	value       int
	left, right *treeNode
}

//为结构体定义方法,结构体的方法需要一个接受者
func (node treeNode) print() {
	fmt.Println(node.value)
}

//go
func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("set value nil to node")
	}
	node.value = value
}

func (node *treeNode) traverse() {
	//空指针使用的好处，在内部判断空指针
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func createNode(value int) *treeNode {
	//go中局部变量的地址也可以给外部变量使用
	//局部变量创建由垃圾回收决定。不返回给外部，编译器在栈上使用，返回给外部使用，在堆上分配
	return &treeNode{value: value}
}

func main() {
	//定义结构体变量
	var root treeNode

	root = treeNode{value: 3}
	//go中指针可以直接用.来取
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}

	//new建一个空的treeNode，返回一个地址
	root.right.left = new(treeNode)
	root.left.right = createNode(2)
	//定义切片时可以直接赋值
	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, nil},
	}
	fmt.Println(nodes)
	//结构体的方法
	//编译器会自动识别指针或值
	//非指针调用时会拷贝一份
	root.print()
	//指针调用时会取地址传过去
	root.setValue(5)

	//go语言中nil指针也可以调用方法，go语言中nil指针是安全的，可以使用的
	//这样调用可以调方法，但是会报错
	//因为nil指针中取value值会出错
	var pRoot *treeNode
	pRoot.setValue(200)
	pRoot = &root
	pRoot.setValue(300)
	pRoot.print()
}
