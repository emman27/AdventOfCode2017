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

func TestPartB(t *testing.T) {
	type testCase struct {
		Input    int
		Expected int
	}

	cases := []testCase{
		testCase{Input: 1, Expected: 2},
		testCase{Input: 2, Expected: 4},
		testCase{Input: 3, Expected: 4},
		testCase{Input: 4, Expected: 5},
		testCase{Input: 5, Expected: 10},
		testCase{Input: 9, Expected: 10},
		testCase{Input: 10, Expected: 11},
		testCase{Input: 11, Expected: 23},
	}
	for idx, c := range cases {
		if res := PartB(c.Input); res != c.Expected {
			t.Fatalf("Test case %d, failed, Input: %d, Expected: %d, Got: %d", idx, c.Input, c.Expected, res)
		}
	}
}

func TestPosition(t *testing.T) {
	type testCase struct {
		Input   int
		OutputX int
		OutputY int
	}
	cases := []testCase{
		testCase{Input: 1, OutputX: 0, OutputY: 0},
		testCase{Input: 2, OutputX: 1, OutputY: 0},
		testCase{Input: 3, OutputX: 1, OutputY: 1},
		testCase{Input: 4, OutputX: 0, OutputY: 1},
		testCase{Input: 5, OutputX: -1, OutputY: 1},
		testCase{Input: 6, OutputX: -1, OutputY: 0},
		testCase{Input: 7, OutputX: -1, OutputY: -1},
		testCase{Input: 8, OutputX: 0, OutputY: -1},
		testCase{Input: 9, OutputX: 1, OutputY: -1},
	}
	for _, c := range cases {
		if x, y := position(c.Input); x != c.OutputX || y != c.OutputY {
			t.Fatalf("Input %d, Expected (%d, %d) but Got (%d, %d)", c.Input, c.OutputX, c.OutputY, x, y)
		}
	}
}
