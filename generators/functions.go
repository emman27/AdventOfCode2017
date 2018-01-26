// Package generators solves day 15 of AoC 2017
package generators

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

var factorA = int64(16807)
var factorB = int64(48271)

func genNext(factor, prev int64) int64 {
	logrus.Info(factor, prev, factor*prev)
	logrus.Info(factor * prev % 2147283647)
	return factor * prev % 2147283647
}

// PartA counts the collisions
func PartA(a, b int64) int {
	count := 0
	var aBin string
	var bBin string
	for i := 0; i < 5; i++ {
		a = genNext(factorA, a)
		b = genNext(factorB, b)
		aBin = toBinary(a)[16:]
		bBin = toBinary(b)[16:]
		if aBin == bBin {
			count++
		}
	}
	return count
}

func toBinary(i int64) string {
	return fmt.Sprintf("%.32b", i)
}
