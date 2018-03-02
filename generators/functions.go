package generators

type generator int64

var factorA generator = 16807
var factorB generator = 48271

type generatorB struct {
	Factor  generator
	Divisor int64
}

var genA = generatorB{factorA, 4}
var genB = generatorB{factorB, 8}

func next(g generator, i int64) int64 {
	return int64(g) * i % 2147483647
}

func nextB(g generatorB, i int64) int64 {
	start := next(g.Factor, i)
	for start%g.Divisor != 0 {
		start = next(g.Factor, start)
	}
	return start
}

// PartA calculates the collisiosn
func PartA(a, b int64) int {
	count := 0
	for i := 0; i < 40000000; i++ {
		a = next(factorA, a)
		b = next(factorB, b)
		if a&65535 == b&65535 {
			count++
		}
	}
	return count
}

// PartB does the same but with extra conditions
func PartB(a, b int64) int {
	count := 0
	for i := 0; i < 5000000; i++ {
		a = nextB(genA, a)
		b = nextB(genB, b)
		if a&65535 == b&65535 {
			count++
		}
	}
	return count
}

// func toBinary(i int64) string {
// 	return fmt.Sprintf("%.32b", i)
// }
