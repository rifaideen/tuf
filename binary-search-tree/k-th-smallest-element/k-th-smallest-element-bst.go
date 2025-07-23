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
func kthSmallest(root *TreeNode, k int) int {
	// Edge case: empty tree has no elements
	if root == nil {
		return 0
	}

	// ans stores the in-order traversal (ascending values)
	// stack simulates recursion for iterative traversal
	ans := []int{}
	stack := []*TreeNode{}

	current := root
	counter := 0

	// Iterative in-order traversal:
	// 1. Traverse leftmost path, pushing nodes to stack
	// 2. When left end is reached, pop and process node
	// 3. Move to right subtree and repeat
	for {
		if current != nil {
			// Reach leftmost node while storing path in stack
			stack = append(stack, current)
			current = current.Left
		} else {
			// When left path is exhausted, backtrack
			if len(stack) == 0 {
				break // Traversal complete
			}

			// Process the top node from stack
			current = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// Record node value and increment count
			ans = append(ans, current.Val)
			counter++

			// Early exit if we found the kth smallest
			if counter == k {
				break
			}

			// Explore right subtree
			current = current.Right
		}
	}

	// Since in-order traversal of BST yields sorted order,
	// the kth element in ans is the kth smallest
	return ans[k-1]
}
