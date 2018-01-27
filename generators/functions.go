package generators

import (
	"fmt"
)

type generator int64

var factorA generator = 16807
var factorB generator = 48271

func next(g generator, i int64) int64 {
	return int64(g) * i % 2147483647
}

// PartA calculates the collisiosn
func PartA(a, b int64) int {
	count := 0
	for i := 0; i < 40000000; i++ {
		a = next(factorA, a)
		b = next(factorB, b)
		binA := toBinary(a)
		binB := toBinary(b)
		if binA[16:] == binB[16:] {
			count++
		}
	}
	return count
}

func toBinary(i int64) string {
	return fmt.Sprintf("%.32b", i)
}
