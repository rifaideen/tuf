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
 *
 * Pre-order traversal in iterative approach using stack.
 * Pre-order = root, left, right
 */
func preorderTraversal(root *TreeNode) []int {
	ans := []int{}

	// handle edge case when root is nil
	if root == nil {
		return ans
	}

	// stack data-structure, initially stores the root
	stack := []*TreeNode{
		root,
	}

	// start the loop until stack becomes empty
	for len(stack) > 0 {
		// pop the stack and remove the popped item manually
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// add the popped node's value
		ans = append(ans, node.Val)

		// when right node is not null, push it to stack
		// to make sure we get the nodes in order root, left and right, we push the right node first
		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		// push it to stack when left node is not null
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return ans
}
