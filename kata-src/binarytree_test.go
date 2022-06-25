package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvertBinaryTree(t *testing.T) {
	tt := []struct {
		name   string
		input  *TreeNode
		output []TreeNode
	}{
		{
			name:   "empty",
			input:  nil,
			output: []TreeNode{},
		},
		{
			name:   "one",
			input:  &TreeNode{Val: 1},
			output: []TreeNode{TreeNode{Val: 1}},
		},
		{
			name:   "two",
			input:  &TreeNode{Val: 2, Left: &TreeNode{Val: 1}},
			output: []TreeNode{TreeNode{Val: 1}, TreeNode{Val: 2}},
		},
		{
			name:   "three",
			input:  &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}},
			output: []TreeNode{TreeNode{Val: 1}, TreeNode{Val: 2}, TreeNode{Val: 3}},
		},
		{
			name:   "four",
			input:  &TreeNode{Val: 4, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}}},
			output: []TreeNode{TreeNode{Val: 1}, TreeNode{Val: 2}, TreeNode{Val: 3}, TreeNode{Val: 4}},
		},
		{
			name:   "five",
			input:  &TreeNode{Val: 5, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}}}},
			output: []TreeNode{TreeNode{Val: 1}, TreeNode{Val: 2}, TreeNode{Val: 3}, TreeNode{Val: 4}, TreeNode{Val: 5}},
		},
		{
			name:   "six",
			input:  &TreeNode{Val: 6, Left: &TreeNode{Val: 5, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}}}}},
			output: []TreeNode{TreeNode{Val: 1}, TreeNode{Val: 2}, TreeNode{Val: 3}, TreeNode{Val: 4}, TreeNode{Val: 5}, TreeNode{Val: 6}},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			assert.EqualValues(t, tc.output, invertbinarytree(tc.input))
		})
	}
}
