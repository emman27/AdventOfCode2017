package passphrase

import (
	"testing"
)

func TestValid(t *testing.T) {
	type testCase struct {
		Input  string
		Output bool
	}

	cases := []testCase{
		{Input: "aa bb cc dd ee", Output: true},
		{Input: "aa bb cc dd aa", Output: false},
		{Input: "aa bb cc dd aaa", Output: true},
	}
	for _, c := range cases {
		if res := Valid(c.Input); res != c.Output {
			t.Fatalf("Want: %v, Got: %v", c.Output, res)
		}
	}
}

func TestValidB(t *testing.T) {
	type testCase struct {
		Input  string
		Output bool
	}

	cases := []testCase{
		{Input: "aa bb cc dd ee", Output: true},
		{Input: "aa bb cc dd aa", Output: false},
		{Input: "aa bb cc dd aaa", Output: true},
		{Input: "abcde fghij", Output: true},
		{Input: "abcde xyz ecdab", Output: false},
		{Input: "a ab abc abd abf abj", Output: true},
		{Input: "iiii oiii ooii oooo", Output: true},
		{Input: "oiii ioii iioi iiio", Output: false},
	}
	for _, c := range cases {
		if res := ValidB(c.Input); res != c.Output {
			t.Fatalf("Want: %v, Got: %v", c.Output, res)
		}
	}
}
