package spiralMemory

import (
	"math"
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
