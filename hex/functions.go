package hex

import (
	"math"
	"strings"
)

type coord struct {
	X int
	Y int
}

// PartA returns the number of steps to the child
func PartA(dir string) int {
	steps := 0
	pt := parseTravel(dir)
	for pt.X != 0 || pt.Y != 0 {
		pt = travel(pt)
		steps++
	}
	return steps
}

// PartB returns the max distance at any point in the path
func PartB(dir string) int {
	return parseTravelB(dir)
}

func travel(pt coord) coord {
	if pt.X > 0 && pt.Y > 0 {
		pt.X--
		pt.Y--
	} else if pt.X < 0 && pt.Y < 0 {
		pt.X++
		pt.Y++
	} else if pt.X > 0 && pt.Y < 0 {
		pt.X--
		pt.Y++
	} else if pt.X < 0 && pt.Y > 0 {
		pt.X++
		pt.Y--
	} else if pt.X == 0 && pt.Y > 0 {
		pt.Y -= 2
	} else if pt.X == 0 && pt.Y < 0 {
		pt.Y += 2
	}
	return pt
}

func parseTravel(dir string) coord {
	pt := coord{}
	fields := strings.FieldsFunc(dir, func(c rune) bool { return c == ',' })
	for _, c := range fields {
		switch c {
		case "ne":
			pt.X++
			pt.Y++
		case "nw":
			pt.X--
			pt.Y++
		case "se":
			pt.Y--
			pt.X++
		case "sw":
			pt.Y--
			pt.X--
		case "n":
			pt.Y += 2
		case "s":
			pt.Y -= 2
		}
	}
	return pt
}

func parseTravelB(dir string) int {
	maxDist := 0
	pt := coord{}
	fields := strings.FieldsFunc(dir, func(c rune) bool { return c == ',' })
	for _, c := range fields {
		switch c {
		case "ne":
			pt.X++
			pt.Y++
		case "nw":
			pt.X--
			pt.Y++
		case "se":
			pt.Y--
			pt.X++
		case "sw":
			pt.Y--
			pt.X--
		case "n":
			pt.Y += 2
		case "s":
			pt.Y -= 2
		}
		if pt.dist() > maxDist {
			maxDist = pt.dist()
		}
	}
	return maxDist
}

func (c coord) dist() int {
	if math.Abs(float64(c.Y)) > math.Abs(float64(c.X)) {
		return int(math.Abs(float64(c.X)) + math.Abs(math.Abs(float64(c.Y))-math.Abs(float64(c.X)))/2)
	}
	return int(math.Abs(float64(c.X)))
}
