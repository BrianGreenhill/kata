package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTravelingSalesman(t *testing.T) {
	tt := []struct {
		name   string
		input  [][]int
		output int
	}{
		{
			name:   "empty",
			input:  [][]int{},
			output: 0,
		},
		{
			name: "one",
			input: [][]int{
				{0, 1, 2},
				{1, 2, 3},
				{2, 3, 4},
			},
			output: 3,
		},
		{
			name: "two",
			input: [][]int{
				{0, 1, 2},
				{1, 2, 3},
				{2, 3, 4},
				{3, 4, 0},
			},
			output: 4,
		},
		{
			name: "three",
			input: [][]int{
				{0, 1, 2},
				{1, 2, 3},
				{2, 3, 4},
				{3, 4, 0},
				{4, 0, 1},
			},
			output: 5,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			assert.EqualValues(t, tc.output, travelingSalesman(tc.input))
		})
	}

}
