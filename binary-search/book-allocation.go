package main

func FindPages(books []int, students int) int {
	// handle edge case, when we don't have enough books for students
	if students > len(books) {
		return -1
	}

	left := 0
	right := 0

	// select the max and sum of books
	for _, page := range books {
		left = max(left, page)
		right += page
	}

	for left <= right {
		mid := (left + right) / 2

		// check if we can allocate the books to all students
		if allocate(books, mid) <= students { // we can allocate
			right = mid - 1 // find the min number of pages we can allocate
		} else {
			left = mid + 1 // we can't allocate books, move the pointer to the right side
		}
	}

	return left
}

func allocate(books []int, pages int) int {
	allocations := 1 // we are calulating the allocations from 1st student
	count := 0

	for _, page := range books {
		if count+page > pages {
			allocations++
			count = page
		} else {
			count += page
		}
	}

	return allocations
}
