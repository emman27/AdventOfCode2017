package inverseCaptcha

import "testing"

type testCase struct {
	Input  string
	Output int
}

func TestInverseCaptcha(t *testing.T) {
	cases := []testCase{
		{Input: "1122", Output: 3},
		{Input: "1111", Output: 4},
		{Input: "1234", Output: 0},
		{Input: "91212129", Output: 9},
	}
	for _, c := range cases {
		actual := InverseCaptcha(c.Input)
		if actual != c.Output {
			t.Fatalf("Test failed, Actual: %d, Expected %d", actual, c.Output)
		}
	}
}
