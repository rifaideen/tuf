package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMinMax(root *TreeNode) (int, int) {
	if root == nil {
		return -1, -1
	}

	return minVal(root, root.Val), maxVal(root, root.Val)
}

func minVal(root *TreeNode, min int) int {
	if root == nil {
		return min
	}

	return minVal(root.Left, root.Val)
}

func maxVal(root *TreeNode, min int) int {
	if root == nil {
		return min
	}

	return maxVal(root.Right, root.Val)
}
