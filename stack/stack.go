package stack

import "fmt"

// Stack datastructure
type Stack struct {
	root *node
}

type node struct {
	next *node
	v    interface{}
}

// New creates a new Stacj
func New() *Stack {
	return &Stack{}
}

// Push an element to the root of the stack
func (stack *Stack) Push(v interface{}) {
	newNode := &node{nil, v}
	if stack.root == nil {
		stack.root = newNode
	} else {
		newNode.next = stack.root
		stack.root = newNode
	}
}

// Pop pops and element from the root of the stack
func (stack *Stack) Pop() interface{} {
	if stack.root == nil {
		return nil
	}
	poppedElement := stack.root
	if poppedElement.next != nil {
		stack.root = poppedElement.next
	}
	return poppedElement
}

// Print the contents of the stack
func (stack *Stack) Print() {
	for node := stack.root; node != nil; {
		fmt.Println(node.v)
		node = node.next
	}
}
