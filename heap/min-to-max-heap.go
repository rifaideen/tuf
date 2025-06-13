package main

func minToMaxHeap(values []int) []int {
	n := len(values)

	if n == 1 {
		return values
	}

	var heapifyDown func(index int)
	heapifyDown = func(index int) {
		left := 2*index + 1
		right := left + 1
		largest := index

		if left < n && values[left] > values[largest] {
			largest = left
		}

		if right < n && values[right] > values[largest] {
			largest = right
		}

		if largest != index {
			values[index], values[largest] = values[largest], values[index]

			heapifyDown(largest)
		}
	}

	for i := n/2 - 1; i >= 0; i-- {
		heapifyDown(i)
	}

	return values
}
