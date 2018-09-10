package main

import (
	"fmt"
	"mygolearn/gocodeanddoc/gostruct/gopackage/gopackagelimit/tree"
)

//组合扩展
//用一个结构体将原结构体包装，定义新结构体的方法从而进行扩展
type myTreeNode struct {
	node *tree.TreeNode
}

//指针接受者要定义一个变量才能继续使用其方法
func (myNode *myTreeNode) postOrder() {
	//结构体为nil里含有内容，用.调用时是会报错的
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()

}

func main() {
	var root tree.TreeNode

	root = tree.TreeNode{Value: 3}
	//go中指针可以直接用.来取
	root.Left = &tree.TreeNode{}
	root.Right = &tree.TreeNode{5, nil, nil}

	//new建一个空的treeNode，返回一个地址
	root.Right.Left = new(tree.TreeNode)
	root.Left.Right = tree.CreateNode(2)
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()
}
