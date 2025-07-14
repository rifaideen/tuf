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
func WidthOfBinaryTree(root *TreeNode) int {
	// We solve this problem by assigning the appropriate index to each node
	// and then comparing the right most node with left most node to find the width
	// among the widths, we choose the max width
	//
	// When there is only one node present at any level, we could not consider that level
	// for width calculation
	//
	// Since the number of nodes can be large, it could lead to overflow error while assigning the
	// index, to avoid that, we do some trick here to avoid that

	ans := 0

	// handle edge case
	if root == nil {
		return ans // 0 width when root it null
	}

	type pair struct {
		node  *TreeNode
		index int
	}

	// queue ds to store pair struct the node and index
	queue := []pair{}
	queue = append(queue, pair{
		node:  root,
		index: 0,
	})

	// while queue is not empty, start popping out
	for len(queue) > 0 {
		n := len(queue)

		// keep the minimum index in the queue before processing.
		minValue := queue[0].index

		// we use this for final comparision to find the max width
		first := 0
		last := 0

		// instead of shrinking the queue inside the below loop, which causes the time limit exception,
		// we process the queue by index i and new nodes are added to this next pair
		// ans replace queue with this next pair
		next := []pair{}

		for i := range n {
			node := queue[i]
			// calculate the current id
			id := node.index - minValue

			// set first and last node's index for comparision
			if i == 0 {
				first = id
			}

			if i == (n - 1) {
				last = id
			}

			// if node has left node, put them in pair and add the index
			// 2xi + 1, where i is the id for left
			if node.node.Left != nil {
				next = append(next, pair{
					node:  node.node.Left,
					index: (2 * id) + 1,
				})
			}

			// if node has right node, put them in pair and add the index
			// 2xi + 2, where i is the id for right
			if node.node.Right != nil {
				next = append(next, pair{
					node:  node.node.Right,
					index: (2 * id) + 2,
				})
			}
		}

		// replace queue with next pairs
		queue = next

		// keep track of max width
		ans = max(ans, last-first+1)
	}

	return ans
}
