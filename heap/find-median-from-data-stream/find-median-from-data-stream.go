package main

type MedianFinder struct {
	h1 *maxheap
	h2 *minheap
}

// DRY RUN:
// h1 = 5,4,3,2,1
// h2 = 6,7,8,9,10
// add 11
// push it to max heap and pop from it
// h1 = 11,5,4,3,2,1
// pop from heap 1, we get 11 and push it to min heap
// h2 = 6,7,8,9,10,11
// observe that heap 2 always have higher elements when the sizes are odd numbers
// we should read the top from heap 2 in case of odd number of heaps

// Assuming we have this input 1,2,3,4,5,6,7,8,9,10
// we split the array into two parts, h1 and h2
// h1 = 1,2,3,4,5 so this has to be in the max heap (havent't arranged here)
// h2 = 6,7,8,9,10 and this has to be in the min heap (this is already arranged)
//
// The idea is when the heaps are in same size, we push it to max heap and pop the value from the max heap and push it to min heap
// when the sizes are not equal, we push the value in min heap, pop the value from min heap and push it to max heap
//
// while finding median, when heaps are same size, we have even number of items, read the top from heap 1 and 2 and divide by 2 and return it.
// otherwise read it from heap 2 top and return it
func Constructor() MedianFinder {
	return MedianFinder{
		h1: &maxheap{
			size: 0,
			nums: []int{},
		},
		h2: &minheap{
			size: 0,
			nums: []int{},
		},
	}
}

func (m *MedianFinder) AddNum(num int) {
	if m.h1.size == m.h2.size {
		m.h1.push(num)
		m.h2.push(m.h1.pop())
	} else {
		m.h2.push(num)
		m.h1.push(m.h2.pop())
	}
}

func (m *MedianFinder) FindMedian() float64 {
	if m.h1.size == m.h2.size {
		return float64(m.h1.nums[0]+m.h2.nums[0]) / 2.0
	}

	return float64(m.h2.nums[0])
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
	largest := i

	if left < h.size && h.nums[left] < h.nums[largest] {
		largest = left
	}

	if right < h.size && h.nums[right] < h.nums[largest] {
		largest = right
	}

	if largest != i {
		h.swap(i, largest)
		h.down(largest)
	}
}

func (h *minheap) swap(i, j int) {
	h.nums[i], h.nums[j] = h.nums[j], h.nums[i]
}

func (h *minheap) empty() bool {
	return h.size == 0
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
