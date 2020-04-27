package zleetcode

import "testing"

type removeDuplicatesIICase struct {
	input  []int
	output []int
	expect []int
	result bool
}

func (c *removeDuplicatesIICase) run() {
	c.output = removeDuplicatesII(c.input)
	if len(c.output) != len(c.expect) {
		c.result = false
	}
	for i := range c.output {
		if c.output[i] != c.expect[i] {
			c.result = false
			return
		}
	}
	c.result = true
}

var removeDuplicatesIICases = []removeDuplicatesIICase{
	{input: []int{}, expect: []int{}},
	{input: []int{1}, expect: []int{1}},
	{input: []int{1, 1}, expect: []int{1, 1}},
	{input: []int{1, 1, 2}, expect: []int{1, 1, 2}},
	{
		input:  []int{1, 2, 3, 3, 4, 4, 4, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6},
		expect: []int{1, 2, 3, 3, 4, 4, 5, 5, 6, 6},
	},
}

func TestRemoveDuplicatesII(t *testing.T) {
	for _, c := range removeDuplicatesIICases {
		if c.run(); !c.result {
			t.Error("input: ", c.input, "expect: ", c.expect, "got: ", c.output)
			return
		}
	}
}
