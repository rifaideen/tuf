package main

func asteroidCollision(asteroids []int) []int {
	st := []int{}

	for _, i := range asteroids {
		if i > 0 { // put it to stack, the asteroids moves towards right side
			st = append(st, i)
		} else { // negative value, moved towards left side and about to collide with asteroids coming towards right side
			// while stack is not empty, and stack top is > 0 && stack top is less than the current asteroid, the stack top get destroyed
			// due to the collided asteroid is bigger
			for len(st) > 0 && st[len(st)-1] > 0 && st[len(st)-1] < (-i) {
				st = st[:len(st)-1]
			}

			// we have some asteroids left, and when both of them being same size, both gets destroyed (means, pop from stack)
			if len(st) > 0 && st[len(st)-1] == (-i) {
				st = st[:len(st)-1]
			} else if len(st) == 0 || st[len(st)-1] < 0 { // stack becomes empty or stack top is already filled with negative asteroids
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
