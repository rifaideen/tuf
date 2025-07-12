package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func changeTree(root *TreeNode) {
	// handle edge case
	if root == nil {
		return
	}

	// we count the immediate left and right node and if the count it greater than root's value
	// we replace the root with counted value
	// else we replace the left and right children node
	total := 0

	// make sure root has left node
	if root.Left != nil {
		total += root.Left.Val
	}

	// make sure root has right node
	if root.Right != nil {
		total += root.Right.Val
	}

	// check if total is greater or equal to total, we replace the root value
	if total >= root.Val {
		root.Val = total
	} else { // total is less then root value, we replace the left and right node's value

		// the reason behind why we set either of the child node is
		// our intuition is to make the child nodes greater than or equal to parent node
		// so making one of the node's value with root value satisfies the intuition
		// still, there is no harm if we set both of them
		if root.Left != nil {
			root.Left.Val = root.Val
		} else if root.Right != nil {
			root.Right.Val = root.Val
		}
	}

	// traverse recursively to the left and right
	changeTree(root.Left)
	changeTree(root.Right)

	// calculate the final count to make sure the root follows the children sum property
	count := 0

	// make sure root has left node
	if root.Left != nil {
		count += root.Left.Val
	}

	// make sure root has right node
	if root.Right != nil {
		count += root.Right.Val
	}

	// except leaf nodes, update the root node's value with calculated count
	// we make sure, we are ignoring the leaf node here.
	if root.Left != nil || root.Right != nil {
		root.Val = count
	}
}
