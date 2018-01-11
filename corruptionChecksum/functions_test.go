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
