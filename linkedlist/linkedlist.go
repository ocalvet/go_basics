package linkedlist

import "fmt"

// LinkedList models a linked list
type LinkedList struct {
	root *node
}

type node struct {
	next  *node
	value interface{}
}

// New creates a linked list
func New() *LinkedList {
	return &LinkedList{}
}

// Iterator Returns an iterator function that can traverse the list
func (list *LinkedList) Iterator() func() (interface{}, bool) {
	current := list.root
	return func() (interface{}, bool) {
		hasNext := current.next != nil
		currentNode := current
		if currentNode != nil {
			current = current.next
			return currentNode.value, hasNext
		} else {
			return nil, hasNext
		}
	}
}

// Add adds a value to the list
func (list *LinkedList) Add(v interface{}) {
	if list.root == nil {
		list.root = &node{nil, v}
	} else {
		// find last node and insert another element
		n := list.root
		for n.next != nil {
			n = n.next
		}
		n.next = &node{nil, v}
	}
}

// Remove removes a value to the list
func (list *LinkedList) Remove(v interface{}) {
	if list.root == nil {
		return
	}
	if list.root.value == v {
		list.root = list.root.next
		return
	}
	// find last node and insert another element
	for n := list.root; n != nil; {
		if n.next.value == v {
			n.next = n.next.next
			return
		}
		n = n.next
	}
}

// Print prints the values in the list
func (list *LinkedList) Print() {
	n := list.root
	if n == nil {
		fmt.Println("Empty")
		return
	}
	for n != nil {
		fmt.Printf("%v\n", n.value)
		n = n.next
	}
}
