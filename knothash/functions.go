// Package knothash solves Day 10 of AOC 2017
package knothash

func stepForward(s []int, steps int) []int {
	steps = steps % len(s)
	return append(s[steps:], s[:steps]...)
}

func reverse(s []int, l int) []int {
	return append(reverseHelper(s[:l]), s[l:]...)
}

func reverseHelper(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// PartA calculates the product of the first 2 elements at the end of the command list
func PartA(cmds []int) int {
	s := []int{}
	for i := 0; i <= 255; i++ {
		s = append(s, i)
	}
	return move(cmds, s)
}

func move(cmds, s []int) int {
	steps := 0
	for idx, cmd := range cmds {
		s = reverse(s, cmd)
		s = stepForward(s, idx+cmd)
		steps += idx + cmd
	}
	// Reset to front
	s = stepForward(s, len(s)-steps%len(s))
	return s[0] * s[1]
}
