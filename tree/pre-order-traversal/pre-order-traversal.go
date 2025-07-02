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

// pre-order traversal using stack, without recursion
func PreorderTraversal2(root *TreeNode) []int {
	// Pre-order: Root, Left, Right
	ans := []int{}

	// handle edge case
	if root == nil {
		return ans
	}

	// create stack and start with root node
	stack := []*TreeNode{
		root,
	}

	// start the loop until stack becomes empty
	for len(stack) > 0 {
		// pop the stack
		node := stack[len(stack)-1]
		// remove the popped item from stack
		stack = stack[:len(stack)-1]

		// push the node value to the answer
		ans = append(ans, node.Val)

		// if node has right, push it to stack
		// we push the right before the left node, because stack is LIFO and if we push the left first,
		// the traversal order will be broken
		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		// if node has left, push it to stack
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return ans
}
