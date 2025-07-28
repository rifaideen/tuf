package main

// Node is a basic structure for a binary tree node.
type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

// Solution is a struct to hold the methods for finding predecessor and successor.
type Solution struct {
}

// findPreSuc finds the predecessor and successor of a given key in a BST.
func (s *Solution) FindPreSuc(root *Node, key int) (*Node, *Node) {
	return s.predecessor(root, key), s.successor(root, key)
}

// successor finds the successor of a given key in a BST.
func (s *Solution) successor(root *Node, key int) *Node {
	var ans *Node = nil
	node := root

	for node != nil {
		if node.Data <= key {
			node = node.Right
		} else {
			ans = node
			node = node.Left
		}
	}
	return ans
}

// predecessor finds the predecessor of a given key in a BST.
func (s *Solution) predecessor(root *Node, key int) *Node {
	var ans *Node = nil
	node := root

	for node != nil {
		if node.Data >= key {
			node = node.Left
		} else {
			ans = node
			node = node.Right
		}
	}
	return ans
}
