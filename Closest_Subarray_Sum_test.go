package zleetcode

import "testing"

var (
	closestSubarraySumCase1Input1 = []int{9, 9, 9}
	closestSubarraySumCase1Input2 = 8
	closestSubarraySumCase1Expect = []int{-1, -1, -1}

	closestSubarraySumCase2Input1 = []int{1, 2, 3, 8}
	closestSubarraySumCase2Input2 = 8
	closestSubarraySumCase2Expect = []int{8, 3, 3}

	closestSubarraySumCase3Input1 = []int{1, 2, 6, 8, 7, 5}
	closestSubarraySumCase3Input2 = 11
	closestSubarraySumCase3Expect = []int{9, 0, 2}

	closestSubarraySumCase4Input1 = []int{1, 1, 1, 1}
	closestSubarraySumCase4Input2 = 4
	closestSubarraySumCase4Expect = []int{4, 0, 3}

	closestSubarraySumCase5Input1 = []int{5, 5, 5, 5}
	closestSubarraySumCase5Input2 = 4
	closestSubarraySumCase5Expect = []int{-1, -1, -1}

	closestSubarraySumCase6Input1 = []int{1, 2, 6, 8, 2, 7, 5}
	closestSubarraySumCase6Input2 = 11
	closestSubarraySumCase6Expect = []int{10, 3, 4}

	closestSubarraySumCase7Input1 = []int{1, 2, 6, 8, 2, 11, 7, 5}
	closestSubarraySumCase7Input2 = 11
	closestSubarraySumCase7Expect = []int{11, 5, 5}
)

func TestClosestSubarraySumCase1(t *testing.T) {
	res := closestSubarraySum(closestSubarraySumCase1Input1, closestSubarraySumCase1Input2)
	for i := range closestSubarraySumCase1Expect {
		if closestSubarraySumCase1Expect[i] != res[i] {
			t.Error("input: ", closestSubarraySumCase1Input1, " ", closestSubarraySumCase1Input2,
				"expected: ", closestSubarraySumCase1Expect,
				"got: ", res)
			break
		}
	}
}

func TestClosestSubarraySumCase2(t *testing.T) {
	res := closestSubarraySum(closestSubarraySumCase2Input1, closestSubarraySumCase2Input2)
	for i := range closestSubarraySumCase2Expect {
		if closestSubarraySumCase2Expect[i] != res[i] {
			t.Error("input: ", closestSubarraySumCase2Input1, " ", closestSubarraySumCase2Input2,
				"expected: ", closestSubarraySumCase2Expect,
				"got: ", res)
			break
		}
	}
}

func TestClosestSubarraySumCase3(t *testing.T) {
	res := closestSubarraySum(closestSubarraySumCase3Input1, closestSubarraySumCase3Input2)
	for i := range closestSubarraySumCase3Expect {
		if closestSubarraySumCase3Expect[i] != res[i] {
			t.Error("input: ", closestSubarraySumCase3Input1, " ", closestSubarraySumCase3Input2,
				"expected: ", closestSubarraySumCase3Expect,
				"got: ", res)
			break
		}
	}
}

func TestClosestSubarraySumCase4(t *testing.T) {
	res := closestSubarraySum(closestSubarraySumCase4Input1, closestSubarraySumCase4Input2)
	for i := range closestSubarraySumCase4Expect {
		if closestSubarraySumCase4Expect[i] != res[i] {
			t.Error("input: ", closestSubarraySumCase4Input1, " ", closestSubarraySumCase4Input2,
				"expected: ", closestSubarraySumCase4Expect,
				"got: ", res)
			break
		}
	}
}

func TestClosestSubarraySumCase5(t *testing.T) {
	res := closestSubarraySum(closestSubarraySumCase5Input1, closestSubarraySumCase5Input2)
	for i := range closestSubarraySumCase5Expect {
		if closestSubarraySumCase5Expect[i] != res[i] {
			t.Error("input: ", closestSubarraySumCase5Input1, " ", closestSubarraySumCase5Input2,
				"expected: ", closestSubarraySumCase5Expect,
				"got: ", res)
		}
	}
}

func TestClosestSubarraySumCase6(t *testing.T) {
	res := closestSubarraySum(closestSubarraySumCase6Input1, closestSubarraySumCase6Input2)
	for i := range closestSubarraySumCase6Expect {
		if closestSubarraySumCase6Expect[i] != res[i] {
			t.Error("input: ", closestSubarraySumCase6Input1, " ", closestSubarraySumCase6Input2,
				"expected: ", closestSubarraySumCase6Expect,
				"got: ", res)
		}
	}
}

func TestClosestSubarraySumCase7(t *testing.T) {
	res := closestSubarraySum(closestSubarraySumCase7Input1, closestSubarraySumCase7Input2)
	for i := range closestSubarraySumCase7Expect {
		if closestSubarraySumCase7Expect[i] != res[i] {
			t.Error("input: ", closestSubarraySumCase7Input1, " ", closestSubarraySumCase7Input2,
				"expected: ", closestSubarraySumCase7Expect,
				"got: ", res)
		}
	}
}
