package queue

import "fmt"

// Queue datastructure
type Queue struct {
	root *node
}

type node struct {
	prev *node
	v    interface{}
}

// New creates a new queue
func New() *Queue {
	return &Queue{}
}

// Enqueue queues an element into the queue
func (Queue *Queue) Enqueue(v interface{}) {
	if Queue.root == nil {
		Queue.root = &node{nil, v}
	}
}

// Dequeue removes an element from the queue
func (Queue *Queue) Dequeue() interface{} {
	return Queue.root.v
}

// Print prints the elements of the queue in the console
func (Queue *Queue) Print() {
	fmt.Println("Printing queue")
}
