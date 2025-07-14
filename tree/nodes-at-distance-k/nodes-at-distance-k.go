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
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	// handle edge case
	if root == nil {
		return []int{}
	}

	// create map to store node and it's parent map
	parents := map[int]*TreeNode{}

	// traverse the nodes and create pointers
	traverse(root, &parents)

	// start with queue data-structure
	queue := []*TreeNode{
		target,
	}

	// keep track of visited nodes to avoid re-visit
	visited := map[*TreeNode]bool{
		target: true,
	}

	// start with current distance 0
	distance := 0

	// loop the queue and process the nodes
	for len(queue) > 0 {

		n := len(queue)

		// stop when reached k distance or node becomes null
		if distance == k {
			break
		}

		// increase the distance
		distance++

		for range n {
			// pop the queue
			node := queue[0]
			queue = queue[1:]

			// move up when node has parent and not visited already
			if parent, ok := parents[node.Val]; ok && !visited[parent] {
				// mark then visited and push it to queue
				visited[parent] = true
				queue = append(queue, parent)
			}

			// move left when node has left and not visited already
			if node.Left != nil && !visited[node.Left] {
				// mark then visited and push it to queue
				visited[node.Left] = true
				queue = append(queue, node.Left)
			}

			// move right when node has right and not visited already
			if node.Right != nil && !visited[node.Right] {
				// mark then visited and push it to queue
				visited[node.Right] = true
				queue = append(queue, node.Right)
			}
		}
	}

	// Loop terminated, whatever is inside the queue is the answer
	// create ans array of length of queue
	ans := make([]int, len(queue))

	// push the values from queue to ans
	for i, q := range queue {
		ans[i] = q.Val
	}

	return ans
}

func traverse(node *TreeNode, parents *map[int]*TreeNode) {
	if node == nil {
		return
	}

	if node.Left != nil {
		(*parents)[node.Left.Val] = node
	}

	if node.Right != nil {
		(*parents)[node.Right.Val] = node
	}

	traverse(node.Left, parents)
	traverse(node.Right, parents)
}
