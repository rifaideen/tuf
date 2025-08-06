package main

import "testing"

func Test_ladderLength(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		beginWord string
		endWord   string
		wordList  []string
		want      int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ladderLength(tt.beginWord, tt.endWord, tt.wordList)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("ladderLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
