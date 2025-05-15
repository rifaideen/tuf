package main

func findOddNumberOccurence(nums []int) []int {
	xor := 0

	// perform xor operation of all numbers
	for _, num := range nums {
		xor ^= num
	}

	// we know that at-lease there will be 1 bit difference between two single numbers
	// so, we need to find the right most set bit
	// we know how to turn the right most set bit from 1 to 0 and keep the rest intact
	// by performing xor of number with number - 1
	// performing xor with (xor of number and number - 1) will give us
	// find the right most set bit
	rightMostSetBit := (xor & (xor - 1)) ^ xor

	// we are trying to split those numbers into 2 separate buckets
	bucket1 := 0
	bucket2 := 0

	for _, num := range nums {
		// check if we are able to perform logical and of num and right most set bit
		if num&rightMostSetBit > 0 {
			// keep it in bucket 1
			bucket1 ^= num
		} else {
			// logic and was not successful, keep it in bucket 2
			bucket2 ^= num
		}
	}

	return []int{bucket1, bucket2}
}
