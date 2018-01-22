// Package knothash solves Day 10 of AOC 2017
package knothash

import (
	"fmt"
)

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

// moveB also returns a number of steps taken
func moveB(cmds, s []int, starting int) ([]int, int) {
	steps := 0
	for idx, cmd := range cmds {
		s = reverse(s, cmd)
		s = stepForward(s, starting+idx+cmd)
		steps += idx + cmd + starting
	}
	// s = stepForward(s, len(s)-steps%len(s))
	return s, steps
}

// PartB finds the dense hash
func PartB(cmds string) string {
	s := []int{}
	for i := 0; i <= 255; i++ {
		s = append(s, i)
	}
	seq := determineSequence(cmds)
	var steps int
	var st int
	for i := 0; i < 64; i++ {
		s, st = moveB(seq, s, i*len(seq))
		steps += st
	}
	s = stepForward(s, len(s)-steps%len(s))
	remaining := combineXOR(s)
	return convertToHex(remaining)
}

func determineSequence(cmds string) []int {
	seq := []int{}
	for _, c := range cmds {
		seq = append(seq, int(c))
	}
	seq = append(seq, []int{17, 31, 73, 47, 23}...)
	return seq
}

func combineXOR(orig []int) []int {
	results := []int{}
	for i := 0; i < 16; i++ {
		results = append(results, xor(orig[i*16:(i+1)*16]))
	}
	return results
}

func xor(orig []int) int {
	curr := orig[0]
	for j := 1; j < 16; j++ {
		curr = curr ^ orig[j]
	}
	return curr
}

func convertToHex(res []int) string {
	ret := ""
	for _, i := range res {
		ret += fmt.Sprintf("%.2x", i)
	}
	return ret
}
