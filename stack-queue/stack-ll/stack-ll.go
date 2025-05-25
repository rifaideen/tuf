package main

type Node struct {
	data int
	next *Node
}

type MyStack struct {
	top    *Node
	values *Node
}

func NewStack() *MyStack {
	return &MyStack{
		top:    nil,
		values: nil,
	}
}

func (t *MyStack) Push(x int) {
	top := &Node{
		data: x,
		next: t.top,
	}

	t.top = top
}

func (t *MyStack) Pop() int {
	data := t.top.data
	t.top = t.top.next

	return data
}

func (t *MyStack) Top() int {
	if t.top == nil {
		panic("stack is empty")
	}

	return t.top.data
}

func (t *MyStack) Empty() bool {
	return t.top == nil
}

type Stack struct {
	top    int
	values []int
}

func New() *Stack {
	return &Stack{
		top:    -1,
		values: []int{},
	}
}

func (t *Stack) Push(x int) {
	t.top++
	t.values = append(t.values, x)
}

func (t *Stack) Pop() int {
	data := t.values[t.top]
	t.values = t.values[:t.top]
	t.top--

	return data
}

func (t *Stack) Top() int {
	if t.top < 0 {
		panic("stack is empty")
	}

	return t.values[t.top]
}

func (t *Stack) Empty() bool {
	return t.top < 0
}
