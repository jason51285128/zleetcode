package zleetcode

import "testing"

type searchInRotatedSortedArrayCase struct {
	input1 []int
	input2 int
	output int
	expect int
	result bool
}

func (c *searchInRotatedSortedArrayCase) run() {
	c.output = searchInRotatedSortedArray(c.input1, c.input2)
	if c.output == c.expect {
		c.result = true
	}
}

var searchInRotatedSortedArrayCases = []searchInRotatedSortedArrayCase{
	{input1: []int{}, input2: 1, expect: -1},
	{input1: []int{1}, input2: 1, expect: 0},
	{input1: []int{1}, input2: 2, expect: -1},
	{input1: []int{3, 4, 5, 6, 7, 8, 9, 0, 1, 2}, input2: 4, expect: 1},
	{input1: []int{3, 4, 5, 6, 7, 8, 9, 0, 1, 2}, input2: 1, expect: 8},
	{input1: []int{3, 4, 5, 6, 7, 8, 9, 0, 1, 2}, input2: 10, expect: -1},
}

func TestSearchInRotatedSortedArray(t *testing.T) {
	for _, c := range searchInRotatedSortedArrayCases {
		if c.run(); c.output != c.expect {
			t.Error("input: ", c.input1, " target: ", c.input2,
				" expect: ", c.expect, " got: ", c.output)
		}
	}
}
