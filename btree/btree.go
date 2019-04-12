package btree

import "fmt"

// BTree btree model
type BTree struct {
	head     *node
	comparer Comparer
}

type node struct {
	left  *node
	right *node
	v     interface{}
}

// Comparer is a function that takes two values and returns
//   -1 if a < b, +1 if a > b and 0 if a = b
type Comparer func(a interface{}, b interface{}) int

// New creates a new btree
func New(comparer Comparer) *BTree {
	return &BTree{nil, comparer}
}

// Add adds an element to the btree
func (btree *BTree) Add(v interface{}) {
	if btree.head == nil {
		btree.head = &node{nil, nil, v}
	} else {
		c := btree.comparer(btree.head.v, v)
		if c == -1 {
			fmt.Println("Less than")
		} else if c == 0 {
			fmt.Println("Equal")
		} else {
			fmt.Println("More than")
		}
	}
}

// Print prints the contents of a btree
func (btree *BTree) Print() {
	fmt.Printf("%v\n", btree.head.v)
}
