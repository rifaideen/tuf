package main

// findOrder returns a valid order in which courses can be taken, satisfying all prerequisites.
// If no valid order exists (i.e., there's a cycle), it returns an empty slice.
// Uses Kahn's Algorithm for topological sorting.
func findOrder(numCourses int, prerequisites [][]int) []int {
	// Step 1: Build the adjacency list where:
	// adjList[prereq] -> list of courses that depend on prereq (i.e., courses you can take *after* prereq)
	// Edge direction: from prerequisite → course (natural dependency flow)
	// This is the correct direction for building topological order.
	adjList := make([][]int, numCourses)
	for _, prereq := range prerequisites {
		course, dependency := prereq[0], prereq[1]
		adjList[dependency] = append(adjList[dependency], course) // dependency -> course
	}

	// Step 2: Calculate in-degree for each node.
	// indegree[i] = number of prerequisites required before taking course i
	indegree := make([]int, numCourses)
	for _, prereq := range prerequisites {
		course, _ := prereq[0], prereq[1]
		indegree[course]++ // each prerequisite increases the in-degree of the course
	}

	// Step 3: Initialize queue with all courses that have no prerequisites
	queue := []int{}
	for course := 0; course < numCourses; course++ {
		if indegree[course] == 0 {
			queue = append(queue, course)
		}
	}

	// Step 4: Use BFS (Kahn's Algorithm) to build topological order
	var result []int

	for len(queue) > 0 {
		// Take a course with no remaining prerequisites
		current := queue[0]
		queue = queue[1:]
		result = append(result, current) // add it to the order

		// For each course that depends on this one, reduce its in-degree
		for _, nextCourse := range adjList[current] {
			indegree[nextCourse]--

			// If all prerequisites for nextCourse are satisfied, it can be taken now
			if indegree[nextCourse] == 0 {
				queue = append(queue, nextCourse)
			}
		}
	}

	// Step 5: If we processed all courses, return the order
	// Otherwise, a cycle exists and no valid order is possible
	if len(result) == numCourses {
		return result
	}

	return []int{} // return empty slice if not all courses can be taken (cycle detected)
}
