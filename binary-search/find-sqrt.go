package main

func FindSqrt(num int) int {
	ans := -1
	left := 0
	right := num

	for left <= right {
		mid := (left + right) / 2

		sqr := mid * mid

		if sqr == num {
			ans = mid
			break
		}

		if sqr < num {
			left = mid + 1
			ans = mid
		} else {
			right = mid - 1
		}
	}

	return ans
}
