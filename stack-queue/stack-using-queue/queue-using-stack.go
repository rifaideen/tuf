package main

type MyQueue struct {
	stack1 *MyStack
	stack2 *MyStack
}

func Constructor() MyQueue {
	return MyQueue{
		stack1: NewStack(),
		stack2: NewStack(),
	}
}

func (t *MyQueue) Push(x int) {
	// push values from stack 1 to stack 2
	for !t.stack1.Empty() {
		t.stack2.Push(t.stack1.Pop())
	}

	// push new value to stack 1
	t.stack1.Push(x)

	// put back values from stack 2 into stack 1
	for !t.stack2.Empty() {
		t.stack1.Push(t.stack2.Pop())
	}
}

func (t *MyQueue) Pop() int {
	return t.stack1.Pop()
}

func (t *MyQueue) Peek() int {
	return t.stack1.Top()
}

func (t *MyQueue) Empty() bool {
	return t.stack1.Empty()
}

type MyStack struct {
	top    int
	values []int
}

func NewStack() *MyStack {
	return &MyStack{
		top:    -1,
		values: []int{},
	}
}

func (t *MyStack) Push(x int) {
	t.top += 1
	t.values = append(t.values, x)
}

func (t *MyStack) Pop() int {
	if t.top < 0 {
		panic("stack underflow")
	}

	value := t.values[t.top]
	t.values = t.values[:t.top]
	t.top -= 1

	return value
}

func (t *MyStack) Top() int {
	if t.top == -1 {
		panic("stack is empty")
	}

	return t.values[t.top]
}

func (t *MyStack) Empty() bool {
	return t.top == -1
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
