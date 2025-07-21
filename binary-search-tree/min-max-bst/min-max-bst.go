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

	// recursively find the min and max value of BST
	return minVal(root, root.Val), maxVal(root, root.Val)
}

func findMinMaxIterative(root *TreeNode) (int, int) {
	if root == nil {
		return -1, -1
	}

	// start with root value being min and max
	min := root.Val
	max := root.Val

	// start the current node from root's left node
	current := root.Left

	// move all the way to left
	for current != nil {
		// update min value to current value and move left
		min = current.Val
		current = current.Left
	}

	// start the current node from root's right node
	current = root.Right

	// move all the way to right
	for current != nil {
		// update min value to current value and move right
		max = current.Val
		current = current.Right
	}

	return min, max
}

func minVal(root *TreeNode, min int) int {
	// when root becomes null, whatever is min currently, return it
	if root == nil {
		return min
	}

	// mowe towards left with current root value being the min value
	return minVal(root.Left, root.Val)
}

func maxVal(root *TreeNode, max int) int {
	// when root becomes null, whatever is max currently, return it
	if root == nil {
		return max
	}

	// mowe towards right with current root value being the max value
	return maxVal(root.Right, root.Val)
}
