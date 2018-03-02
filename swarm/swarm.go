// Package swarm solves Day 20 of Advent of Code 2017
package swarm

import (
	"regexp"
	"strconv"

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

// Tick adjusts the velocity and position of a particle
func (p *Particle) Tick() *Particle {
	p.Velocity = Geometry{
		X: p.Velocity.X + p.Acceleration.X,
		Y: p.Velocity.Y + p.Acceleration.Y,
		Z: p.Velocity.Z + p.Acceleration.Z,
	}
	p.Position = Geometry{
		X: p.Position.X + p.Velocity.X,
		Y: p.Position.Y + p.Velocity.Y,
		Z: p.Position.Z + p.Velocity.Z,
	}
	return p
}

// parseCommand turns meaningless text into a particle
func parseCommand(s string, id int) Particle {
	p := Particle{ID: id}
	r, err := regexp.Compile(`<\-?[0-9]+,\-?[0-9]+,\-?[0-9]+>`)
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
	x, err := strconv.Atoi(values[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(values[1])
	if err != nil {
		panic(err)
	}
	z, err := strconv.Atoi(values[2])
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
