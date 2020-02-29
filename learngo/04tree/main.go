package main

import "fmt"
import "learngo/04tree/tree"

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	l := myTreeNode{myNode.node.Left}
	r := myTreeNode{myNode.node.Right}
	l.postOrder()
	r.postOrder()
	myNode.node.Print()
}
func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	nodes := []tree.Node{
		{Value: 3}, {}, {6, nil, &root},
	}
	fmt.Println(nodes)

	root.Print()
	root.Right.Left.SetValue2(4)
	root.Right.Left.Print()
	root.Right.Left.SetValue(4)
	root.Right.Left.Print()

	proot := &root
	proot.Print()
	proot.SetValue(5)
	proot.Print()

	var rootp *tree.Node
	rootp.SetValue(6)
	rootp = &root
	rootp.SetValue(7)
	rootp.Print()

	fmt.Println("Traverse")
	root.Traverse()
	fmt.Println("Traverse2")
	root.Traverse2()
	nodecount := 0
	root.TraverseFunc(func(tn *tree.Node) { nodecount++ })
	fmt.Println("nodecount:", nodecount)

	fmt.Println("Traverse post order")
	myroot := myTreeNode{&root}
	myroot.postOrder()

	fmt.Println("Traverse with channel")
	c := root.TraverseWithChannel()
	maxNode := 0
	for node := range c {
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}
	fmt.Println("max node value: ", maxNode)
}
