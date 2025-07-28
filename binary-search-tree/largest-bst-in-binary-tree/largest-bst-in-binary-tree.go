package main

import "math"

// NodeInfo holds information about a subtree needed to determine if it's a valid BST
// and to compute the size of the largest BST subtree.
type NodeInfo struct {
	max     int // Maximum value in the subtree
	min     int // Minimum value in the subtree
	maxSize int // Size of the largest BST in this subtree
}

// New creates a new NodeInfo instance.
// Used to return sentinel values for nil nodes or invalid BSTs.
func New(max, min, maxSize int) *NodeInfo {
	return &NodeInfo{
		max:     max,
		min:     min,
		maxSize: maxSize,
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// largestBst returns the size of the largest Binary Search Tree (BST) subtree
// within the given binary tree.
func largestBst(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return largestBstHelper(root).maxSize
}

// largestBstHelper traverses the tree in post-order and returns a NodeInfo
// containing:
//   - The max and min values in the current subtree (to validate BST property)
//   - The size of the largest valid BST in the subtree
//
// Intuition:
// We use a bottom-up approach. For each node, we assume we already know the BST properties
// of its left and right subtrees. Then we check:
//  1. Is the current subtree a valid BST?
//     - All values in the left subtree must be < root.Val
//     - All values in the right subtree must be > root.Val
//     - Both left and right subtrees must themselves be BSTs.
//  2. If valid, the size is 1 + size(left BST) + size(right BST)
//  3. If not valid, we return the maximum BST size found in either subtree.
//
// We use extreme values (MaxInt, MinInt) to represent invalid BSTs or nil nodes,
// so they don't interfere with comparisons.
func largestBstHelper(root *TreeNode) *NodeInfo {
	// Base case: empty tree
	// Return extreme values so that any parent node can still potentially be a valid BST.
	// maxSize = 0 because no nodes exist.
	if root == nil {
		return New(math.MaxInt, math.MinInt, 0) // Note: max < min indicates empty
	}

	// Recursively get info from left and right subtrees
	left := largestBstHelper(root.Left)
	right := largestBstHelper(root.Right)

	// Check if the current node can form a valid BST with its children
	// A valid BST requires:
	//   - All values in the left subtree < root.Val
	//   - All values in the right subtree > root.Val
	//   - Both left and right subtrees are valid BSTs (implied by our recursive structure)
	//
	// Note: We compare root.Val against left.max and right.min
	if left.max < root.Val && root.Val < right.min {
		// Current subtree is a valid BST
		// Update the min (smallest in left or root) and max (largest in right or root)
		// Size is 1 (current node) + sizes of valid left and right BSTs
		return New(
			max(root.Val, right.max),     // new max
			min(root.Val, left.min),      // new min
			1+left.maxSize+right.maxSize, // new size
		)
	}

	// Current subtree is NOT a valid BST
	// So we cannot include it in a larger BST
	// Return the largest BST size found in either left or right subtree
	// Also return extreme values for max/min to ensure parent nodes won't
	// mistakenly treat this as a valid BST.
	return New(
		math.MaxInt,                      // Invalid max (won't satisfy < condition)
		math.MinInt,                      // Invalid min (won't satisfy > condition)
		max(left.maxSize, right.maxSize), // Best BST size from children
	)
}
