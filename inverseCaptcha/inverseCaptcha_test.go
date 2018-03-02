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

func TestInverseCaptchaVersion2(t *testing.T) {
	cases := []testCase{
		{Input: "1212", Output: 6},
		{Input: "1221", Output: 0},
		{Input: "123425", Output: 4},
		{Input: "123123", Output: 12},
		{Input: "12131415", Output: 4},
	}
	for _, c := range cases {
		actual := Version2(c.Input)
		if actual != c.Output {
			t.Fatalf("Test failed, Actual: %d, Expected %d", actual, c.Output)
		}
	}
}
