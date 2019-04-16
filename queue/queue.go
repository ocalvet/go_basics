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

func New() *Queue {
	return &Queue{}
}

func (Queue *Queue) Enqueue(v interface{}) {
	if Queue.root == nil {
		Queue.root = &node{nil, v}
	}
}

func (Queue *Queue) Dequeue() interface{} {
	return Queue.root.v
}

func (Queue *Queue) Print() {
	fmt.Println("Printing queue")
}
