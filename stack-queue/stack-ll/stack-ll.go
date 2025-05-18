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
