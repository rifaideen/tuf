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
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// base case when root is nil or root matches either p or q, return root
	if root == nil || root == p || root == q {
		return root
	}

	// traverse left recursively
	left := LowestCommonAncestor(root.Left, p, q)
	// traverse right recursively
	right := LowestCommonAncestor(root.Right, p, q)

	// if left gets nil, return right
	if left == nil {
		return right
	}

	// if right gets nil, return left
	if right == nil {
		return left
	}

	// both left and right got nil, return root
	return root
}
