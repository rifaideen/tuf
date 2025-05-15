package main

func minBitFlips(start int, goal int) int {
	// performing xor on the start and goal will give us the idea of how many bits can be flipped
	ans := start ^ goal

	// now, we need to count how many bits are set in the answer
	count := 0

	// instead of looping from 0 to 31st bit, we use the division concept the optimize the calculation
	for ans > 0 {
		count += ans % 2
		ans /= 2
	}

	return count
}
