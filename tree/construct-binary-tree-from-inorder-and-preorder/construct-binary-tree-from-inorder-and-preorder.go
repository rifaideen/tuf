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

// buildTree constructs a binary tree from given preorder and inorder traversal slices.
func BuildTree(preorder []int, inorder []int) *TreeNode {
	// Create a map to store value-to-index mapping of inorder traversal
	inorderValueToIndex := make(map[int]int)
	n := len(inorder)

	for i := 0; i < n; i++ {
		inorderValueToIndex[inorder[i]] = i
	}

	// Start building the tree recursively using both traversals
	return buildSubtree(
		inorder, 0, n-1,
		preorder, 0, n-1,
		inorderValueToIndex,
	)
}

// buildSubtree recursively builds a subtree using:
// - inorder: slice containing inorder traversal
// - inStart, inEnd: start and end indices in the inorder slice
// - preorder: slice containing preorder traversal
// - preStart, preEnd: start and end indices in the preorder slice
// - inorderValueToIndex: map for quick lookup of root index in inorder
func buildSubtree(
	inorder []int, inStart, inEnd int,
	preorder []int, preStart, preEnd int,
	inorderValueToIndex map[int]int,
) *TreeNode {

	// Base case: If there are no elements to construct the subtree
	if inStart > inEnd || preStart > preEnd {
		return nil
	}

	// The first element in preorder is always the root of the current subtree
	rootVal := preorder[preStart]
	root := &TreeNode{Val: rootVal}

	// Find the position of the root in the inorder traversal
	inRoot := inorderValueToIndex[rootVal]

	// Calculate how many nodes are on the left side of the root in the inorder traversal
	leftSubtreeSize := inRoot - inStart

	// Recursively build the left subtree:
	// - Inorder range: from inStart to inRoot - 1
	// - Preorder range: from preStart + 1 to preStart + leftSubtreeSize
	root.Left = buildSubtree(
		inorder, inStart, inRoot-1,
		preorder, preStart+1, preStart+leftSubtreeSize,
		inorderValueToIndex,
	)

	// Recursively build the right subtree:
	// - Inorder range: from inRoot + 1 to inEnd
	// - Preorder range: from preStart + leftSubtreeSize + 1 to preEnd
	root.Right = buildSubtree(
		inorder, inRoot+1, inEnd,
		preorder, preStart+leftSubtreeSize+1, preEnd,
		inorderValueToIndex,
	)

	return root
}
