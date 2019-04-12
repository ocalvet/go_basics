package linkedlist

import "fmt"

// LinkedList models a linked list
type LinkedList struct {
	head *node
}

type node struct {
	next  *node
	value interface{}
}

// New creates a linked list
func New() *LinkedList {
	return &LinkedList{}
}

// Add adds a value to the list
func (list *LinkedList) Add(v interface{}) {
	if list.head == nil {
		list.head = &node{nil, v}
	} else {
		// find last node and insert another element
		n := list.head
		for n.next != nil {
			n = n.next
		}
		n.next = &node{nil, v}
	}
}

// Print prints the values in the list
func (list *LinkedList) Print() {
	n := list.head
	if n == nil {
		fmt.Println("Empty")
		return
	}
	for n != nil {
		fmt.Printf("%v\n", n.value)
		n = n.next
	}
}
