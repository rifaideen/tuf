package main

// using brute-force method
func ReversePair(nums []int) int {
	ans := 0

	for i := range len(nums) - 1 {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > 2*nums[j] {
				ans++
			}
		}
	}

	return ans
}
