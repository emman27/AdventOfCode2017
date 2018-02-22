package duet

import (
	"strconv"
	"strings"

	"github.com/emman27/aoc2017/utils"
)

// PartA finds the value of the recovered frequency the first time a rcv instruction is executed
// with a nonzero value
func PartA(filename string) int {
	commands := readInput(filename)
	registers := map[byte]int{}
	var lastPlayed int
	idx := 0
	for idx >= 0 && idx <= len(commands) {
		cmd := commands[idx]
		switch cmd.Command {
		case "snd":
			lastPlayed = registers[cmd.Register]
		case "add":
			val, err := strconv.Atoi(cmd.Value)
			if err != nil {
				panic(err)
			}
			registers[cmd.Register] += val
		case "set":
			val, err := strconv.Atoi(cmd.Value)
			if err != nil {
				val = registers[cmd.Value[0]]
			}
			registers[cmd.Register] = val
		case "mul":
			val, err := strconv.Atoi(cmd.Value)
			if err != nil {
				val = registers[cmd.Value[0]]
			}
			registers[cmd.Register] *= val
		case "mod":
			val, err := strconv.Atoi(cmd.Value)
			if err != nil {
				val = registers[cmd.Value[0]]
			}
			registers[cmd.Register] %= val
		case "rcv":
			if registers[cmd.Register] != 0 {
				return lastPlayed
			}
		case "jgz":
			if registers[cmd.Register] > 0 {
				val, err := strconv.Atoi(cmd.Value)
				if err != nil {
					panic(err)
				}
				idx += val
			} else {
				idx++
			}
		}
		if cmd.Command != "jgz" {
			idx++
		}
	}
	return lastPlayed
}

type command struct {
	Command  string
	Register byte
	Value    string
}

func readInput(filename string) []command {
	commands := utils.ReadFile(filename)
	results := []command{}
	for _, cmd := range commands {
		stuff := strings.Fields(cmd)
		var value string
		if len(stuff) == 3 {
			value = stuff[2]
		}
		results = append(results, command{Command: stuff[0], Register: stuff[1][0], Value: value})
	}
	return results
}
