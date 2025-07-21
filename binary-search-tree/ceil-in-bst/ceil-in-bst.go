package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// findCeil returns the ceiling node of value k in a binary search tree (BST).
// The ceiling is defined as the smallest node value that is greater than or equal to k.
func findCeil(root *TreeNode, k int) *TreeNode {
	// Base case: If the tree is empty, there's no ceiling.
	if root == nil {
		return nil
	}

	// If the current node's value matches k exactly, it's the best possible ceiling.
	if root.Val == k {
		return root
	}

	// If the current node's value is greater than k, it could be a potential ceiling.
	// But we first explore the left subtree to see if there's a smaller value
	// that still satisfies being >= k.
	if root.Val > k {
		leftResult := findCeil(root.Left, k)

		// If the left subtree returns a valid ceiling (not nil),
		// that node is a better (smaller) ceiling than the current root.
		if leftResult != nil {
			return leftResult
		}

		// If no better candidate is found in the left subtree,
		// the current node is the smallest valid ceiling found so far.
		return root
	}

	// If the current node's value is less than k, the ceiling must be in the right subtree.
	// We can safely ignore the left subtree because any value there would also be less than k.
	return findCeil(root.Right, k)
}
