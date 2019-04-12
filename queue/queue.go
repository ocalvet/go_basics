package Queue

// Queue datastructure
type Queue struct {
	top *node
}

type node struct {
	prev *node
	v    interface{}
}

func New() *Queue {
	return &Queue{}
}

func (Queue *Queue) Enqueue(v interface{}) {
	if Queue.top == nil {
		Queue.top = &node{nil, v}
	}
}

func (Queue *Queue) Dequeue() interface{} {
	return Queue.top.v
}

func (Queue *Queue) Print() {

}
