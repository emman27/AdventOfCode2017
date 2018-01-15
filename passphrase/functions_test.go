package passphrase

import "testing"

func TestValid(t *testing.T) {
	type testCase struct {
		Input  string
		Output bool
	}

	cases := []testCase{
		testCase{Input: "aa bb cc dd ee", Output: true},
		testCase{Input: "aa bb cc dd aa", Output: false},
		testCase{Input: "aa bb cc dd aaa", Output: true},
	}
	for _, c := range cases {
		if res := Valid(c.Input); res != c.Output {
			t.Fatalf("Want: %v, Got: %v", c.Output, res)
		}
	}
}
