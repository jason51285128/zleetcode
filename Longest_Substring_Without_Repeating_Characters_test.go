package zleetcode

import "testing"

const (
	inputcase1  = ""
	expectcase1 = 0
	inputcase2  = "x"
	expectcase2 = 1
	inputcase3  = "hello"
	expectcase3 = 3
	inputcase4  = "abcxdaefg"
	expectcase4 = 8
	inputcase5  = "abcdaefag"
	expectcase5 = 6
	inputcase6  = "abcdaefaglmnopq"
	expectcase6 = 10
	inputcase7  = "abcdefa"
	expectcase7 = 6
	inputcase8  = "abcdefgh"
	expectcase8 = 8
)

//greedy method testcase
func TestLswrcGreedy01(t *testing.T) {
	if res := lswrcGreedy(inputcase1); res != expectcase1 {
		t.Error("input: ", inputcase1, " expected: ", expectcase1, " got: ", res)
	}
}

func TestLswrcGreedy02(t *testing.T) {
	if res := lswrcGreedy(inputcase2); res != expectcase2 {
		t.Error("input: ", inputcase2, " expected: ", expectcase2, " got: ", res)
	}
}

func TestLswrcGreedy03(t *testing.T) {
	if res := lswrcGreedy(inputcase3); res != expectcase3 {
		t.Error("input: ", inputcase3, " expected: ", expectcase3, " got: ", res)
	}
}

func TestLswrcGreedy04(t *testing.T) {
	if res := lswrcGreedy(inputcase4); res != expectcase4 {
		t.Error("input: ", inputcase4, " expected: ", expectcase4, " got: ", res)
	}
}

func TestLswrcGreedy05(t *testing.T) {
	if res := lswrcGreedy(inputcase5); res != expectcase5 {
		t.Error("input: ", inputcase5, " expected: ", expectcase5, " got: ", res)
	}
}

func TestLswrcGreedy06(t *testing.T) {
	if res := lswrcGreedy(inputcase6); res != expectcase6 {
		t.Error("input: ", inputcase6, " expected: ", expectcase6, " got: ", res)
	}
}

func TestLswrcGreedy07(t *testing.T) {
	if res := lswrcGreedy(inputcase7); res != expectcase7 {
		t.Error("input: ", inputcase7, " expected: ", expectcase7, " got: ", res)
	}
}

func TestLswrcGreedy08(t *testing.T) {
	if res := lswrcGreedy(inputcase8); res != expectcase8 {
		t.Error("input: ", inputcase8, " expected: ", expectcase8, " got: ", res)
	}
}

//DP method testcase
func TestLswrcDP01(t *testing.T) {
	if res := lswrcDP(inputcase1); res != expectcase1 {
		t.Error("input: ", inputcase1, " expected: ", expectcase1, " got: ", res)
	}
}

func TestLswrcDP02(t *testing.T) {
	if res := lswrcDP(inputcase2); res != expectcase2 {
		t.Error("input: ", inputcase2, " expected: ", expectcase2, " got: ", res)
	}
}

func TestLswrcDP03(t *testing.T) {
	if res := lswrcDP(inputcase3); res != expectcase3 {
		t.Error("input: ", inputcase3, " expected: ", expectcase3, " got: ", res)
	}
}

func TestLswrcDP04(t *testing.T) {
	if res := lswrcDP(inputcase4); res != expectcase4 {
		t.Error("input: ", inputcase4, " expected: ", expectcase4, " got: ", res)
	}
}

func TestLswrcDP05(t *testing.T) {
	if res := lswrcDP(inputcase5); res != expectcase5 {
		t.Error("input: ", inputcase5, " expected: ", expectcase5, " got: ", res)
	}
}

func TestLswrcDP06(t *testing.T) {
	if res := lswrcDP(inputcase6); res != expectcase6 {
		t.Error("input: ", inputcase6, " expected: ", expectcase6, " got: ", res)
	}
}

func TestLswrcDP07(t *testing.T) {
	if res := lswrcDP(inputcase7); res != expectcase7 {
		t.Error("input: ", inputcase7, " expected: ", expectcase7, " got: ", res)
	}
}

func TestLswrcDP08(t *testing.T) {
	if res := lswrcDP(inputcase8); res != expectcase8 {
		t.Error("input: ", inputcase8, " expected: ", expectcase8, " got: ", res)
	}
}
