package main

import "testing"

func Test_removeKdigits(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		num  string
		k    int
		want string
	}{
		{
			num:  "1432219",
			k:    3,
			want: "1219",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeKdigits(tt.num, tt.k)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("removeKdigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
