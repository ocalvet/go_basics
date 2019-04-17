package queue

import "fmt"

// Queue datastructure
type Queue struct {
	root *node
}

type node struct {
	next *node
	v    interface{}
}

// New creates a new queue
func New() *Queue {
	return &Queue{}
}

// Enqueue queues an element into the queue
func (queue *Queue) Enqueue(v interface{}) {
	if queue.root == nil {
		queue.root = &node{nil, v}
		return
	}
	var n *node
	for n = queue.root; n.next != nil; n = n.next {
	}
	last := &node{nil, v}
	n.next = last
}

// Dequeue removes an element from the queue
func (queue *Queue) Dequeue() interface{} {
	return queue.root.v
}

// Print prints the elements of the queue in the console
func (queue *Queue) Print() {
	for n := queue.root; n != nil; n = n.next {
		fmt.Println(n.v)
	}
}
