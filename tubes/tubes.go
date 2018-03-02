package tubes

import (
	"errors"
	"strings"

	"github.com/emman27/aoc2017/utils"
)

type diagram struct {
	Points map[int]map[int]byte
}

func (d *diagram) Set(x, y int, value byte) {
	_, exists := d.Points[x]
	if !exists {
		d.Points[x] = map[int]byte{}
	}
	d.Points[x][y] = value
}

type packet struct {
	X         int
	Y         int
	Direction direction
	Trail     []byte
	Moves     int
}

type direction int

// Constants relating to the 4 possible preferred directions for a packet
const (
	UP    direction = 1
	RIGHT direction = 2
	DOWN  direction = 3
	LEFT  direction = 0
)

func (d direction) opposite() direction {
	return (d + 2) % 4
}

func (d diagram) isDeadEnd(x, y int) bool {
	_, existsDown := d.Points[x+1][y]
	_, existsUp := d.Points[x-1][y]
	row, rowExists := d.Points[x]
	if !rowExists {
		row = map[int]byte{}
	}
	_, existsLeft := row[y-1]
	_, existsRight := row[y+1]
	return utils.BoolToInt(existsDown)+utils.BoolToInt(existsUp)+utils.BoolToInt(existsLeft)+utils.BoolToInt(existsRight) <= 1
}

func (d diagram) At(x, y int) (byte, error) {
	row, exists := d.Points[x]
	if !exists {
		return byte(0), errors.New("no path here")
	}
	val, exists := row[y]
	if !exists {
		return byte(0), errors.New("no path here")
	}
	return val, nil
}

func move(x, y int, d direction) (int, int) {
	switch d {
	case UP:
		return x - 1, y
	case DOWN:
		return x + 1, y
	case LEFT:
		return x, y - 1
	case RIGHT:
		return x, y + 1
	}
	panic("no step available")
}

func (p packet) move(d diagram) packet {
	var step direction
	var newPacket packet
	switch p.Direction {
	case UP:
		if _, err := d.At(p.X-1, p.Y); err == nil {
			step = UP
		}
	case DOWN:
		if _, err := d.At(p.X+1, p.Y); err == nil {
			step = DOWN
		}
	case LEFT:
		if _, err := d.At(p.X, p.Y-1); err == nil {
			step = LEFT
		}
	case RIGHT:
		if _, err := d.At(p.X, p.Y+1); err == nil {
			step = RIGHT
		}
	}
	if step == 0 {
		directions := []direction{UP, DOWN, LEFT, RIGHT}
		for _, dir := range directions {
			x, y := move(p.X, p.Y, dir)
			if _, err := d.At(x, y); err == nil && dir != p.Direction.opposite() {
				step = dir
			}
		}
	}
	newPacket.X, newPacket.Y = move(p.X, p.Y, step)
	newPacket.Trail = p.Trail
	newPacket.Direction = step
	if char, _ := d.At(newPacket.X, newPacket.Y); !strings.Contains("-|+", string(char)) {
		newPacket.Trail = append(p.Trail, char)
	}
	newPacket.Moves = p.Moves + 1
	return newPacket
}

// PartA finds the order of letters taken
func PartA(filename string) string {
	lines := utils.ReadFile(filename)
	d := generateMap(lines)
	p := startPacket(d)
	p = p.move(d)
	for !d.isDeadEnd(p.X, p.Y) {
		p = p.move(d)
	}
	return string(p.Trail)
}

// PartB counts the length of the path of our dear packet
func PartB(filename string) int {
	lines := utils.ReadFile(filename)
	d := generateMap(lines)
	p := startPacket(d)
	p = p.move(d)
	for !d.isDeadEnd(p.X, p.Y) {
		p = p.move(d)
	}
	return p.Moves
}

func startPacket(d diagram) packet {
	p := packet{Direction: DOWN, Moves: 1}
	for y := range d.Points[0] {
		p.Y = y
	}
	return p
}

func generateMap(lines []string) diagram {
	d := diagram{Points: map[int]map[int]byte{}}
	for x := 0; x < len(lines); x++ {
		for y := 0; y < len(lines[x]); y++ {
			if lines[x][y] != ' ' {
				d.Set(x, y, lines[x][y])
			}
		}
	}
	return d
}
