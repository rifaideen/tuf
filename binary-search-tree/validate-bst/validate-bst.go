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
// Main function to validate if a tree is a BST
func isValidBST(root *TreeNode) bool {
	// Start validation with full integer range as initial bounds
	// All node values must stay strictly between -infinity and +infinity
	return isValid(root, math.MinInt, math.MaxInt)
}

// Helper function with recursive validation logic
func isValid(root *TreeNode, minVal, maxVal int) bool {
	// Base case: Empty subtree is always valid
	if root == nil {
		return true
	}

	// Current node check:
	// - Must be GREATER than minVal (ancestor's lower bound)
	// - Must be LESS than maxVal (ancestor's upper bound)
	// Equal values are invalid (BST requires unique elements)
	if root.Val >= maxVal || root.Val <= minVal {
		return false
	}

	// Recursive case:
	// For LEFT subtree:
	// - Keeps same lower bound (minVal)
	// - New upper bound becomes current node's value
	// For RIGHT subtree:
	// - New lower bound becomes current node's value
	// - Keeps same upper bound (maxVal)
	// Both subtrees must be valid for the whole tree to be valid
	return isValid(root.Left, minVal, root.Val) && isValid(root.Right, root.Val, maxVal)
}
