package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printBoundary(root *TreeNode) []int {
	ans := []int{}

	// handle edge case
	if root == nil {
		return ans
	}

	// store the root value to the answer
	ans = append(ans, root.Val)

	// We split the tree into left, leaf and right boundary

	// left boundary traverse all the way to the left excluding the leaf nodes
	// anytime when left node is null, we switch to the right node
	addLeftBoundary(root, &ans)
	// traverse the nodes and add only the leaf nodes to the answers
	addLeaves(root, &ans)
	// right boundary traverse all the way to the right excluding the leaf nodes
	// anytime when left node is null, we switch to the right node
	addRightBoundary(root, &ans)

	return ans
}

// check if the node is leaf node
func isLeaf(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

// traverse left nodes and use right node when left node becomes null and add the values
// in the answer array, excluding leaf nodes
func addLeftBoundary(root *TreeNode, ans *[]int) {
	current := root.Left

	for current != nil {
		// add only if current is not leaf node
		if !isLeaf(current) {
			*ans = append(*ans, current.Val)
		}

		// move the current to left
		if current.Left != nil {
			current = current.Left
		} else { // left node is empty, move to right node
			current = current.Right
		}
	}
}

func addLeaves(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}

	if isLeaf(root) {
		*ans = append(*ans, root.Val)
		return
	}

	addLeaves(root.Left, ans)
	addLeaves(root.Right, ans)
}

// traverse right nodes and use left node when right node becomes null and add the values
// in temporary array, excluding leaf nodes
// once the iteration is done, loop the temp array in reverse order and push it to answer array
func addRightBoundary(root *TreeNode, ans *[]int) {
	current := root.Right

	temp := []int{}

	for current != nil {
		if !isLeaf(current) {
			temp = append(temp, current.Val)
		}

		if current.Right != nil {
			current = current.Right
		} else {
			current = current.Left
		}
	}

	for i := len(temp) - 1; i >= 0; i-- {
		*ans = append(*ans, temp[i])
	}
}
