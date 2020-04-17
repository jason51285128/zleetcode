package zleetcode

import "testing"

type findmCase struct {
	input  []int
	expect int
	got    int
}

func (c *findmCase) run() bool {
	c.got = findm(c.input)
	if c.got == c.expect {
		return true
	}
	return false
}

var findmCases = []findmCase{
	{input: []int{1, 2, 4, 3}, expect: 2},
	{input: []int{1, 3, 2}, expect: 1},
	{input: []int{1, 2, 10, 8, 5, 4, 3, 9, 11, 14, 13}, expect: 2},
}

func TestFindm(t *testing.T) {
	for _, c := range findmCases {
		c.run()
		if c.got != c.expect {
			t.Error(c)
			return
		}
	}
}
