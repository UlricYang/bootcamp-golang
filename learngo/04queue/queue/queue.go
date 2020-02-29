package queue

import "fmt"

type Queue []interface{}

func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
	fmt.Printf("address(q)=%p\n", q)

}

func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	fmt.Printf("address(q)=%p\n", q)
	return head
}

func (q *Queue) IsEmpty() bool {
	fmt.Printf("address(q)=%p\n", q)
	return len(*q) == 0
}
