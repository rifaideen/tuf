package main

func maxMeetings(start, end []int) int {
	// push the meeting start time, end time and index (1 based) and sort them by end time
	h := &MinHeap{
		size:   0,
		values: make([][]int, 0, len(start)),
	}

	for key, s := range start {
		h.Push([]int{s, end[key], key + 1})
	}

	// create variables to store the answer and when does the current meeting ends
	sum := 0
	freeAt := -1 // start with no meeting

	for h.size > 0 {
		meeting := h.Pop()

		// first meeting or meeting start time is greater than last meeting end time
		if meeting[0] > freeAt { // meeting starts after the current meeting ends
			freeAt = meeting[1] // meeting ends at meeting[1]
			sum++
		}
	}

	return sum
}

type MinHeap struct {
	values [][]int
	size   int
}

func (h *MinHeap) Push(value []int) {
	h.values = append(h.values, value)
	h.Up(h.size)

	h.size++
}

func (h *MinHeap) Pop() []int {
	top := h.values[0]

	h.size--
	h.values[0] = h.values[h.size]
	h.values = h.values[:h.size]

	h.Down(0)

	return top
}

func (h *MinHeap) Up(index int) {
	if index > 0 {
		parent := (index - 1) / 2

		if h.values[index][1] < h.values[parent][1] {
			h.Swap(index, parent)
			h.Up(parent)
		}
	}
}

func (h *MinHeap) Down(index int) {
	left := (2 * index) + 1
	right := left + 1
	smallest := index

	if left < h.size && h.values[left][1] < h.values[smallest][1] {
		smallest = left
	}

	if right < h.size && h.values[right][1] < h.values[smallest][1] {
		smallest = right
	}

	if index != smallest {
		h.Swap(index, smallest)
		h.Down(smallest)
	}
}

func (h *MinHeap) Swap(i, j int) {
	h.values[i], h.values[j] = h.values[j], h.values[i]
}
