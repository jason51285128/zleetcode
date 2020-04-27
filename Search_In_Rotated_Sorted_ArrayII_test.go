package zleetcode

import "testing"

type searchInRotatedSortedArrayIICase struct {
	input1 []int
	input2 int
	output int
	expect int
	result bool
}

func (c *searchInRotatedSortedArrayIICase) run() {
	c.output = searchInRotatedSortedArrayII(c.input1, c.input2)
	if c.output == c.expect {
		c.result = true
	}
}

var searchInRotatedSortedArrayIICases = []searchInRotatedSortedArrayIICase{
	{input1: []int{}, input2: 1, expect: -1},
	{input1: []int{1}, input2: 1, expect: 0},
	{input1: []int{1}, input2: 2, expect: -1},
	{input1: []int{1, 3, 1, 1, 1}, input2: 3, expect: 1},
	{input1: []int{3, 4, 5, 5, 6, 7, 8, 9, 0, 1, 1, 2}, input2: 5, expect: 2},
	{input1: []int{3, 4, 5, 5, 6, 7, 8, 9, 0, 1, 1, 2}, input2: 1, expect: 10},
	{input1: []int{3, 4, 5, 5, 6, 7, 8, 9, 0, 1, 1, 2}, input2: 10, expect: -1},
}

func TestSearchInRotatedSortedArrayII(t *testing.T) {
	for _, c := range searchInRotatedSortedArrayIICases {
		if c.run(); c.output != c.expect {
			t.Error("input: ", c.input1, " target: ", c.input2,
				" expect: ", c.expect, " got: ", c.output)
		}
	}
}
