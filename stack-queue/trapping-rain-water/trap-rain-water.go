package main

func trap(height []int) int {
	n := len(height)
	ans := 0

	// left and right pointers
	left := 0
	right := n - 1

	// store left max and right max heights
	leftMax := 0
	rightMax := 0

	// loop until left and right are less or equal
	for left <= right {
		// see if we can store between current and right side building
		if height[left] <= height[right] {
			// if current height is greater than existing left max height, overwrite with current one
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				ans += leftMax - height[left]
			}

			// move left pointer
			left++
		} else {
			// if current height is greater than existing right max height, overwrite with current one
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				ans += rightMax - height[right]
			}

			// move right pointer
			right--
		}
	}

	return ans
}
