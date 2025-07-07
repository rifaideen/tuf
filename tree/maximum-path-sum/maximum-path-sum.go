package main

import "math"

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
func MaxPathSum(root *TreeNode) int {
	maxSum := math.MinInt
	height(root, &maxSum)

	return maxSum
}

func height(root *TreeNode, maxSum *int) int {
	if root == nil {
		return 0
	}

	// to avoid the values in negative, we put the max check here
	left := max(0, height(root.Left, maxSum))
	right := max(0, height(root.Right, maxSum))

	// find the maximum sum by adding left, right and then root value
	// left and right already has their actual value included
	*maxSum = max(*maxSum, left+right+root.Val)

	// here we take the maximum of left and right then add root value, becuase we are interested
	// in sum value, not the height
	return max(left, right) + root.Val
}
