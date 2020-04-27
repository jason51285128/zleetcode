package zleetcode

import "testing"

type removeDuplicatesCase struct {
	input  []int
	output []int
	expect []int
	result bool
}

func (c *removeDuplicatesCase) run() {
	c.output = removeDuplicates(c.input)
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

var removeDuplicatesCases = []removeDuplicatesCase{
	{input: []int{}, expect: []int{}},
	{input: []int{1}, expect: []int{1}},
	{input: []int{1, 1, 2}, expect: []int{1, 2}},
	{input: []int{1, 2, 3, 3, 4, 4, 4, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6}, expect: []int{1, 2, 3, 4, 5, 6}},
}

func TestRemoveDuplicates(t *testing.T) {
	for _, c := range removeDuplicatesCases {
		if c.run(); !c.result {
			t.Error("input: ", c.input, "expect: ", c.expect, "got: ", c.output)
			return
		}
	}
}
