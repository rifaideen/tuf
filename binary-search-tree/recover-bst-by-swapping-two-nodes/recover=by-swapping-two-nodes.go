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
func recoverTree(root *TreeNode) {
	s := &Solution{}
	s.inorder(root)

	if s.first != nil && s.last != nil {
		s.first.Val, s.last.Val = s.last.Val, s.first.Val
	} else {
		s.first.Val, s.second.Val = s.second.Val, s.first.Val
	}
}

type Solution struct {
	first  *TreeNode
	second *TreeNode
	last   *TreeNode
	prev   *TreeNode
}

func (s *Solution) inorder(node *TreeNode) {
	if node == nil {
		return
	}

	s.inorder(node.Left)

	if s.prev != nil && node.Val < s.prev.Val {
		if s.first == nil {
			s.first = s.prev
			s.second = node
		} else {
			s.last = node
		}
	}

	s.prev = node
	s.inorder(node.Right)
}
