package main

import (
	"fmt"
	"test01/muke/queue"
	tree2 "test01/muke/tree"
)

type myTreeNode struct {
	node *tree2.Node
}

/*
	后序遍历
*/
func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
}

func main()  {
	//var tree = tree2.Node{Value: 0}
	//tree.Left = &tree2.Node{Value: 1}
	//tree.Right = &tree2.Node{Value: 2}
	//tree.Left.Right = new(tree2.Node)
	//tree.Right.Left = &tree2.Node{Value: 5}
	//tree.Traverse()
	//myRoot := myTreeNode{&tree}
	//myRoot.postOrder()

	var q = queue.Queue{1}
	q.Push(2)
	q.Push(3)
	q.Push(4)
	fmt.Println(q)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
}

