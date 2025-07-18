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
func InorderTraversal(root *TreeNode) []int {
	// using morris traversal algorithm
	// We create the thread to the right most node on the left sub-tree to point to current root node
	// Once we are at the right most node, we traverse back to the current root node and remove the thread after we travelled
	ans := []int{}

	// handle edge case
	if root == nil {
		return ans
	}

	current := root

	for current != nil {
		// current doesn't have left node
		if current.Left == nil {
			// push it to answer and move right
			ans = append(ans, current.Val)
			current = current.Right
		} else { // current has left node
			// before moving left, we must create thread on the right most node on the left child of current node
			prev := current.Left

			// move right, until we have right node and it doesn't point to current node
			for prev.Right != nil && prev.Right != current {
				prev = prev.Right
			}

			if prev.Right == nil { // previous right most node doesn't point to current node
				// create thread to current node and move current node to left
				prev.Right = current

				// we have createed the thread, we can move to left node now
				current = current.Left
			} else { // previous right most node points to current node
				// erase the thread, push current to answer and move to right
				prev.Right = nil
				// we are at current node possibly traverssed via thread, add the current to answer
				ans = append(ans, current.Val)
				// move to right side
				current = current.Right
			}
		}
	}

	return ans
}
