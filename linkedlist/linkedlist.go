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

// Remove removes a value to the list
func (list *LinkedList) Remove(v interface{}) {
	if list.head == nil {
		return
	}
	if list.head.value == v {
		list.head = list.head.next
		return
	}
	// find last node and insert another element
	for n := list.head; n != nil; {
		if n.next.value == v {
			n.next = n.next.next
			return
		}
		n = n.next
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
