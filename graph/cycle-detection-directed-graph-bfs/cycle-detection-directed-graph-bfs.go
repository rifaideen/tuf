package main

// detectCycle checks for cycles in a directed graph using Kahn's algorithm
// INTUITION:
//
//	In a dependency graph (e.g., tasks with prerequisites), a cycle means:
//	"Task A requires Task B, which requires Task C, which requires Task A..."
//	→ No task in the cycle can ever start (all have unresolvable prerequisites)
//
//	We simulate executing tasks in safe order:
//	1. Count how many prerequisites each task has (inDegree)
//	2. Start with tasks having NO prerequisites (inDegree=0)
//	3. When a task completes, it unblocks dependent tasks
//	4. If ANY task remains unexecuted at the end → cycle exists
func detectCycle(vertices int, adjList [][]int) bool {
	processedCount := 0               // Tracks how many tasks we successfully executed
	inDegree := make([]int, vertices) // "Prerequisite counters" for each task

	// Count prerequisites for every task
	// Why? To identify which tasks can start immediately (0 prerequisites)
	for node := range adjList {
		for _, dependent := range adjList[node] {
			inDegree[dependent]++ // dependent needs node as prerequisite
		}
	}

	// Queue of tasks ready to execute (no remaining prerequisites)
	queue := []int{}

	// Find all initially executable tasks
	for node := range inDegree {
		if inDegree[node] == 0 { // No prerequisites? Ready to run!
			queue = append(queue, node)
		}
	}

	// Process tasks in dependency-safe order
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:] // Remove from queue (task is executing)
		processedCount++  // Successfully executed one more task

		// When current task finishes, it unblocks its dependents
		for _, dependent := range adjList[current] {
			inDegree[dependent]-- // Remove one prerequisite (current task)

			// If dependent now has NO prerequisites left...
			if inDegree[dependent] == 0 {
				queue = append(queue, dependent) // ...it becomes executable
			}
		}
	}

	// CYCLE DETECTION LOGIC:
	//   In a valid dependency graph (DAG), we should execute ALL tasks
	//   If ANY task remains unexecuted (processedCount < vertices):
	//     → Some tasks have unresolvable prerequisites (cycle exists)
	//     → Example: A→B→C→A leaves all 3 tasks with 1+ prerequisites forever
	return processedCount != vertices
}
