package main

type Node struct {
	Data int
	Next *Node
}

type MyQueue struct {
	top    *Node
	bottom *Node
}

func Constructor() MyQueue {
	return MyQueue{
		top:    nil,
		bottom: nil,
	}
}

func (q *MyQueue) Push(x int) {
	node := &Node{
		Data: x,
	}

	if q.top == nil {
		q.top, q.bottom = node, node
		return
	}

	q.bottom.Next = node
	q.bottom = q.bottom.Next
}

func (q *MyQueue) Pop() int {
	if q.top == nil {
		panic("queue is empty")
	}

	value := q.top.Data

	q.top = q.top.Next

	return value
}

func (q *MyQueue) Peek() int {
	if q.top == nil {
		panic("queue is empty")
	}

	return q.top.Data
}

func (q *MyQueue) Empty() bool {
	return q.top == nil
}
