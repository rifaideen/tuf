package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */
// bstFromPreorder constructs a Binary Search Tree (BST) from a given preorder traversal.
//
// Intuition:
// - In a preorder traversal: [root, left subtree, right subtree]
// - The first element is always the root.
// - In a BST, all values in the left subtree are less than the root,
//   and all values in the right subtree are greater than the root.
//
// We reconstruct the tree by processing elements in order,
// using the BST property and maintaining an upper bound for each subtree.
//
// Why an upper bound?
// - It defines the "ceiling" for valid values in the current subtree.
// - When we go into the left subtree, the current node's value becomes the upper bound.
// - When we go into the right subtree, the original upper bound is preserved.
//
// We use a pointer to the index to ensure we process each element exactly once.
func bstFromPreorder(preorder []int) *TreeNode {
	i := 0 // Index to track current position in preorder array
	return traverse(preorder, &i, math.MaxInt)
}

// traverse recursively builds the BST from the preorder list.
//
// Parameters:
// - preorder: the input preorder traversal
// - i: pointer to the current index in the preorder array (shared across recursion)
// - upperBound: the maximum allowed value for the current subtree (based on BST property)
//
// Returns:
// - A pointer to the root of the constructed subtree
//
// Key Insight:
//   - If the current value exceeds the upper bound, it doesn't belong in this subtree.
//     So we return nil and let the caller handle it (e.g., move to right subtree).
//   - We only consume a value (increment i) if it fits within the bounds.
func traverse(preorder []int, i *int, upperBound int) *TreeNode {
	// Base case: if we've processed all elements or current value exceeds the bound,
	// this value doesn't belong in the current subtree.
	if *i >= len(preorder) || preorder[*i] > upperBound {
		return nil
	}

	// Current value is valid — consume it and create a node
	val := preorder[*i]
	(*i)++ // Move to next element (important: update shared index)

	// Recursively build left and right subtrees with updated bounds:
	// - Left subtree: all values must be < val (so val is the new upper bound)
	// - Right subtree: values must be < original upper bound, but > val (implied by BST)
	root := &TreeNode{
		Val:   val,
		Left:  traverse(preorder, i, val),        // Left child: values < val
		Right: traverse(preorder, i, upperBound), // Right child: values < upperBound (but > val by preorder & BST)
	}

	return root
}
