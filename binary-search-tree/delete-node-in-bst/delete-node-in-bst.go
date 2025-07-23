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
func deleteNode(root *TreeNode, key int) *TreeNode {
	// If the current node is nil, the key isn't found, so return nil.
	// This is the base case for recursion.
	if root == nil {
		return nil
	}

	// If we find the node to delete
	if root.Val == key {
		// Case 1: Node has no left child.
		// We can simply replace this node with its right child (which may be nil).
		if root.Left == nil {
			return root.Right
		}

		// Case 2: Node has a left child.
		// We want to preserve the BST property. One approach is to attach the right subtree
		// to the rightmost (largest) node in the left subtree, then return the left subtree
		// as the new root of this subtree.

		// Find the rightmost node in the left subtree
		last := root.Left // 'last' will track the rightmost node
		current := root.Left

		// Traverse to the rightmost node in the left subtree
		for current != nil {
			last = current
			current = current.Right
		}

		// Attach the right subtree of the node to be deleted to the right of 'last'
		// This ensures all values in the right subtree are still greater than all in the left
		last.Right = root.Right

		// Return the left subtree as the new root, effectively removing the current node
		return root.Left
	}

	// If the key is smaller than the current node's value,
	// it must be in the left subtree. Recursively delete it and update the left child.
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)

		return root
	}

	// If the key is greater than the current node's value,
	// it must be in the right subtree. Recursively delete it and update the right child.
	root.Right = deleteNode(root.Right, key)
	return root
}
