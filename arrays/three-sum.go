package main

import "slices"

func ThreeSum(nums []int) [][]int {
	ans := [][]int{}

	// sort nums
	slices.Sort(nums)
	n := len(nums)

	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j := i + 1
		k := n - 1

		for j < k {
			total := nums[i] + nums[j] + nums[k]

			if total == 0 {
				// move pointers and push it to answer
				ans = append(ans, []int{
					nums[i],
					nums[j],
					nums[k],
				})

				j++
				k--

				for j < k && nums[j] == nums[j-1] {
					j++
				}

				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if total > 0 {
				k--
			} else if total < 0 {
				j++
			}
		}
	}

	return ans
}
