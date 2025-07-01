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
func levelOrder(root *TreeNode) [][]int {
	// create array of array to store the answers
	ans := [][]int{}

	// when root is empty, stop the process
	if root == nil {
		return ans
	}

	// create queue data-structure to store the tree node and push the root first
	queue := []*TreeNode{
		root,
	}

	// after pushing the root, we start the loop until queue becomes empty
	for len(queue) > 0 {
		// count the current size of the queue and start the loop upto n times
		n := len(queue)

		// create new variable to store the all the values in the current level (from left to right direction)
		levels := []int{}

		// start the loop upto n times
		for range n {
			// pop the queue and keep the queue upto date
			node := queue[0]
			// element removed, update the queue
			queue = queue[1:]

			// store the node's value to the levels
			levels = append(levels, node.Val)

			// check if we have left node and push the left node back to the queue
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			// check if we have right node and push the right node back to the queue
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// processing the nodes at level wise is done, push the collected values to the answer
		ans = append(ans, levels)
	}

	// processing is done, return the answer
	return ans
}
