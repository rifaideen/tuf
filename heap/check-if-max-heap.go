package main

func isMaxheap(values []int) bool {
	n := len(values)

	if n == 1 {
		return true
	}

	for i := 0; i < n/2; i++ {
		left := 2*i + 1
		right := left + 1

		if left < n && values[left] > values[i] {
			return false
		}

		if right < n && values[right] > values[i] {
			return false
		}
	}

	return true
}
