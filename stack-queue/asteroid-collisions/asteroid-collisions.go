package main

func asteroidCollision(asteroids []int) []int {
	st := []int{}

	for _, i := range asteroids {
		if i > 0 {
			st = append(st, i)
		} else {
			for len(st) > 0 && st[len(st)-1] > 0 && st[len(st)-1] < (-i) {
				st = st[:len(st)-1]
			}

			if len(st) > 0 && st[len(st)-1] == (-i) {
				st = st[:len(st)-1]
			} else if len(st) == 0 || st[len(st)-1] < 0 {
				st = append(st, i)
			}
		}
	}

	return st
}

// type Stack struct {
// 	top    int
// 	values []int
// }

// func NewStack() *Stack {
// 	return &Stack{
// 		top:    -1,
// 		values: []int{},
// 	}
// }

// func (t *Stack) Push(x int) {
// 	t.top++
// 	t.values = append(t.values, x)
// }

// func (t *Stack) Pop() int {
// 	data := t.values[t.top]
// 	t.values = t.values[:t.top]
// 	t.top--

// 	return data
// }

// func (t *Stack) Top() int {
// 	if t.top < 0 {
// 		panic("stack is empty")
// 	}

// 	return t.values[t.top]
// }

// func (t *Stack) Empty() bool {
// 	return t.top < 0
// }
