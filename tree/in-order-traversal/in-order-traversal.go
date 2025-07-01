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
func inorderTraversal(root *TreeNode) []int {
	return traverse(root)
}

func traverse(root *TreeNode) []int {
	// stop when root becomes empty
	if root == nil {
		return []int{}
	}

	ans := []int{}

	// process and append values on the left recursively
	ans = append(ans, traverse(root.Left)...)
	// append root value
	ans = append(ans, root.Val)
	// process and append values on the right recursively
	ans = append(ans, traverse(root.Right)...)

	return ans
}
