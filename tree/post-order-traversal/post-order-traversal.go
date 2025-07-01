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
func postorderTraversal(root *TreeNode) []int {
	return traverse(root)
}

func traverse(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ans := []int{}

	ans = append(ans, traverse(root.Left)...)
	ans = append(ans, traverse(root.Right)...)
	ans = append(ans, root.Val)

	return ans
}
