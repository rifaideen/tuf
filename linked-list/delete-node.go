package main

func deleteNode(node *Node) {
	*node = *node.next
}
