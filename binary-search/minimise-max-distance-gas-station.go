package main

import "github.com/emirpasic/gods/queues/priorityqueue"

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
	howMany := make([]int, len(nums)-1)

	type pair struct {
		diff  float64
		index int
	}

	pq := priorityqueue.NewWith(nil)

	for i := range len(nums) - 1 {
		pq.Enqueue(pair{
			diff:  float64(nums[i+1]) - float64(nums[i]),
			index: i,
		})
	}

	for gasStation := 1; gasStation <= k; gasStation++ {
		top, _ := pq.Peek()
		pq.Dequeue()

		sectionIndex := top.(pair).index
		howMany[sectionIndex]++
		initialDiff := float64(nums[sectionIndex+1] - nums[sectionIndex])
		newSectionLength := initialDiff / float64(howMany[sectionIndex]+1)

		pq.Enqueue(pair{
			diff:  newSectionLength,
			index: sectionIndex,
		})
	}

	ans, _ := pq.Dequeue()

	return ans.(pair).diff
}
