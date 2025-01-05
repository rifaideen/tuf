package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(nums []int, target int) []int {
	visited := map[int]int{}

	for i := 0; i < len(nums); i++ {
		current := nums[i]
		x := target - current

		if index, ok := visited[x]; ok {
			return []int{index, i}
		}

		visited[current] = i
	}

	return []int{}
}
