package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findFloor(root *TreeNode, k int) *TreeNode {
	// Base case: If the tree is empty, there's no node to consider as floor.
	if root == nil {
		return nil
	}

	// If the current node's value matches k exactly, it's the best possible floor.
	if root.Val == k {
		return root
	}

	// If current node's value is less than k:
	// This node is a potential floor candidate (since floor is ≤ k),
	// but we explore the right subtree to see if there's a **larger** value
	// that still satisfies being ≤ k (to find the largest possible floor).
	if root.Val < k {
		rightResult := findFloor(root.Right, k)

		// If the right subtree returns a valid floor (not nil),
		// that node is a better (larger) floor than the current root.
		if rightResult != nil {
			return rightResult
		}

		// If no better candidate is found in the right subtree,
		// the current node is the best floor we've found so far.
		return root
	}

	// If current node's value is greater than k:
	// The floor must be in the left subtree because:
	// - Values in the left subtree are smaller than root.Val
	// - We need a value ≤ k (so root.Val is invalid as a floor)
	// We can safely ignore the right subtree here.
	return findFloor(root.Left, k)
}
