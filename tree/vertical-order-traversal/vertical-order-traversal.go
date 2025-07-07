package main

import (
	"cmp"
	"slices"
)

type NodeInfo struct {
	Val int
	Row int
	Col int
}

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
func VerticalTraversal(root *TreeNode) [][]int {
	ans := [][]int{}

	if root == nil {
		return ans
	}

	nodes := []*NodeInfo{}

	// dfs search, using pre-order traversal
	var dfs func(node *TreeNode, row, col int)

	dfs = func(node *TreeNode, row, col int) {
		if node == nil {
			return
		}

		nodes = append(nodes, &NodeInfo{
			Val: node.Val,
			Row: row,
			Col: col,
		})

		dfs(node.Left, row+1, col-1)
		dfs(node.Right, row+1, col+1)
	}

	// step 1: collect nodes
	dfs(root, 0, 0)

	// sort the nodes by column, then level, then value
	slices.SortFunc(nodes, func(a, b *NodeInfo) int {
		// sort when cols are not same
		if a.Col != b.Col {
			return cmp.Compare(a.Col, b.Col)
		}

		// sort when rows are not same
		if a.Row != b.Row {
			return cmp.Compare(a.Row, b.Row)
		}

		// sort when values are not same
		return cmp.Compare(a.Val, b.Val)
	})

	// extract the nodes by group
	if len(nodes) > 0 {
		// since the nodes[0] always have the node with lowest column, we pick it from there
		currentCol := nodes[0].Col
		currentGroup := []int{}

		for _, node := range nodes {
			// when node's column and current column doesn't match, it means we have done processing the current column
			// reset the current column and current group
			if node.Col != currentCol { // current column processed
				ans = append(ans, currentGroup)
				currentGroup = []int{node.Val}
				currentCol = node.Col
			} else { // node's column and current column match, push it to current group
				currentGroup = append(currentGroup, node.Val)
			}
		}

		// add the last group, which wasn't added in the above iteration
		ans = append(ans, currentGroup)
	}

	return ans
}
