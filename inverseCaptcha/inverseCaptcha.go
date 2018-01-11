package inverseCaptcha

import (
	"log"
	"strconv"
	"time"
)

// InverseCaptcha calculates the inverse captcha for a string.
func InverseCaptcha(s string) int {
	s = s + string(s[0])
	start := time.Now()
	total := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			val, err := strconv.Atoi(string(s[i]))
			if err != nil {
				panic(err)
			}
			total += val
		}
	}
	log.Printf("Took %f seconds", time.Now().Sub(start).Seconds())
	return total
}

// InverseCaptchaVersion2 also considers the halfway round character
func InverseCaptchaVersion2(s string) int {
	total := 0
	newString := s + s
	half := len(s) / 2
	for i := 0; i < len(s); i++ {
		wanted := i + half
		if s[i] == newString[wanted] {
			val, err := strconv.Atoi(string(s[i]))
			if err != nil {
				panic(err)
			}
			total += val
		}
	}
	return total
}
