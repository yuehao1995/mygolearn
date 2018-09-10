package main

import (
	"fmt"
	"mygolearn/gocodeanddoc/gostruct/gopackage/gopackagelimit/tree"
)

func main() {
	//引用其他包时要记得大写是go封装的形式，包是封装的单位，如果想要内容要对其他包进行开放时，要记得大写
	var root tree.TreeNode

	root = tree.TreeNode{Value: 3}
	//go中指针可以直接用.来取
	root.Left = &tree.TreeNode{}
	root.Right = &tree.TreeNode{5, nil, nil}

	//new建一个空的treeNode，返回一个地址
	root.Right.Left = new(tree.TreeNode)
	root.Left.Right = tree.CreateNode(2)
	//定义切片时可以直接赋值
	nodes := []tree.TreeNode{
		{Value: 3},
		{},
		{6, nil, nil},
	}
	fmt.Println(nodes)
	//结构体的方法
	//编译器会自动识别指针或值
	//非指针调用时会拷贝一份
	root.Print()
	//指针调用时会取地址传过去
	root.SetValue(5)

	//go语言中nil指针也可以调用方法，go语言中nil指针是安全的，可以使用的
	//这样调用可以调方法，但是会报错
	//因为nil指针中取value值会出错
	var pRoot *tree.TreeNode
	pRoot.SetValue(200)
	pRoot = &root
	pRoot.SetValue(300)
	pRoot.Print()
}
