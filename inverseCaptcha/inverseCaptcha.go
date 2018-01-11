package inverseCaptcha

import "strconv"

// InverseCaptcha calculates the inverse captcha for a string.
func InverseCaptcha(s string) int {
	total := 0
	for i := 0; i < len(s); i++ {
		if i+1 == len(s) {
			if s[i] == s[0] {
				val, err := strconv.Atoi(string(s[i]))
				if err != nil {
					panic(err)
				}
				total += val
			}
		} else if s[i] == s[i+1] {
			val, err := strconv.Atoi(string(s[i]))
			if err != nil {
				panic(err)
			}
			total += val
		}
	}
	return total
}
