package passphrase

import (
	"sort"
	"strings"

	set "gopkg.in/fatih/set.v0"
)

// PartA counts the number of valid passphrases
func PartA(phrases []string) int {
	return Counter(phrases, Valid)
}

// PartB counts the number of valid passphrases without allowing anagrams
func PartB(phrases []string) int {
	return Counter(phrases, ValidB)
}

// Counter does the counting in an asynchronous method
func Counter(phrases []string, accumulator func(string) bool) int {
	ch := make(chan int, len(phrases))
	num := 0
	for _, pp := range phrases {
		go func(phrase string) {
			if accumulator(phrase) {
				ch <- 1
			} else {
				ch <- 0
			}
		}(pp)
	}
	for range phrases {
		num += <-ch
	}
	return num
}

// Valid tells if a passphrase is completely unique
func Valid(passphrase string) bool {
	list := strings.Fields(passphrase)
	s := set.New()
	for _, word := range list {
		if s.Has(word) {
			return false
		}
		s.Add(word)
	}
	return true
}

// ValidB does not allow anagrams of each other
func ValidB(passphrase string) bool {
	list := strings.Fields(passphrase)
	s := set.New()
	for _, word := range list {
		x := []byte(word)
		w := make([]string, len(word))
		for _, char := range x {
			w = append(w, string(char))
		}
		sort.Strings(w)
		res := strings.Join(w, "")
		if s.Has(res) {
			return false
		}
		s.Add(res)
	}
	return true
}
