package main

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NodeInfo struct {
	Node *TreeNode
	Line int
}

func topView(root *TreeNode) []int {
	ans := []int{}

	// handle edge case
	if root == nil {
		return ans
	}

	// create queue data-structure and start with root node
	// queue stores node and the line number
	queue := []*NodeInfo{}

	queue = append(queue, &NodeInfo{
		Node: root,
		Line: 0,
	})

	// to keep track of processed nodes and to avoid overwriting them
	// we require this because we are interested in top view, so same line should not overwrite it
	visited := map[int]int{}

	for len(queue) > 0 {
		// pop the queue
		top := queue[0]
		queue = queue[1:]

		// explicitly assign node and line
		node, line := top.Node, top.Line

		// make sure we are not visited the line already
		if _, ok := visited[line]; !ok {
			visited[line] = node.Val
		}

		// push the left node to the queue
		if node.Left != nil {
			queue = append(queue, &NodeInfo{
				Node: node.Left,
				Line: line - 1, // when going left node, we need to subtract the line by 1
			})
		}

		// push the right node to the queue
		if node.Right != nil {
			queue = append(queue, &NodeInfo{
				Node: node.Right,
				Line: line + 1, // when going right node, we need to add the line by 1
			})
		}
	}

	lines := []int{}
	for k := range visited {
		lines = append(lines, k)
	}

	sort.Ints(lines)

	// pop the lines in sorted order and take the values from the visited map
	for _, line := range lines {
		ans = append(ans, visited[line])
	}

	return ans
}
