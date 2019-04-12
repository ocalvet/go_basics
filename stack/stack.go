package stack

// Stack datastructure
type Stack struct {
	top *node
}

type node struct {
	prev *node
	v    interface{}
}

func New() *Stack {
	return &Stack{}
}

func (stack *Stack) Push(v interface{}) {
	if stack.top == nil {
		stack.top = &node{nil, v}
	}
}

func (stack *Stack) Pop() interface{} {
	return stack.top.v
}

func (stack *Stack) Print() {
	
}
