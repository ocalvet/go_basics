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

// New creates a new Stacj
func New() *Stack {
	return &Stack{}
}

// Push an element to the top of the stack
func (stack *Stack) Push(v interface{}) {
	newNode := &node{nil, v}
	if stack.top == nil {
		stack.top = newNode
	} else {
		newNode.prev = stack.top
		stack.top = newNode
	}
}

// Pop pops and element from the top of the stack
func (stack *Stack) Pop() interface{} {
	if stack.top == nil {
		return nil
	}
	poppedElement := stack.top
	if poppedElement.prev != nil {
		stack.top = poppedElement.prev
	}
	return poppedElement
}

// Print the contents of the stack
func (stack *Stack) Print() {
	for node := stack.top; node != nil; {
		fmt.Println(node.v)
		node = node.prev
	}
}
