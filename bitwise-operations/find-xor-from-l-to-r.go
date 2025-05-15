package main

func findRange(left, right int) int {
	// since xor(n) calculates the result from 1 to n
	// we can solve using the forumula xor(left - 1)^ xor(right)
	return xor(left-1) ^ xor(right)
}

// find xor from 1 to n
func xor(n int) int {
	switch n % 4 {
	case 1:
		return 1
	case 2:
		return n + 1
	case 3:
		return 0
		// either we can keep it here or outside the switch statement
		// case 4:
		// 	return n
	}

	return n
}
