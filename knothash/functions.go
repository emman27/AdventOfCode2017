// Package knothash solves Day 10 of AOC 2017
package knothash

import "github.com/sirupsen/logrus"

type store struct {
	Items []int
	Curr  int
}

func (s *store) stepForward(steps int) {
	s.Curr = (s.Curr + steps) % len(s.Items)
}

func (s *store) reverse(l int) {
	var max int
	var right int
	if m := s.Curr + l; m >= len(s.Items) {
		max = m % len(s.Items)
		right = len(s.Items)
	} else {
		right = s.Curr + l
	}
	correctLength := len(s.Items)
	reversed := reverseHelper(append(s.Items[s.Curr:right], s.Items[:max]...))
	s.Items = append(append(s.Items[:s.Curr], reversed...), s.Items[right:]...)
	s.Items = append(reversed[len(reversed)-max:], s.Items...)[:correctLength]
	logrus.Info(s.Items)
}

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
