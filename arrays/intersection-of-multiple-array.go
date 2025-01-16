package main

import "sort"

func Intersection(nums [][]int) []int {
	counts := map[int]int{}
	size := len(nums)
	ans := []int{}

	for _, arr := range nums {
		for _, num := range arr {
			counts[num] += 1

			if counts[num] == size {
				ans = append(ans, num)
			}
		}
	}

	return sort.IntSlice(ans)
}
