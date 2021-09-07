package tree

import "fmt"

type Node struct {
	Value int
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Println(node.Value)
}

func (node *Node) setValue(value int) {
	node.Value = value
}
/*
	这种写法在C++是非法行为, 不能在函数体中的局部地址返回给外部使用
	局部变量是分配在栈上面的, 函数一旦退出局部变量就会被立刻销毁,
	传出去就要在堆上分配, 在推上分配就要手动释放 --> 这是C++的做法
	其他语言一般都是分配在推上面的 比如 new...

	在golang中, 无需知道分配, 交给系统, 之后会有垃圾回收机制
*/
func CreateTreeNode(value int) *Node {
	return &Node{ Value: value }
}

/*
	中序遍历
*/
func (node *Node) Traverse()  {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}