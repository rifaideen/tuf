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

// In-Order traversal using stack data structure, without recursion
func InorderTraversal2(root *TreeNode) []int {
	ans := []int{}

	// handle edge case
	if root == nil {
		return ans
	}

	// create stack to store the nodes temporarily
	stack := []*TreeNode{}

	// current node we process, we start with root node
	current := root

	// start the loop until we process all the nodes
	for {
		// when current is not empty, we can visit the left nodes, until we get null
		if current != nil {
			// push the current node to the stack and move left
			stack = append(stack, current)
			// move the current node to the left
			current = current.Left
		} else {
			// current became empty, check if the stack also empty and break it immediately
			if len(stack) == 0 {
				break
			}

			// pop the stack and assign it to current
			current = stack[len(stack)-1]
			// remove the popped element from stack
			stack = stack[:len(stack)-1]

			// push the current value to the answer
			ans = append(ans, current.Val)

			// move the node to the right.
			// it goes to the if block because we assigned current node and it then process all the left nodes for this
			current = current.Right
		}
	}

	return ans
}
