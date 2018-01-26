// Package defrag solves Day 14 of AoC2017
package defrag

import (
	"fmt"
	"strconv"

	"github.com/emman27/aoc2017/knothash"
)

// PartA returns the number of used squares
func PartA(input string) int {
	count := 0
	for i := 0; i < 128; i++ {
		keystring := input + "-" + strconv.Itoa(i)
		denseHash := knothash.PartB(keystring)
		for _, char := range denseHash {
			bitMask := convertToBitmask(char)
			for _, c := range bitMask {
				if c == '1' {
					count++
				}
			}
		}
	}
	return count
}

func convertToBitmask(c rune) string {
	i, err := strconv.ParseInt(string(c), 16, 0)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%.4b", i)
}
