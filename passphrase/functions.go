package passphrase

import (
	"strings"

	set "gopkg.in/fatih/set.v0"
)

// PartA counts the number of valid passphrases
func PartA(phrases []string) int {
	ch := make(chan int, len(phrases))
	num := 0
	for _, pp := range phrases {
		go func(phrase string) {
			if Valid(phrase) {
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
