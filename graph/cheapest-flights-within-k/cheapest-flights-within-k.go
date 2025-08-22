package main

// Edge represents a flight from one city to another with a specific cost.
type Edge struct {
	to   int // destination city
	cost int // price of flight
}

// State represents a traveler's current state during BFS traversal.
type State struct {
	city      int // current city
	totalCost int // total cost to reach this city
	numStops  int // number of stops (i.e., flights taken so far)
}

// findCheapestPrice returns the cheapest price from src to dst with at most k stops.
// Uses a modified BFS (like a queue-based relaxation) to explore paths within stop limit.
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	// Build adjacency list: graph[i] = list of flights from city i
	graph := make([][]*Edge, n)
	// fares[i] stores the cheapest cost found so far to reach city i
	fares := make([]int, n)

	for i := range n {
		graph[i] = []*Edge{}
		fares[i] = 1e9 // initialize with a large value (infinity)
	}

	// Starting point: cost to reach src is 0
	fares[src] = 0

	// Populate the graph with flight routes
	for _, flight := range flights {
		from := flight[0]
		to := flight[1]
		price := flight[2]

		graph[from] = append(graph[from], &Edge{
			to:   to,
			cost: price,
		})
	}

	// BFS-like queue to explore paths, prioritizing by arrival order (not cost)
	// We use simple queue (not priority queue) because we're limiting by stops, not just cost
	var queue []State
	queue = append(queue, State{
		city:      src,
		totalCost: 0,
		numStops:  0,
	})

	// Process each state in FIFO order
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// If we've exceeded the stop limit, don't expand further from this path
		if current.numStops > k {
			continue
		}

		// Check all outgoing flights from current city
		for _, edge := range graph[current.city] {
			nextCity := edge.to
			flightCost := edge.cost
			newTotalCost := current.totalCost + flightCost

			// Only proceed if we found a cheaper way to reach nextCity
			if newTotalCost < fares[nextCity] {
				fares[nextCity] = newTotalCost

				// Add the next state to the queue for further exploration
				queue = append(queue, State{
					city:      nextCity,
					totalCost: newTotalCost,
					numStops:  current.numStops + 1, // one additional flight (i.e., one stop)
				})
			}
		}
	}

	// If destination is unreachable within k stops, return -1
	if fares[dst] == 1e9 {
		return -1
	}

	return fares[dst]
}
