package spiralMemory

import (
	"math"

	"github.com/sirupsen/logrus"
)

// PartA calculates the Manhatten Distance from the port to center
func PartA(port int) int {
	ring := ringNumber(port)
	return ring + distToPillar(port, ring)
}

func ringNumber(port int) int {
	return int(math.Ceil(((math.Sqrt(float64(port))) - 1) / 2))
}

func distToPillar(port, ringNumber int) int {
	fport := float64(port)
	bottomRightCorner := (ringNumber*2 + 1) * (ringNumber*2 + 1)
	bottom := float64(bottomRightCorner - ringNumber)
	left := float64(bottomRightCorner - 3*ringNumber)
	top := float64(bottomRightCorner - 5*ringNumber)
	right := float64(bottomRightCorner - 7*ringNumber)
	return int(math.Min(
		math.Min(math.Abs(fport-top), math.Abs(fport-bottom)),
		math.Min(math.Abs(fport-left), math.Abs(fport-right)),
	))
}

func position(port int) (int, int) {
	ring := ringNumber(port)
	bottomRightCornerNumber := (ring*2 + 1) * (ring*2 + 1)
	x := ring
	y := -ring
	for i := bottomRightCornerNumber; i > port; i-- {
		if x > -ring && y == -ring {
			x--
		} else if x == -ring && y < ring {
			y++
		} else if y == ring && x < ring {
			x++
		} else if x == ring && y > -ring {
			y--
		} else {
			panic(ring)
		}
	}
	return x, y
}

// PartB does the calculation for part B
func PartB(min int) int {
	storage := NewStorage()
	storage.Set(0, 0, 1)
	curr := 1
	for i := 2; curr <= min; i++ {
		x, y := position(i)
		curr = storage.Calculate(x, y)
		logrus.Info(x, y, curr)
		storage.Set(x, y, curr)
	}
	return curr
}
