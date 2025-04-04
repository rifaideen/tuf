package main

func searchKey(node *Node, key int) bool {
	current := node

	for {
		if current.data == key {
			return true
		}

		if current.next == nil {
			return false
		}

		current = current.next
	}
}
