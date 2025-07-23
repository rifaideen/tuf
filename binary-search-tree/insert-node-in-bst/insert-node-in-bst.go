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
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	// Base case: If the current node is nil, we've reached the spot where we should insert the new value.
	if root == nil {
		// Create a new node with the given value and return it.
		// This node becomes the child of the previous recursive call's parent.
		return &TreeNode{
			Val: val,
		}
	}

	// Recursive case: Decide whether to go left or right based on BST properties.
	if val < root.Val {
		// The value belongs in the left subtree.
		// Recursively insert into the left subtree. The result (newly created node)
		// is assigned back to root.Left to link the new node to the tree.
		node := insertIntoBST(root.Left, val)
		root.Left = node

		// Return the current node (root) to maintain the tree structure.
		// This ensures that parent nodes correctly retain their left/right pointers.
		return root
	} else {
		// The value belongs in the right subtree.
		// Recursively insert into the right subtree. The result (newly created node)
		// is assigned back to root.Right to link the new node to the tree.
		node := insertIntoBST(root.Right, val)
		root.Right = node

		// Return the current node (root) to maintain the tree structure.
		return root
	}
}
