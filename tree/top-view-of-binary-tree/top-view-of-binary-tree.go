package main

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
	visited := map[int]struct{}{}

	// we need to maintain the result in sorted manner, hence we sort the result with line
	h := &MinHeap{
		values: [][]int{},
		size:   0,
	}

	for len(queue) > 0 {
		// pop the queue
		top := queue[0]
		queue = queue[1:]

		// explicitly assign node and line
		node, line := top.Node, top.Line

		// make sure we are not visited the line already
		if _, ok := visited[line]; !ok {
			visited[line] = struct{}{}
			h.Push(node, line)
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

	// pop the values from stack and return
	for h.size > 0 {
		ans = append(ans, h.Pop()[0])
	}

	return ans
}

type MinHeap struct {
	values [][]int // [value, index] pair
	size   int
}

func (h *MinHeap) Push(node *TreeNode, line int) {
	h.values = append(h.values, []int{node.Val, line})
	h.Up(h.size)

	h.size++
}

func (h *MinHeap) Pop() []int {
	top := h.values[0]

	h.size--
	h.values[0] = h.values[h.size]
	h.values = h.values[:h.size]

	h.Down(0)

	return top
}

func (h *MinHeap) Up(index int) {
	if index > 0 {
		parent := (index - 1) / 2

		if h.values[index][1] < h.values[parent][1] {
			h.Swap(index, parent)

			h.Up(parent)
		}
	}
}

func (h *MinHeap) Down(index int) {
	left := (2 * index) + 1
	right := left + 1

	smallest := index

	if left < h.size && h.values[left][1] < h.values[smallest][1] {
		smallest = left
	}

	if right < h.size && h.values[right][1] < h.values[smallest][1] {
		smallest = right
	}

	if index != smallest {
		h.Swap(index, smallest)
		h.Down(smallest)
	}
}

func (h *MinHeap) Swap(i, j int) {
	h.values[i], h.values[j] = h.values[j], h.values[i]
}
