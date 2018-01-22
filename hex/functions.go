package hex

import (
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
