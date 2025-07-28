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
func findTarget(root *TreeNode, k int) bool {
	iterator := NewIterator(root)

	left := iterator.Next()
	right := iterator.Before()

	for left < right {
		if left+right == k {
			return true
		}

		if left+right > k {
			right = iterator.Before()
		} else {
			left = iterator.Next()
		}
	}

	return false
}

type Iterator struct {
	lstack []*TreeNode
	rstack []*TreeNode
}

func NewIterator(root *TreeNode) *Iterator {
	i := &Iterator{
		lstack: []*TreeNode{},
		rstack: []*TreeNode{},
	}

	i.PushLeft(root)
	i.PushRight(root)

	return i
}

func (i *Iterator) PushLeft(node *TreeNode) {
	for node != nil {
		i.lstack = append(i.lstack, node)
		node = node.Left
	}
}

func (i *Iterator) PushRight(node *TreeNode) {
	for node != nil {
		i.rstack = append(i.rstack, node)
		node = node.Right
	}
}

func (i *Iterator) Next() int {
	node := i.lstack[len(i.lstack)-1]
	i.lstack = i.lstack[:len(i.lstack)-1]
	i.PushLeft(node.Right)

	return node.Val
}

func (i *Iterator) Before() int {
	node := i.rstack[len(i.rstack)-1]
	i.rstack = i.rstack[:len(i.rstack)-1]
	i.PushRight(node.Left)

	return node.Val
}
