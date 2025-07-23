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
func isValidBST(root *TreeNode) bool {
	return isValid(root, math.MinInt, math.MaxInt)
}

func isValid(root *TreeNode, minVal, maxVal int) bool {
	// handle edge case, when the root is null, it is considered valid
	if root == nil {
		return true
	}

	// at any moment, root value goes out of range in either case, it's considered invalid
	if root.Val >= maxVal || root.Val <= minVal {
		return false
	}

	// recursively verify the left and right sub-trees
	return isValid(root.Left, minVal, root.Val) && isValid(root.Right, root.Val, maxVal)
}
