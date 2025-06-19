package main

func isNStraightHand(hand []int, groupSize int) bool {
	n := len(hand)

	// check if we can perfectly divide the number of cards by group size
	if n%groupSize > 0 {
		return false
	}

	// create instance of min heap and push the cards
	h := &minheap{
		nums: make([]int, 0, len(hand)),
		size: 0,
	}

	for _, card := range hand {
		h.push(card)
	}

	// while heap is not empty, we start pop the value
	for !h.empty() {
		// the window size, number of items we can arrange per group
		cycle := groupSize
		// keep track of valid last card retrieved
		last := -1
		// temporarily store the un-used cards and put it back to the heap after a cycle is completed
		temp := []int{}
		// the number of cards arranged
		cards := 0

		// while the window is valid and heap is not empty, start popping out
		for cycle > 0 && !h.empty() {
			card := h.pop()

			// place the current card only when card == last + 1 or if this is the first iteration
			if last == -1 || last+1 == card {
				// card is consecutive or this is first iteration
				last = card // card is valid, replace with last card
				cycle--     // 1 operation completed
				cards++     // 1 valid card found
			} else {
				// value not in consecutive order, un-used card, goes into temp array
				temp = append(temp, card)
			}
		}

		// loop completed, number of cards processed should be equal to group size
		if cards != groupSize {
			return false // cards and group size doesn't match. stop early.
		}

		// put the un-used cards back to the heap
		for _, card := range temp {
			h.push(card)
		}
	}

	return true
}

type minheap struct {
	nums []int
	size int
}

func (h *minheap) push(value int) {
	h.nums = append(h.nums, value)
	h.up(h.size)
	h.size++
}

func (h *minheap) pop() int {
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

func (h *minheap) up(index int) {
	if index > 0 {
		parent := (index - 1) / 2

		if h.nums[index] < h.nums[parent] {
			h.swap(index, parent)
			h.up(parent)
		}
	}
}

func (h *minheap) down(i int) {
	left := (2 * i) + 1
	right := left + 1
	smallest := i

	if left < h.size && h.nums[left] < h.nums[smallest] {
		smallest = left
	}

	if right < h.size && h.nums[right] < h.nums[smallest] {
		smallest = right
	}

	if smallest != i {
		h.swap(i, smallest)
		h.down(smallest)
	}
}

func (h *minheap) swap(i, j int) {
	h.nums[i], h.nums[j] = h.nums[j], h.nums[i]
}

func (h *minheap) empty() bool {
	return h.size == 0
}
