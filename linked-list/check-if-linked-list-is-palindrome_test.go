package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		head *Node
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPalindrome(tt.head)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindromeOptimized(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		head *Node
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPalindromeOptimized(tt.head)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("isPalindromeOptimized() = %v, want %v", got, tt.want)
			}
		})
	}
}
