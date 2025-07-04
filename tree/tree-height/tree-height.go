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
func MaxDepth(root *TreeNode) int {
	// handle edge case when root becomes null
	if root == nil {
		return 0
	}

	// find the left node's height recursively
	left := MaxDepth(root.Left)
	// find the right node's height recursively
	right := MaxDepth(root.Right)

	// choose the maximum height between left and right and increment by 1 (root node inclusive)
	return max(left, right) + 1
}
