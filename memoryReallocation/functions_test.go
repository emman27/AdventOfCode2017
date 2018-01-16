package memoryReallocation

import "testing"

func TestReallocate(t *testing.T) {
	type testCase struct {
		Input  []int
		Output []int
	}
	cases := []testCase{
		testCase{Input: []int{0, 2, 7, 0}, Output: []int{2, 4, 1, 2}},
		testCase{Input: []int{2, 4, 1, 2}, Output: []int{3, 1, 2, 3}},
		testCase{Input: []int{3, 1, 2, 3}, Output: []int{0, 2, 3, 4}},
		testCase{Input: []int{0, 2, 3, 4}, Output: []int{1, 3, 4, 1}},
		testCase{Input: []int{1, 3, 4, 1}, Output: []int{2, 4, 1, 2}},
	}
	for _, c := range cases {
		res := Reallocate(c.Input)
		for idx, val := range res {
			if val != c.Output[idx] {
				t.Fatalf("Failed. Got %v, Want %v", res, c.Output)
			}
		}
	}
}

func TestPartA(t *testing.T) {
	type testCase struct {
		Input  []int
		Output int
	}
	cases := []testCase{
		testCase{Input: []int{0, 2, 7, 0}, Output: 5},
	}
	for _, c := range cases {
		if res := PartA(c.Input); res != c.Output {
			t.Fatalf("Failed. Got %d, Want %d", res, c.Output)
		}
	}
}

func TestPartB(t *testing.T) {
	type testCase struct {
		Input  []int
		Output int
	}
	cases := []testCase{
		testCase{Input: []int{0, 2, 7, 0}, Output: 4},
	}
	for _, c := range cases {
		if res := PartB(c.Input); res != c.Output {
			t.Fatalf("Failed. Got %d, Want %d", res, c.Output)
		}
	}
}
