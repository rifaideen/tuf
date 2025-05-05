package main

import "testing"

func Test_wordBreak(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s        string
		wordDict []string
		want     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := wordBreak(tt.s, tt.wordDict)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("wordBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}
