package tree

import "fmt"

func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}

func (node *Node) Traverse2() {
	node.TraverseFunc(func(n *Node) { n.Print() })
	fmt.Println()
}
func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}

func (node *Node) TraverseWithChannel() chan *Node {
	out := make(chan *Node, 0)
	go func() {
		node.TraverseFunc(func(n *Node) {
			out <- n
		})
		close(out)
	}()
	return out
}
