package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	root  *TreeNode
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	it := BSTIterator{
		root:  root,
		stack: []*TreeNode{},
	}

	it.PushLeft(root)

	return it
}

func (t *BSTIterator) Next() int {
	// pop the node
	node := t.stack[len(t.stack)-1]
	t.stack = t.stack[:len(t.stack)-1]
	t.PushLeft(node.Right)

	return node.Val
}

func (t *BSTIterator) HasNext() bool {
	return len(t.stack) > 0
}

func (t *BSTIterator) PushLeft(node *TreeNode) {
	for node != nil {
		t.stack = append(t.stack, node)
		node = node.Left
	}
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
