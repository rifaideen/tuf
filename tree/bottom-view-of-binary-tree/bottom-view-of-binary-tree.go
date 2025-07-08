package main

import (
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NodeInfo struct {
	Node *TreeNode
	Line int
}

func bottomView(root *TreeNode) []int {
	ans := []int{}

	// handle edge-case
	if root == nil {
		return ans
	}

	// create queue to list of nodes to be processed
	queue := []*NodeInfo{}
	// start with root
	queue = append(queue, &NodeInfo{
		Node: root,
		Line: 0,
	})

	// keep track of visitied line and it's value,
	// each time when same line visited, it would be replaced by latest value
	visited := map[int]int{} // line: value pair

	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]

		node, line := top.Node, top.Line

		// update visited line with current node's value
		visited[line] = node.Val

		// We solve this by forming imaginary line, where root is at line 0, left nodes are -1,-2,etc
		// Right nodes are +1,+2,+3 etc
		// push left node to queue
		if node.Left != nil {
			queue = append(queue, &NodeInfo{
				Node: node.Left,
				Line: line - 1, // -1 subtracted from line for left node
			})
		}

		// push right node to queue
		if node.Right != nil {
			queue = append(queue, &NodeInfo{
				Node: node.Right,
				Line: line + 1, // +1 add from line for right node
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
