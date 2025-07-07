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
func DiameterOfBinaryTree(root *TreeNode) int {
	diameter := 0
	height(root, &diameter)

	return diameter
}

func height(root *TreeNode, diameter *int) int {
	if root == nil {
		return 0
	}

	left := height(root.Left, diameter)
	right := height(root.Right, diameter)

	*diameter = max(*diameter, left+right)

	return max(left, right) + 1
}
