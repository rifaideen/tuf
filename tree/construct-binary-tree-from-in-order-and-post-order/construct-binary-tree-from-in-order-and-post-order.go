package main

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

// buildTree constructs a binary tree from inorder and postorder traversal slices.
func BuildTree(inorder []int, postorder []int) *TreeNode {
	// Create a map to quickly find the index of a value in the inorder traversal
	valueToInorderIndex := make(map[int]int)
	n := len(inorder)

	// Map each value to its index in the inorder array
	for i, val := range inorder {
		valueToInorderIndex[val] = i
	}

	// Start recursive construction of the binary tree
	return buildSubtree(
		inorder, 0, n-1,
		postorder, 0, n-1,
		valueToInorderIndex,
	)
}

// buildSubtree recursively builds a subtree using:
// - inorder: slice containing inorder traversal
// - inStart, inEnd: start and end indices in the inorder slice
// - postorder: slice containing postorder traversal
// - pStart, pEnd: start and end indices in the postorder slice
// - valueToInorderIndex: map for quick lookup of root index in inorder
func buildSubtree(
	inorder []int, inStart, inEnd int,
	postorder []int, pStart, pEnd int,
	valueToInorderIndex map[int]int,
) *TreeNode {

	// Base case: no elements to construct the subtree
	if inStart > inEnd || pStart > pEnd {
		return nil
	}

	// Root is always the last element in current postorder segment
	rootValue := postorder[pEnd]
	root := &TreeNode{Val: rootValue}

	// Find position of the root in the inorder traversal
	inRoot := valueToInorderIndex[rootValue]

	// Calculate size of the left subtree
	leftSubtreeSize := inRoot - inStart

	// Recursively build left subtree
	// Inorder: from inStart to inRoot - 1
	// Postorder: from pStart to pStart + leftSubtreeSize - 1
	root.Left = buildSubtree(
		inorder, inStart, inRoot-1,
		postorder, pStart, pStart+leftSubtreeSize-1,
		valueToInorderIndex,
	)

	// Recursively build right subtree
	// Inorder: from inRoot + 1 to inEnd
	// Postorder: from pStart + leftSubtreeSize to pEnd - 1
	root.Right = buildSubtree(
		inorder, inRoot+1, inEnd,
		postorder, pStart+leftSubtreeSize, pEnd-1,
		valueToInorderIndex,
	)

	return root
}
