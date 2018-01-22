// Package stream solves Day 9!
package stream

func removeCancels(s string) string {
	for idx, char := range s {
		if char == '!' {
			return removeCancels(s[:idx] + s[idx+2:])
		}
	}
	return s
}

// removeGarbage removes
func removeGarbage(s string) string {
	startIdx := -1
	for idx, char := range s {
		if char == '<' && startIdx == -1 {
			startIdx = idx
		} else if char == '>' && startIdx != -1 {
			s = s[0:startIdx] + s[idx+1:]
			return removeGarbage(s)
		}
	}
	return s
}

func score(s string, base int) int {
	startIdx := -1
	openCount := 0
	for idx, char := range s {
		if char == '{' && startIdx == -1 {
			startIdx = idx
			openCount++
		} else if char == '{' {
			openCount++
		} else if char == '}' && startIdx != -1 {
			openCount--
			if openCount == 0 {
				return score(s[startIdx+1:idx], base+1) + score(s[idx+1:], base) + base
			}
		}
	}
	return 0
}

// PartA calculates the score for part A!
func PartA(s string) int {
	return score(removeGarbage(removeCancels(s)), 1)
}

func removeGarbageWithCount(s string) (string, int) {
	startIdx := -1
	for idx, char := range s {
		if char == '<' && startIdx == -1 {
			startIdx = idx
		} else if char == '>' && startIdx != -1 {
			s = s[0:startIdx] + s[idx+1:]
			str, count := removeGarbageWithCount(s)
			return str, count + idx - startIdx - 1
		}
	}
	return s, 0
}

// PartB counts the number of rubbish characters
func PartB(s string) int {
	_, c := removeGarbageWithCount(removeCancels(s))
	return c
}
