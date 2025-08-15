package main

// canFinish determines if it's possible to finish all courses given the prerequisites.
// This is a classic "Course Schedule" problem that reduces to detecting cycles in a directed graph.
// We use Kahn's Algorithm for Topological Sorting to solve it.
func canFinish(numCourses int, prerequisites [][]int) bool {
	// Step 1: Build the adjacency list representation of the graph.
	// adjList[i] contains all courses that depend on course i (i is a prerequisite for them).
	// Note: Edge direction is from course -> its prerequisite (reverse of dependency order).
	adjList := make([][]int, numCourses)
	for _, prereq := range prerequisites {
		course, dependency := prereq[0], prereq[1]
		adjList[course] = append(adjList[course], dependency)
	}

	// Step 2: Calculate the in-degree (number of incoming edges) for each node.
	// indegree[i] = number of prerequisites that must be taken before course i.
	indegree := make([]int, numCourses)
	for node := range adjList {
		for _, neighbor := range adjList[node] {
			indegree[neighbor]++
		}
	}

	// Step 3: Initialize a queue with all nodes having in-degree 0.
	// These courses have no prerequisites and can be taken immediately.
	queue := []int{}
	for node, degree := range indegree {
		if degree == 0 {
			queue = append(queue, node)
		}
	}

	// Step 4: Count how many courses we can successfully "take" using BFS.
	counter := 0
	for len(queue) > 0 {
		// Take a course with no remaining prerequisites
		node := queue[0]
		queue = queue[1:]
		counter++

		// For each course that depends on this one, remove the dependency
		// (simulate completing 'node', thus reducing prerequisite count for its dependents)
		for _, neighbor := range adjList[node] {
			indegree[neighbor]--

			// If the course now has no prerequisites, it can be taken next
			if indegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// Step 5: If we were able to take all courses (counter == numCourses),
	// then there's no cycle and all courses can be finished.
	// Otherwise, a cycle exists and we return false.
	return counter == numCourses
}
