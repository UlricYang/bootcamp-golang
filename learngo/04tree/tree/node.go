package tree

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// node参数是传值，不是传引用，go所有都是传值(即为拷贝变量)
func (node *Node) Print() {
	fmt.Println(node.Value)
}

// 为了正儿八经Set值，需要改变内容，所以接收者要用指针方式
// 结构过大也使用指针接收者
// 一致性：如有指针接收者，最好都用指针
func (node *Node) SetValue(Value int) {
	if node == nil {
		fmt.Println("Setting Value to nil node")
		return
	}
	node.Value = Value
}

func (node Node) SetValue2(Value int) {
	node.Value = Value
}

// 返回了局部变量，这也是对的，无需知道是分配在堆还是栈上，有垃圾回收机制
// 如果是返回的变量，可能是栈上;如果是指针，则大概是在堆上
func CreateNode(Value int) *Node {
	return &Node{Value: Value}
}
