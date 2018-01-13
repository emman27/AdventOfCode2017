package spiralMemory

import "testing"

func TestPartA(t *testing.T) {
	type testCase struct {
		Input    int
		Expected int
	}

	cases := []testCase{
		testCase{Input: 1, Expected: 0},
		testCase{Input: 12, Expected: 3},
		testCase{Input: 23, Expected: 2},
		testCase{Input: 1024, Expected: 31},
	}
	for idx, c := range cases {
		if res := PartA(c.Input); res != c.Expected {
			t.Fatalf("Test case %d failed, Input: %d, Expected: %d, Got: %d", idx, c.Input, c.Expected, res)
		}
	}
}

func TestRingNumber(t *testing.T) {
	type testCase struct {
		Input    int
		Expected int
	}

	cases := []testCase{
		testCase{Input: 1, Expected: 0},
		testCase{Input: 2, Expected: 1},
		testCase{Input: 5, Expected: 1},
		testCase{Input: 9, Expected: 1},
		testCase{Input: 10, Expected: 2},
	}
	for idx, c := range cases {
		if res := ringNumber(c.Input); res != c.Expected {
			t.Fatalf("Test case %d failed, Input: %d, Expected: %d, Got: %d", idx, c.Input, c.Expected, res)
		}
	}
}
