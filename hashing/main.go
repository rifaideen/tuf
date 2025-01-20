package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(
		"Count frequency = ",
		CountFrequency([]int{10, 5, 10, 10, 10, 5, 20, 20, 20, 20, 20}),
	)

	high, low := Frequency([]int{10, 5, 10, 10, 10, 5, 20, 20, 20, 20, 20})
	fmt.Printf("High frequency = %d\nLow frequency = %d\n", high, low)

	matches := [][]int{
		{1, 3},
		{2, 3},
		{3, 6},
		{5, 6},
		{5, 7},
		{4, 5},
		{4, 8},
		{4, 9},
		{10, 4},
		{10, 9},
	}
	findWinners(matches)

	fmt.Println("Contiguos Array of [0,1,0] is", ContiguosArray([]int{0, 1, 0}))

	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	fmt.Println(GroupAnagram(strs))
}

func CountFrequency(nums []int) (frequency map[int]int) {
	frequency = map[int]int{}

	for _, num := range nums {
		frequency[num] += 1
	}

	return
}

func Frequency(nums []int) (high, low int) {
	frequency := map[int]int{}

	for _, num := range nums {
		frequency[num] += 1
	}

	h := 0
	l := math.MaxInt

	for key, val := range frequency {
		if val > h {
			h = val
			high = key
		}

		if val < l {
			l = val
			low = key
		}
	}

	return
}
