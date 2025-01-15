package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(nums []int, target int) []int {
	visited := make(map[int]int) // value: index pair

	for i := range nums {
		j, ok := visited[target-nums[i]]

		if ok {
			return []int{j, i}
		}

		visited[nums[i]] = i
	}

	return []int{-1, -1}
}
