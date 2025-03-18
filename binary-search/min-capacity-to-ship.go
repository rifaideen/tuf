package main

func ShipWithinDays(weights []int, days int) int {
	left := -1
	right := 0
	ans := 501

	for _, weight := range weights {
		left = max(left, weight)
		right += weight
	}

	for left <= right {
		mid := (left + right) / 2
		actual := countDays(weights, mid)

		if actual <= days {
			ans = min(ans, mid)
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return ans
}

func countDays(weights []int, capacity int) int {
	days := 1
	sum := 0

	for _, weight := range weights {
		sum += weight

		if sum == capacity {
			days++
			sum = 0
		} else if sum >= capacity {
			days++
			sum = weight
		}
	}

	return days
}
