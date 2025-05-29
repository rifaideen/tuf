package main

type StockSpanner struct {
	stack [][2]int // index => value pair
	index int
}

func Constructor() StockSpanner {
	return StockSpanner{
		index: -1,
	}
}

func (t *StockSpanner) Next(price int) int {
	t.index++

	// by finding pse element and using monotonic stack in decreasing order (highest values at front and next highest values follows till the top)

	// while stack top <= current price, we pop the stack
	for len(t.stack) > 0 && t.stack[len(t.stack)-1][1] <= price {
		t.stack = t.stack[:len(t.stack)-1]
	}

	// checking if stack is not empty and pick the index stored on top,  defaults to -1 when stack is empty
	left := -1

	if len(t.stack) > 0 {
		left = t.stack[len(t.stack)-1][0]
	}

	// push it to stack
	t.stack = append(t.stack, [2]int{t.index, price})

	return t.index - left
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
