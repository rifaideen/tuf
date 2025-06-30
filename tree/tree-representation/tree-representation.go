package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func main() {
	tree := &Node{
		Value: 1,
	}

	tree.Left = &Node{
		Value: 2,
	}

	tree.Right = &Node{
		Value: 3,
	}

	tree.Left.Left = &Node{
		Value: 4,
	}

	tree.Left.Right = &Node{
		Value: 5,
	}

	fmt.Println(tree)
}
