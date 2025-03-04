package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindRotatedCount(t *testing.T) {
	type testcase struct {
		input  []int
		output int
	}

	testcases := []testcase{
		{
			input:  []int{4, 5, 6, 7, 0, 1, 2, 3},
			output: 4,
		},
		{
			input:  []int{3, 4, 5, 1, 2},
			output: 3,
		},
		{
			input: []int{
				123599, 146544, 167713, 171529, 173758, 177264, 178937, 184403, 190836, 202013, 221088, 235369, 255064, 265870, 283759, 306962, 339797, 350910, 382851, 397079, 413843, 420673, 428991, 433410, 433712, 479635, 512424, 515402, 543673, 547017, 577958, 580560, 582120, 627947, 635471, 649604, 654079, 666601, 685972, 702400, 703050, 704948, 710230, 729866, 740006, 752919, 775434, 781915, 806724, 828173, 840647, 845691, 853036, 908771, 919220, 929581, 937378, 951577, 975151, 108346, 112959,
			},
			output: 59,
		},
	}

	for _, testcase := range testcases {
		output := FindRotatedCount(testcase.input)
		assert.Equal(t, testcase.output, output)
	}
}
