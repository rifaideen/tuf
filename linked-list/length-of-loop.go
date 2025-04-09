package main

func lengthOfLoop(head *Node) int {
	// detect the loop first using tortoise & hare algorithm
	slow, fast := head, head

	// move the pointers until the fast.next reaches nil
	for fast != nil && fast.next != nil {
		slow = slow.next      // move one step at a time
		fast = fast.next.next // move two steps at a time

		if fast == slow { // loop detected

			// find the length of the loop by moving the fast pointer by one and keep the slow in same place
			return func() int {
				fast = fast.next
				count := 1

				// until they meet, increment the counter
				for slow != fast {
					fast = fast.next
					count++
				}

				return count
			}()
		}
	}

	// no loop detected
	return 0
}
