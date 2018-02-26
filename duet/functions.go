package duet

import (
	"strconv"
	"strings"
	"time"

	"github.com/emman27/aoc2017/utils"
	"github.com/sirupsen/logrus"
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

type program struct {
	ID        int64
	Registers map[byte]int64
	Input     *utils.Queue
	Output    *utils.Queue
}

func newProgram(ID int64, input, output *utils.Queue) *program {
	return &program{
		ID:        ID,
		Registers: map[byte]int64{'p': ID},
		Input:     input,
		Output:    output,
	}
}

func (p *program) execute(filename string) int {
	count := 0
	commands := readInput(filename)
	idx := int64(0)
	logrus.Infof("Program %d is starting up", p.ID)
	for idx >= 0 && idx <= int64(len(commands)) {
		cmd := commands[idx]
		// logrus.Infof("Program %d is running %s %v %s", p.ID, cmd.Command, string([]byte{cmd.Register}), cmd.Value)
		// logrus.Infof("Program %d Input Queue: %v, Output Queue: %v", p.ID, p.Input, p.Output)
		switch cmd.Command {
		case "snd":
			val, err := strconv.Atoi(string(cmd.Register))
			iVal := int64(val)
			if err != nil {
				iVal = p.Registers[cmd.Register]
			}
			count++
			// logrus.Infof("Program %d Input Queue: %v, Output Queue: %v", p.ID, p.Input, p.Output)
			p.Output.Push(iVal)
		case "add":
			val, err := strconv.Atoi(cmd.Value)
			iVal := int64(val)
			if err != nil {
				iVal = p.Registers[cmd.Value[0]]
			}
			p.Registers[cmd.Register] += iVal
		case "set":
			val, err := strconv.Atoi(cmd.Value)
			iVal := int64(val)
			if err != nil {
				iVal = p.Registers[cmd.Value[0]]
			}
			p.Registers[cmd.Register] = iVal
		case "mul":
			val, err := strconv.Atoi(cmd.Value)
			iVal := int64(val)
			if err != nil {
				iVal = p.Registers[cmd.Value[0]]
			}
			p.Registers[cmd.Register] *= iVal
		case "mod":
			val, err := strconv.Atoi(cmd.Value)
			iVal := int64(val)
			if err != nil {
				iVal = p.Registers[cmd.Value[0]]
			}
			p.Registers[cmd.Register] %= iVal
		case "rcv":
			newVal, err := p.Input.Pop(1 * time.Second)
			if err != nil {
				return count
			}
			p.Registers[cmd.Register] = newVal.(int64)
		case "jgz":
			val, err := strconv.Atoi(string([]byte{cmd.Register}))
			var value int64
			if err != nil {
				value = p.Registers[cmd.Register]
			} else {
				value = int64(val)
			}
			if value > 0 {
				val, err := strconv.Atoi(cmd.Value)
				iVal := int64(val)
				if err != nil {
					iVal = p.Registers[cmd.Value[0]]
				}
				idx += iVal
				continue
			}
		}
		idx++
	}
	return count
}

// PartB counts the number of times program 1 sends a value
func PartB(filename string) int {
	out0 := &utils.Queue{Items: []interface{}{}}
	out1 := &utils.Queue{Items: []interface{}{}}
	result := make(chan int)
	program0 := newProgram(0, out1, out0)
	program1 := newProgram(1, out0, out1)
	go program0.execute(filename)
	go func() {
		result <- program1.execute(filename)
	}()
	return <-result
}
