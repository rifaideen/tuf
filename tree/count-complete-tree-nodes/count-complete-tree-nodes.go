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
func countNodes(root *TreeNode) int {
	// handle edge case
	if root == nil {
		return 0
	}

	lh := leftHeight(root)
	rh := rightHeight(root)

	// when left and right heights are same, we find 2^height - 1
	// when left and right nodes has same height, it represents the complete binary tree where left and right are equally filled
	if lh == rh {
		return (1 << lh) - 1
	}

	// when heights of left and right nodes are not same, we apply the formula
	// 1 + count left + count right
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

// find left height
func leftHeight(node *TreeNode) int {
	height := 0

	current := node

	for current != nil {
		height++
		current = current.Left
	}

	return height
}

// find right height
func rightHeight(node *TreeNode) int {
	height := 0

	current := node

	for current != nil {
		height++
		current = current.Right
	}

	return height
}
