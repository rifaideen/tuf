package main

func minimumMultiplications(arr []int, start, end int) int {
	// If we're already at the target, no multiplication is needed.
	if start == end {
		return 0
	}

	const mod = 100000 // The problem specifies modulo 100000

	// distances[i] stores the minimum number of multiplications to reach 'i'
	// Initialize all distances to a large number (infinity)
	distances := make([]int, mod)

	for i := range distances {
		distances[i] = 1e9 // "infinity" to mark unvisited or unreachable
	}

	// Start with the initial value; 0 multiplications needed to be at 'start'
	distances[start] = 0

	// Queue holds pairs: [current_value, steps_taken]
	// We use a slice as a queue (FIFO) for BFS
	queue := [][]int{{start, 0}}

	// BFS to explore all reachable states level by level
	for len(queue) > 0 {
		node, time := queue[0][0], queue[0][1]
		queue = queue[1:] // Pop from front

		// Try multiplying current value with each number in the array
		for _, num := range arr {
			neighbor := (node * num) % mod // Next state after multiplication + mod
			newTime := time + 1            // One more step

			// If we found a shorter way to reach 'neighbor', update it
			if newTime < distances[neighbor] {
				distances[neighbor] = newTime

				// If we reached the target, return the number of steps immediately
				// Since BFS guarantees the first time we reach 'end' is the shortest
				if neighbor == end {
					return newTime
				}

				// Add the new state to the queue for further exploration
				queue = append(queue, []int{neighbor, newTime})
			}
		}
	}

	// If BFS completes without reaching 'end', it's unreachable
	return -1
}
