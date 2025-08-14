package main

func topoSort(vertices int, adjList [][]int) []int {
	ans := []int{}
	stack := []int{}
	visited := make([]int, vertices)

	for i := range vertices {
		if visited[i] == 0 {
			dfs(i, visited, &stack, adjList)
		}
	}

	for i := len(stack) - 1; i >= 0; i-- {
		ans = append(ans, stack[i])
	}

	return ans
}

func dfs(node int, visited []int, stack *[]int, adjList [][]int) {
	visited[node] = 1

	for _, neighbor := range adjList[node] {
		if visited[neighbor] == 0 {
			dfs(neighbor, visited, stack, adjList)
		}
	}

	(*stack) = append((*stack), node)
}
