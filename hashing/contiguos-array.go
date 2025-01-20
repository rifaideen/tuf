package main

func ContiguosArray(nums []int) int {
	m := map[int]int{ // count => index pair
		0: -1,
	}

	count := 0  // counter
	maxlen := 0 // maximum length of contiguos subarray we have seen so far

	for i, num := range nums {
		// decrement when 0 and increment when 1
		if num == 0 {
			count--
		} else {
			count++
		}

		// did we see the same count already? then calculate the maxlen
		if j, ok := m[count]; ok {
			maxlen = max(maxlen, i-j)
		} else { // we haven't seen before, add it to map
			m[count] = i
		}
	}

	return maxlen
}
