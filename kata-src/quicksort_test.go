package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuicksort(t *testing.T) {
	tt := []struct {
		name   string
		input  []int
		output []int
	}{
		{
			name:   "empty",
			input:  []int{},
			output: []int{},
		},
		{
			name:   "one",
			input:  []int{1},
			output: []int{1},
		},
		{
			name:   "two",
			input:  []int{2, 1},
			output: []int{1, 2},
		},
		{
			name:   "three",
			input:  []int{3, 2, 1},
			output: []int{1, 2, 3},
		},
		{
			name:   "four",
			input:  []int{4, 3, 2, 1},
			output: []int{1, 2, 3, 4},
		},
		{
			name:   "five",
			input:  []int{5, 4, 3, 2, 1},
			output: []int{1, 2, 3, 4, 5},
		},
		{
			name:   "six",
			input:  []int{6, 5, 4, 3, 2, 1},
			output: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			assert.EqualValues(t, tc.output, quicksort(tc.input))
		})
	}
}
