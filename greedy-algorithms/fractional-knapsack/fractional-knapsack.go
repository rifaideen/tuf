package main

import "fmt"

type Item struct {
	value      int
	weight     int
	qty_per_kg float64
}

// you are given weights and the values (weight[0] represents number of quantity at values[0]) with capacity
// our task is to maximize the number of items we can fill that fits the capacity
//
// we first calculate the quantity per kg by dividing quantity / weight push the item into the max heap
// which maintains the priority based on quantity per kg.
//
// then we start to pop from the heap while maintaining running sum and the weight we take so far.
//
// by the time we are unable to pick the entire quantity, we pick the fractional quantity to fill the bag to full capacity
func fractionalKnapsack(values, weights []int, capacity int) float64 {
	h := MaxHeap{
		size:   0,
		values: make([]Item, 0, len(weights)),
	}

	for i := range values {
		h.Push(Item{
			value:      values[i],
			weight:     weights[i],
			qty_per_kg: float64(values[i]) / float64(weights[i]),
		})
	}

	sum := 0.0
	current_weight := 0

	for h.size > 0 {
		value := h.Pop()
		fmt.Println(value)

		if current_weight+value.weight <= capacity {
			sum += float64(value.value)
			current_weight += value.weight
		} else {
			// we fill the last portion in fraction and stop the loop
			sum += float64((capacity - current_weight)) * value.qty_per_kg
			break
		}
	}

	return sum
}

type MaxHeap struct {
	values []Item
	size   int
}

func (h *MaxHeap) Push(value Item) {
	h.values = append(h.values, value)
	h.Up(h.size)

	h.size++
}

func (h *MaxHeap) Pop() Item {
	if h.size == 0 {
		panic("heap is empty")
	}

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

		if h.values[index].qty_per_kg > h.values[parent].qty_per_kg {
			h.Swap(index, parent)
			h.Up(parent)
		}
	}
}

func (h *MaxHeap) Down(index int) {
	left := 2*index + 1
	right := left + 1
	largest := index

	if left < h.size && h.values[left].qty_per_kg > h.values[largest].qty_per_kg {
		largest = left
	}

	if right < h.size && h.values[right].qty_per_kg > h.values[largest].qty_per_kg {
		largest = right
	}

	if index != largest {
		h.Swap(index, largest)
		h.Down(largest)
	}
}

func (h *MaxHeap) Swap(i, j int) {
	h.values[i], h.values[j] = h.values[j], h.values[i]
}
