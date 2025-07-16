package main

import (
	"strconv"
	"strings"
)

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

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// serialize encodes a binary tree into a comma-separated string using level-order traversal.
// Nil nodes are represented by "#" to denote null pointers in the tree structure.
func (t *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	var result []*int // Using *int to differentiate between missing and present values

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if node == nil {
				result = append(result, nil)
			} else {
				result = append(result, &node.Val)
				queue = append(queue, node.Left, node.Right)
			}
		}
	}

	// Build the final string from the result slice
	var sb strings.Builder

	for i, valPtr := range result {
		if valPtr != nil {
			sb.WriteString(strconv.Itoa(*valPtr))
		} else {
			sb.WriteString("#")
		}

		if i < len(result)-1 {
			sb.WriteString(",")
		}
	}

	return sb.String()
}

// Deserializes your encoded data to tree.
// deserialize decodes a serialized binary tree from a comma-separated string.
// Uses BFS to reconstruct the tree level-by-level.
func (t *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	nodes := strings.Split(data, ",")
	index := 0

	// Create root node
	rootVal, _ := strconv.Atoi(nodes[index])
	root := &TreeNode{Val: rootVal}
	index++

	queue := []*TreeNode{root}

	for index < len(nodes) && len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]

		// Process left child
		leftVal := nodes[index]
		index++
		if leftVal != "#" {
			val, _ := strconv.Atoi(leftVal)
			currentNode.Left = &TreeNode{Val: val}
			queue = append(queue, currentNode.Left)
		}

		// Process right child
		if index >= len(nodes) {
			break
		}

		rightVal := nodes[index]
		index++
		if rightVal != "#" {
			val, _ := strconv.Atoi(rightVal)
			currentNode.Right = &TreeNode{Val: val}
			queue = append(queue, currentNode.Right)
		}
	}

	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
