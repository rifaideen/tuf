package main

import (
	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/utils"
)

// Brute force solution
func MinimiseMaxDistance(nums []int, k int) float64 {
	howMany := make([]int, len(nums)-1)

	for gasStation := 1; gasStation <= k; gasStation++ {
		maxSection := -1
		maxIndex := -1

		for i := range len(nums) - 1 {
			diff := nums[i+1] - nums[i]

			sectionLength := diff / (howMany[i] + 1)

			if sectionLength > maxSection {
				maxSection = sectionLength
				maxIndex = i
			}
		}

		howMany[maxIndex] = maxSection
	}

	ans := -1.0

	for i := range len(nums) - 1 {
		diff := float64(nums[i+1] - nums[i])
		sectionLength := diff / float64(howMany[i]+1)
		ans = max(ans, sectionLength)
	}

	return ans
}

// Better solution using priority queue
// It's not tested yet.
func MinimiseMaxDistancePQ(nums []int, k int) float64 {
	// create new array of size n-1 filled with 0 initially
	howMany := make([]int, len(nums)-1)

	type pair struct {
		priority float64 // priority between gas stations
		index    int     // item index
	}

	// priorotize the pair based on priority attribute in descending order
	priorotize := func(a, b any) int {
		priorityA := a.(pair).priority
		priorityB := b.(pair).priority

		return -utils.Float64Comparator(priorityA, priorityB) // "-" descending order
	}

	// create priority queue
	pq := priorityqueue.NewWith(priorotize)

	// feed the priority queue
	for i := range len(nums) - 1 {
		pq.Enqueue(pair{
			// calculate distance between next and current gas station and update the priority
			priority: float64(nums[i+1]) - float64(nums[i]),
			index:    i,
		})
	}

	// we must place the k number of gas stations
	for gasStation := 1; gasStation <= k; gasStation++ {
		// pull the highest pair and update the priority
		top, ok := pq.Dequeue()

		if !ok {
			panic("queue is empty")
		}

		sectionIndex := top.(pair).index
		howMany[sectionIndex]++
		initialDiff := float64(nums[sectionIndex+1] - nums[sectionIndex])
		newSectionLength := initialDiff / float64(howMany[sectionIndex]+1)

		pq.Enqueue(pair{
			priority: newSectionLength,
			index:    sectionIndex,
		})
	}

	ans, _ := pq.Dequeue()

	return ans.(pair).priority
}
