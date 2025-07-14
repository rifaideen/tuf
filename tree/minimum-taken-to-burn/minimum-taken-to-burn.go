package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minTimeToBurn(root *TreeNode, target *TreeNode) int {
	// handle edge case
	if root == nil {
		return 0
	}

	ans := 0

	parents := map[int]*TreeNode{}

	// create node and parent pointers to lookup the parent nodes of any node
	traverse(root, &parents)

	// queue data-structure to burn the nodes
	queue := []*TreeNode{
		target, // start with target node
	}

	// keep track of visited nodes to avoid revisiting
	visited := map[*TreeNode]bool{
		target: true, // target is visited
	}

	for len(queue) > 0 {
		n := len(queue)

		// keep a flag to see if we have burnt any nodes or not
		burned := false

		for range n {
			node := queue[0]
			queue = queue[1:]

			// move up when not visited already
			if parent, ok := parents[node.Val]; ok && !visited[parent] {
				visited[parent] = true
				queue = append(queue, parent)
				burned = true
			}

			// move left when not visited already
			if node.Left != nil && !visited[node.Left] {
				visited[node.Left] = true
				queue = append(queue, node.Left)
				burned = true
			}

			// move right when not visited already
			if node.Right != nil && !visited[node.Right] {
				visited[node.Right] = true
				queue = append(queue, node.Right)
				burned = true
			}
		}

		// increase the answer only when nodes are burned
		if burned {
			ans++
		}
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
