package tree

import (
	"fmt"
)

//go中无构造函数,需要使用时可以用工厂函数
//要改变内容必须使用指针接受者
//结构过大也要考虑指针接受者，拷贝性能问题
//一致性，如果有指针接受者，最好都使用指针接收者
//只有go语言有值接收者，且go中值接受者与指针接受者都可以直接使用
type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

//为结构体定义方法,结构体的方法需要一个接受者
func (node TreeNode) Print() {
	fmt.Println(node.Value)
}

//go
func (node *TreeNode) SetValue(value int) {
	if node == nil {
		fmt.Println("set value nil to node")
	}
	node.Value = value
}

func (node *TreeNode) traverse() {
	//空指针使用的好处，在内部判断空指针
	if node == nil {
		return
	}
	node.Left.traverse()
	node.Print()
	node.Right.traverse()
}

func CreateNode(value int) *TreeNode {
	//go中局部变量的地址也可以给外部变量使用
	//局部变量创建由垃圾回收决定。不返回给外部，编译器在栈上使用，返回给外部使用，在堆上分配
	return &TreeNode{Value: value}
}
