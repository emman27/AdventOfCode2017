package corruptionChecksum

import "testing"

type testCase struct {
	Input    [][]int
	Expected int
}

func TestCalcuateMaxDiff(t *testing.T) {
	cases := []testCase{
		{Input: [][]int{[]int{5, 1, 9, 5}, []int{7, 5, 3}, []int{2, 4, 6, 8}}, Expected: 18},
	}
	for idx, c := range cases {
		if res := Calculate(c.Input, MaxDiff); res != c.Expected {
			t.Fatalf("Test #%d failed. Expected: %d, Actual: %d", idx, c.Expected, res)
		}
	}
}

type t2Case struct {
	Input    []int
	Expected int
}

func TestEvenDivisResults(t *testing.T) {
	cases := []t2Case{
		{Input: []int{5, 9, 2, 8}, Expected: 4},
		{Input: []int{9, 4, 7, 3}, Expected: 3},
		{Input: []int{3, 8, 6, 5}, Expected: 2},
	}
	for idx, c := range cases {
		if res, err := EvenDivisResults(c.Input); res != c.Expected || err != nil {
			t.Fatalf("Test #%d failed. Expected: %d, Actual: %d", idx, c.Expected, res)
		}
	}
}
