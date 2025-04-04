package main

func insertBeginning(node *Node, num int) *Node {
	return &Node{
		data: num,
		next: node,
	}
}
