package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func traverse(nodes *TreeNode) [][]int {
	pre := []int{}
	post := []int{}
	in := []int{}

	stack := [][]any{
		{
			nodes,
			1,
		},
	}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		node := top[0].(*TreeNode)
		value := top[1].(int)

		switch value {
		case 1:
			// pre-order
			pre = append(pre, node.Val)

			value++
			stack = append(stack, []any{node, value})

			if node.Left != nil {
				stack = append(stack, []any{node.Left, 1})
			}
		case 2:
			// in-order
			in = append(in, node.Val)

			value++
			stack = append(stack, []any{node, value})

			if node.Right != nil {
				stack = append(stack, []any{node.Right, 1})
			}
		case 3:
			// post-order
			post = append(post, node.Val)
		}
	}

	return [][]int{
		pre,
		in,
		post,
	}
}
