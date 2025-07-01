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
func preorderTraversal(root *TreeNode) []int {
	return traverse(root)
}

func traverse(node *TreeNode) []int {
	// stop processing when node is empty
	if node == nil {
		return []int{}
	}

	// add root node's value
	ans := []int{node.Val}

	// append the left node values recursively
	ans = append(ans, preorderTraversal(node.Left)...)
	// append the right node values recursively
	ans = append(ans, preorderTraversal(node.Right)...)

	return ans
}
