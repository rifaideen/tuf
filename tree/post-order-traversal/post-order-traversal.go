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
func postorderTraversal(root *TreeNode) []int {
	return traverse(root)
}

func traverse(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ans := []int{}

	ans = append(ans, traverse(root.Left)...)
	ans = append(ans, traverse(root.Right)...)
	ans = append(ans, root.Val)

	return ans
}

// Binary tree postorder traversal using two stacks
func PostorderTraversal2(root *TreeNode) []int {
	ans := []int{}

	// handle edge case
	if root == nil {
		return ans
	}

	// create two stacks and push root to stack 1
	stack1 := []*TreeNode{
		root,
	}
	stack2 := []*TreeNode{}

	// loop until stack 1 becomes empty
	for len(stack1) > 0 {
		// pop the value from stack 1 and push it to stack 2
		current := stack1[len(stack1)-1]
		stack1 = stack1[:len(stack1)-1]

		stack2 = append(stack2, current)

		// check if current node has left node, push it to stack 1
		if current.Left != nil {
			stack1 = append(stack1, current.Left)
		}

		// check if current node has right node, push it to stack 1
		if current.Right != nil {
			stack1 = append(stack1, current.Right)
		}
	}

	// Pop the values from stack2 and push it to the answers
	for len(stack2) > 0 {
		ans = append(ans, stack2[len(stack2)-1].Val)
		stack2 = stack2[:len(stack2)-1]
	}

	return ans
}

// Binary tree postorder traversal using single stacks
func PostorderTraversal3(root *TreeNode) []int {
	ans := []int{}

	// handle edge case: If the tree is empty, there's nothing to traverse, so just return an empty slice.
	if root == nil {
		return ans
	}

	// create stack: We'll use a stack to keep track of nodes we need to visit, and 'current' will point to the node we're currently exploring.
	stack := []*TreeNode{}
	current := root

	// while current node or stack is not empty: We continue as long as there are still nodes to visit (either 'current' points to one, or there are nodes waiting in the stack).
	for current != nil || len(stack) > 0 {
		// move towards left, when current is not null: If there's a current node, we push it onto the stack and go left, because in postorder, we visit left children first.
		if current != nil {
			// push the current and move left
			stack = append(stack, current)
			current = current.Left
		} else {
			// current is null, pop the stack: If 'current' is null, it means we've gone as far left as possible. Now, we need to check the right child of the node at the top of our stack.

			// Get a temporary reference to the right child of the top node on the stack.
			temp := stack[len(stack)-1].Right

			// If there's no right child, or if the right child has already been processed (meaning it's the same as the node we just processed from the stack, if we were tracking that),
			// then we're done with this node's children, and we can process the node itself.
			if temp == nil || temp == stack[len(stack)-1] { // Note: The `temp == stack[len(stack)-1]` condition here is a common way to signal that the right child was just processed and is ready for its parent to be visited.
				// Pop the node from the stack.
				temp = stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				// Add its value to our result.
				ans = append(ans, temp.Val)

				// This inner loop handles cases where we've processed a right subtree and can now process its parent, and potentially its parent's parent, and so on, if they also had their right children processed.
				for len(stack) > 0 && temp == stack[len(stack)-1].Right {
					temp = stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					ans = append(ans, temp.Val)
				}
			} else {
				// If there's a right child and it hasn't been processed yet, we need to go explore that subtree.
				current = temp
			}
		}
	}

	return ans
}
