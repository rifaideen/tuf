package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printPath(root *TreeNode, x *TreeNode) []int {
	ans := []int{}
	printPathRecursive(root, x, &ans)

	return ans
}

func printPathRecursive(node, x *TreeNode, path *[]int) bool {
	// handle edge case, return false when node becomes null
	if node == nil {
		return false
	}

	// append the node value to the path
	*path = append(*path, node.Val)

	// node and x value is matching, return true
	if node.Val == x.Val {
		return true
	}

	// if left recursion or right recursion yield the result, return true
	if printPathRecursive(node.Left, x, path) || printPathRecursive(node.Right, x, path) {
		return true
	}

	// node not found, pop the last inserted value and return false
	*path = (*path)[:len(*path)-1]

	return false
}
