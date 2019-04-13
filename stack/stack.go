package stack

import "fmt"

// Stack datastructure
type Stack struct {
	top *node
}

type node struct {
	prev *node
	v    interface{}
}

func New() *Stack {
	return &Stack{}
}

func (stack *Stack) Push(v interface{}) {
	newNode := &node{nil, v}
	if stack.top == nil {
		stack.top = newNode
	} else {
		newNode.prev = stack.top.prev
		stack.top = newNode
	}
}

func (stack *Stack) Pop() interface{} {
	return stack.top.v
}

func (stack *Stack) Print() {
	for node := stack.top; node != nil; {
		fmt.Println(node.v)
		node = node.prev
	}
}
