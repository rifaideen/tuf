package main

import "testing"

func Test_getIntersectionNode(t *testing.T) {
	// Helper function to create linked lists
	createList := func(nums []int) *Node {
		if len(nums) == 0 {
			return nil
		}
		head := &Node{data: nums[0]}
		current := head
		for i := 1; i < len(nums); i++ {
			current.next = &Node{data: nums[i]}
			current = current.next
		}
		return head
	}

	// Helper function to find a node in a linked list by its value
	findNode := func(head *Node, val int) *Node {
		current := head
		for current != nil {
			if current.data == val {
				return current
			}
			current = current.next
		}
		return nil
	}

	tests := []struct {
		name         string
		listA        []int
		listB        []int
		intersectVal int
		want         int
	}{
		{
			name:         "Intersection at node 8",
			listA:        []int{4, 1, 8, 4, 5},
			listB:        []int{5, 6, 1, 8, 4, 5},
			intersectVal: 8,
			want:         8,
		},
		{
			name:         "No intersection",
			listA:        []int{1, 9, 1, 2, 4},
			listB:        []int{3, 2, 9},
			intersectVal: 0,
			want:         0,
		},
		{
			name:         "Intersection at the beginning of list A",
			listA:        []int{2, 6, 4},
			listB:        []int{1, 5, 2, 6, 4},
			intersectVal: 2,
			want:         2,
		},
		{
			name:         "Intersection at the beginning of list B",
			listA:        []int{1, 5, 2, 6, 4},
			listB:        []int{2, 6, 4},
			intersectVal: 2,
			want:         2,
		},
		{
			name:         "One list is empty",
			listA:        []int{},
			listB:        []int{1, 2, 3},
			intersectVal: 0,
			want:         0,
		},
		{
			name:         "Both lists are empty",
			listA:        []int{},
			listB:        []int{},
			intersectVal: 0,
			want:         0,
		},
		{
			name:         "Intersection at the end",
			listA:        []int{1, 2, 3},
			listB:        []int{4, 5, 3},
			intersectVal: 3,
			want:         3,
		},
		{
			name:         "Identical lists",
			listA:        []int{1, 2, 3},
			listB:        []int{1, 2, 3},
			intersectVal: 1,
			want:         1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headA := createList(tt.listA)
			headB := createList(tt.listB)

			var intersectNode *Node
			if tt.intersectVal != 0 {
				intersectNode = findNode(headA, tt.intersectVal)
				if intersectNode != nil {
					// Find the node in list B where the intersection should occur
					var currentB *Node = headB
					if len(tt.listB) > 0 {
						for currentB.next != nil && findNode(currentB.next, tt.intersectVal) == nil {
							currentB = currentB.next
						}
						if currentB != nil {
							currentB.next = intersectNode
						} else if findNode(headB, tt.intersectVal) != nil {
							// Intersection at the head of B
							// No need to modify, headB already points to it (or will after creation)
						}
					} else if len(tt.listA) > 0 && findNode(headA, tt.intersectVal) == headA {
						// If list B is empty and intersection is at head of A
						headB = headA
					}
				}
			}

			got := getIntersectionNode(headA, headB)

			var gotVal int
			if got != nil {
				gotVal = got.data
			}

			if gotVal != tt.want {
				t.Errorf("getIntersectionNode() = %v, want %v", gotVal, tt.want)
			}
		})
	}
}
