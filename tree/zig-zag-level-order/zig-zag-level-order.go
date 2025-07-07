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
func ZigzagLevelOrder(root *TreeNode) [][]int {
	ans := [][]int{}

	// handle edge case
	if root == nil {
		return ans
	}

	// stack data structure to hold the nodes to be processed
	stack := []*TreeNode{
		root,
	}

	// determines to traversal order, starts with left to right direction
	ltr := true

	for len(stack) > 0 {
		// current stack size and loop upto n size
		n := len(stack)

		// keep the values at current level
		levels := []int{}
		// list of nodes to be process into the stack on next iteration
		next := []*TreeNode{}

		// left to right
		for range n {
			// pop the stack
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// push the node value to levels
			levels = append(levels, node.Val)

			// is left to right, left node goes first
			if ltr && node.Left != nil {
				next = append(next, node.Left)
			}

			if node.Right != nil {
				next = append(next, node.Right)
			}

			// is right to left, left node goes last
			if !ltr && node.Left != nil {
				next = append(next, node.Left)
			}
		}

		// toggle the direction
		ltr = !ltr

		// push the levels to answer
		ans = append(ans, levels)
		// push the next nodes to be processed onto the stack
		stack = append(stack, next...)
	}

	return ans
}
