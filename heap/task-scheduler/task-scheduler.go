package main

// Task Scheduler - 621 Leetcode
// We count the frequency of the given tasks by creating array of size 26 to represent A-Z
//
// Create max-heap to store the frequency count, once completed, we start to pop until the heap becomes empty.
//
// We start the loop while heap not equals to empty, we create a variable to store the number of cycles (n + 1) 1 time per item and n time required to complete the cycle.
//
// we create couple of variables to store the number of tasks completed, temporary array to store the popped frequency.
//
// push the items back into the heap
func leastInterval(tasks []byte, n int) int {
	frequency := [26]int{}

	// count the task frequency
	for _, task := range tasks {
		frequency[task-'A']++
	}

	// create max-heap and store the collected frequency, ignoring zeroes.
	h := &maxheap{
		size: 0,
		nums: []int{},
	}

	for _, value := range frequency {
		if value > 0 {
			h.push(value)
		}
	}

	// store the running sum
	sum := 0

	// while heap is not empty, we start to pop the values from the heap upto n + 1 times in the inner loop
	for !h.empty() {
		// we can pop upto n + 1 times or heap becomes empty before completing n + 1 operations.
		cycle := n + 1
		// create temporary array to store the reduced frequency after popping it from heap
		temp := []int{}
		// the number of cycles completed
		completed := 0

		// start the inner loop while cycle greater than 0 and heap is not empty
		for cycle > 0 && !h.empty() {
			frequency := h.pop()

			// if we have more frequency to process, we need to put it to the temporary array
			if frequency > 1 {
				temp = append(temp, frequency-1)
			}

			// we have completed one task and one cycle completed
			completed++
			cycle--
		}

		// put the temporary array back to the heap
		for _, val := range temp {
			h.push(val)
		}

		// while performing the above operation, if the heap becomes empty, we have completed all the operations and we can take the completion count to the sum
		// else, we need to consider n+1 to the sum
		if h.empty() {
			sum += completed
		} else {
			sum += n + 1
		}
	}

	return sum
}

type maxheap struct {
	nums []int
	size int
}

func (h *maxheap) push(value int) {
	h.nums = append(h.nums, value)
	h.up(h.size)
	h.size++
}

func (h *maxheap) pop() int {
	if h.size == 0 {
		panic("heap empty")
	}

	top := h.nums[0]

	h.size--
	h.nums[0] = h.nums[h.size]
	h.nums = h.nums[:h.size]

	h.down(0)

	return top
}

func (h *maxheap) up(index int) {
	if index > 0 {
		parent := (index - 1) / 2

		if h.nums[index] > h.nums[parent] {
			h.swap(index, parent)
			h.up(parent)
		}
	}
}

func (h *maxheap) down(i int) {
	left := (2 * i) + 1
	right := left + 1
	largest := i

	if left < h.size && h.nums[left] > h.nums[largest] {
		largest = left
	}

	if right < h.size && h.nums[right] > h.nums[largest] {
		largest = right
	}

	if largest != i {
		h.swap(i, largest)
		h.down(largest)
	}
}

func (h *maxheap) swap(i, j int) {
	h.nums[i], h.nums[j] = h.nums[j], h.nums[i]
}

func (h *maxheap) empty() bool {
	return h.size == 0
}
