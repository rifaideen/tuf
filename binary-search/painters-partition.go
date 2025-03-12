package main

func FindLargestMinDistance(boards []int, k int) int {
	left := 0
	right := 0

	for _, board := range boards {
		left = max(left, board)
		right += board
	}

	for left <= right {
		mid := (left + right) / 2

		if calculate(boards, mid) <= k {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}

func calculate(boards []int, val int) int {
	count := 0
	painters := 1

	for _, board := range boards {
		if count+board <= val {
			count += board
		} else {
			painters++
			count = board
		}
	}

	return painters
}
