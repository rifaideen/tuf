package main

func jobSequencing(deadlines, profits []int) []int {
	h := &MaxHeap{
		size:   0,
		values: [][]int{},
	}

	maxDeadline := -1

	for i := range profits {
		maxDeadline = max(maxDeadline, deadlines[i])
		h.Push([]int{i + 1, deadlines[i], profits[i]})
	}

	sum := 0
	schedule := make([]int, maxDeadline+1)
	earnings := 0

	for h.size > 0 {
		value := h.Pop()

		for j := value[1]; j >= 0; j-- { // loop from the deadline till 0
			if schedule[j] == 0 {
				schedule[j] = value[0] // store the id
				earnings += value[2]   // sum the earnings
				sum++
				break
			}
		}
	}

	return []int{sum, earnings}
}

type MaxHeap struct {
	values [][]int // [id, deadline, profit]
	size   int
}

func (h *MaxHeap) Push(value []int) {
	h.values = append(h.values, value)
	h.Up(h.size)

	h.size++
}

func (h *MaxHeap) Pop() []int {
	top := h.values[0]

	h.size--
	h.values[0] = h.values[h.size]
	h.values = h.values[:h.size]

	h.Down(0)

	return top
}

func (h *MaxHeap) Up(index int) {
	if index > 0 {
		parent := (index - 1) / 2

		if h.values[index][2] > h.values[parent][2] {
			h.Swap(index, parent)
			h.Up(parent)
		}
	}
}

func (h *MaxHeap) Down(index int) {
	left := (2 * index) + 1
	right := left + 1
	smallest := index

	if left < h.size && h.values[left][2] > h.values[smallest][2] {
		smallest = left
	}

	if right < h.size && h.values[right][2] > h.values[smallest][2] {
		smallest = right
	}

	if index != smallest {
		h.Swap(index, smallest)
		h.Down(smallest)
	}
}

func (h *MaxHeap) Swap(i, j int) {
	h.values[i], h.values[j] = h.values[j], h.values[i]
}
