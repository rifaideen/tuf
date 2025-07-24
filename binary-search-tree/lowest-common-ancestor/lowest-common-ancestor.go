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
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// Base case: if the current node is nil, we've reached a leaf's child,
	// so there's no ancestor here.
	if root == nil {
		return nil
	}

	// In a BST, all nodes in the left subtree have values less than the root.
	// If both p and q are less than the current root, then both must be in the left subtree.
	// Therefore, the LCA must also be in the left subtree.
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else if p.Val > root.Val && q.Val > root.Val {
		// Similarly, if both p and q are greater than the current root,
		// they must both be in the right subtree, so the LCA is in the right subtree.
		return lowestCommonAncestor(root.Right, p, q)
	}

	// If we reach here, it means that p and q are on different sides of the current root
	// (or one of them is the root itself). This implies that the current root
	// is the lowest common ancestor because we're at the "split" point.
	//
	// Why?
	// - If one node is on the left and the other on the right,
	//   their paths diverge at this root.
	// - If one of the nodes is the current root (e.g., p.Val == root.Val),
	//   then the root is the first common ancestor (since we need the lowest).
	//
	// So, we return the current root as the LCA.
	return root
}
