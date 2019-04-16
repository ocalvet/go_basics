package main

import (
	"fmt"

	"go_basics/btree"
	"go_basics/linkedlist"
	"go_basics/stack"
)

func main() {
	// runLinkedListSample()
	// runStackSample()
	runQueueSample()
	// runBTreeSample()
}

func runQueueSample() {
	fmt.Println("Queue example")
}

func runLinkedListSample() {
	fmt.Println("Link list example")
	list := linkedlist.New()
	list.Add(4)
	list.Add(23)
	list.Add(2)
	list.Print()
	list.Remove(2)
	list.Print()
}

func runStackSample() {
	fmt.Println("STACK")
	st := stack.New()
	st.Push(1)
	st.Push(4)
	st.Push(6)
	st.Print()
	st.Pop()
	st.Print()
}

func runBTreeSample() {
	bt := btree.New(func(a interface{}, b interface{}) int {
		vA := a.(int)
		vB := b.(int)
		if vA < vB {
			return -1
		} else if vA == vB {
			return 0
		} else {
			return 1
		}
	})

	bt.Add(3)
	bt.Print()
}
