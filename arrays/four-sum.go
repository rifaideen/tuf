package main

import "slices"

func FourSum(nums []int, target int) [][]int {
	ans := [][]int{}

	n := len(nums)
	// sort the numbers
	slices.Sort(nums)

	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			k := j + 1
			l := n - 1

			for k < l {
				total := nums[i] + nums[j] + nums[k] + nums[l]

				if total == target {
					ans = append(ans, []int{
						nums[i],
						nums[j],
						nums[k],
						nums[l],
					})

					k++
					l--

					for k < l && nums[k] == nums[k-1] {
						k++
					}

					for k < l && nums[l] == nums[l+1] {
						l--
					}
				} else if total > target {
					l--
				} else if total < target {
					k++
				}
			}
		}
	}

	return ans
}
