package main

func totalFruit(fruits []int) int {
	left := 0
	maxLen := 0
	basket := map[int]int{}

	for right, fruit := range fruits {
		basket[fruit]++

		if len(basket) > 2 {
			basket[fruits[left]]--

			if basket[fruits[left]] == 0 {
				delete(basket, fruits[left])
			}

			left++
		} else {
			maxLen = max(maxLen, right-left+1)
		}
	}

	return maxLen
}
