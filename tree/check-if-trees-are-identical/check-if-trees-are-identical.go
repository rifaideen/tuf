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
func IsSameTree(p *TreeNode, q *TreeNode) bool {
	// when both nodes are nil, it's identical
	if p == nil && q == nil {
		return true
	}

	// when one of node is empty, it's not identical
	if p == nil || q == nil {
		return false
	}

	// when both values are not same, it's not identical
	if p.Val != q.Val {
		return false
	}

	// when left and right is same, it's identical
	return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}
