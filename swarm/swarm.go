// Package swarm solves Day 20 of Advent of Code 2017
package swarm

import (
	"errors"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/emman27/aoc2017/utils"
)

// Geometry represents a 3-dimensional value
type Geometry struct {
	X int
	Y int
	Z int
}

// Particle represents a particle
type Particle struct {
	ID           int
	Position     Geometry
	Velocity     Geometry
	Acceleration Geometry
}

// parseCommand turns meaningless text into a particle
func parseCommand(s string, id int) Particle {
	p := Particle{ID: id}
	r, err := regexp.Compile(`<(\-|\ )?[0-9]+,(\-|\ )?[0-9]+,(\-|\ )?[0-9]+>`)
	if err != nil {
		panic(err)
	}
	matches := r.FindAllString(s, -1)
	p.Position = parseGeometry(matches[0])
	p.Velocity = parseGeometry(matches[1])
	p.Acceleration = parseGeometry(matches[2])
	return p
}

func parseGeometry(s string) Geometry {
	values := utils.SplitByComma(s[1 : len(s)-1])
	x, err := strconv.Atoi(strings.Trim(values[0], " "))
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(strings.Trim(values[1], " "))
	if err != nil {
		panic(err)
	}
	z, err := strconv.Atoi(strings.Trim(values[2], " "))
	if err != nil {
		panic(err)
	}
	return Geometry{x, y, z}
}

func manhattenDistance(g Geometry) int {
	return utils.AbsInt(g.X) + utils.AbsInt(g.Y) + utils.AbsInt(g.Z)
}

// PartA finds the nearest particle in the longrun based on Manhatten distance
func PartA(filename string) int {
	lines := utils.ReadFile(filename)
	particles := []Particle{}
	for idx, l := range lines {
		particles = append(particles, parseCommand(l, idx))
	}
	id := -1
	min := 2147483647 // lazy max value of int
	for _, p := range particles {
		if d := manhattenDistance(p.Acceleration); d < min {
			id = p.ID
			min = d
		}
	}
	return id
}

// PartB finds the number of uncollided particles
func PartB(filename string) int {
	lines := utils.ReadFile(filename)
	particles := []Particle{}
	for idx, l := range lines {
		particles = append(particles, parseCommand(l, idx))
	}
	return len(particles)
}

func collisionPoint(p1, p2 Particle) (bool, int) {
	xc := float64(p1.Position.X - p2.Position.X)
	xb := float64(p1.Velocity.X - p2.Velocity.X)
	xa := float64(p1.Acceleration.X - p2.Acceleration.X)

	return false, 0
}

func solveQuadraticOnlyPositiveIntegerRoots(a, b, c float64) (float64, error) {
	if math.Pow(b, 2)-4*a*c < 0 {
		return 0, errors.New("no root found")
	}
	rootLower := (-b - math.Sqrt(math.Pow(b, 2)-4*a*c)) / (2 * a)
	rootHigher := (-b + math.Sqrt(math.Pow(b, 2)-4*a*c)) / (2 * a)
	higher := math.Max(rootLower, rootHigher)
	return higher, nil
}
