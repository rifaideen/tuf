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
func searchBST(root *TreeNode, val int) *TreeNode {
	// hande base case, when root becomes null
	if root == nil {
		return nil
	}

	// if root values matches the lookup value
	if val == root.Val {
		return root
	}

	// if lookup value is less than root value, it must be on root left
	if val < root.Val {
		return searchBST(root.Left, val)
	}

	// lookup value is greater than root value, it must be on root right
	return searchBST(root.Right, val)
}
