package main

import "strings"

// findOrder returns a valid order of characters in the alien alphabet
// based on the lexicographically sorted list of words.
func findOrder(words []string) string {
	// We'll use indices 0–25 to represent 'a' to 'z'
	// adjList[i] will contain all characters that come AFTER 'a'+i
	// i.e., an edge from i → j means 'a'+i must come before 'a'+j
	adjList := make([][]int, 26)

	// exists[i] = true if character 'a'+i appears anywhere in the input
	exists := make([]bool, 26)

	// Mark every character that appears in any word
	for _, word := range words {
		for _, char := range word {
			exists[char-'a'] = true
		}
	}

	// Compare every pair of consecutive words to extract ordering clues
	for i := 0; i < len(words)-1; i++ {
		s1 := words[i]
		s2 := words[i+1]
		minLen := min(len(s1), len(s2))

		// Compare character by character until the first mismatch
		for j := range minLen {
			if s1[j] != s2[j] {
				// The first differing character determines the order:
				// s1[j] < s2[j] in the alien alphabet
				from := int(s1[j] - 'a') // index of earlier character
				to := int(s2[j] - 'a')   // index of later character

				// Add directed edge: from → to
				adjList[from] = append(adjList[from], to)

				// Only the first difference matters for lexicographical order
				// Break immediately after adding the edge
				break
			}
		}
	}

	// Now we have a directed graph of character dependencies
	// Use BFS (Kahn's algorithm) for topological sort to get a valid order
	return bfs(adjList, exists)
}

// bfs performs topological sorting using Kahn's algorithm
// It returns the characters in a valid order, only including those that exist
func bfs(adjList [][]int, exists []bool) string {
	// indegree[i] = number of edges pointing to node i
	// i.e., how many characters must come before 'a'+i
	indegree := make([]int, 26)

	// Calculate indegree by scanning all edges in the graph
	for i := range 26 {
		for _, neighbor := range adjList[i] {
			indegree[neighbor]++
		}
	}

	// Queue starts with all existing characters that have no incoming edges
	// (i.e., no other char needs to come before them)
	queue := []int{}
	for i := range 26 {
		if exists[i] && indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	var str strings.Builder

	// Process nodes in topological order
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		// Append the character corresponding to this node
		str.WriteRune(rune(node) + 'a')

		// Remove this node and update its neighbors
		for _, neighbor := range adjList[node] {
			indegree[neighbor]--

			// If a neighbor now has no dependencies, it's ready to be placed
			if indegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// The result is the reconstructed alien alphabet order
	// Note: If not all existing nodes were processed, there's a cycle.
	// But this code assumes valid input — no cycle or invalid order.
	return str.String()
}
