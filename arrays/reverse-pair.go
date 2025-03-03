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

func ReversePairs(nums []int) int {
	return rmergesort(nums, 0, len(nums)-1)
}

func rmergesort(nums []int, left, right int) int {
	if left >= right {
		return 0
	}

	count := 0
	mid := (left + right) / 2
	count += rmergesort(nums, left, mid)
	count += rmergesort(nums, mid+1, right)
	count += countPairs(nums, left, mid, right)
	rmerge(nums, left, mid, right)

	return count
}

func countPairs(nums []int, low, mid, high int) int {
	right := mid + 1
	count := 0

	for i := low; i <= mid; i++ {
		for right <= high && nums[i] > (2*nums[right]) {
			right++
		}

		count += right - (mid + 1)
	}

	return count
}

func rmerge(nums []int, low, mid, high int) {
	ans := []int{}
	left := low
	right := mid + 1

	for left <= mid && right <= high {
		if nums[left] <= nums[right] {
			ans = append(ans, nums[left])
			left++
		} else {
			ans = append(ans, nums[right])
			right++
		}
	}

	for left <= mid {
		ans = append(ans, nums[left])
		left++
	}

	for right <= high {
		ans = append(ans, nums[right])
		right++
	}

	for i := low; i <= high; i++ {
		nums[i] = ans[i-low]
	}
}
