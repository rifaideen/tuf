package main

import "math"

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
func IsBalanced(root *TreeNode) bool {
	return maxDepth(root) != -1
}

func maxDepth(root *TreeNode) int {
	// handle edge case when root becomes null
	if root == nil {
		return 0
	}

	// find the left node's height recursively
	left := maxDepth(root.Left)
	// find the right node's height recursively
	right := maxDepth(root.Right)

	// in any case we receive -1, we return -1 as well
	if left == -1 || right == -1 {
		return -1
	}

	// instead of just calculating the height, we return -1 when
	// absolute difference between left and right > 1
	if math.Abs(float64(left-right)) > 1 {
		return -1
	}

	// choose the maximum height between left and right and increment by 1 (root node inclusive)
	return max(left, right) + 1
}
