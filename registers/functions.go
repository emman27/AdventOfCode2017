package registers

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

type registry map[string]int

func readData(filename string) []string {
	cmds := []string{}
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cmds = append(cmds, scanner.Text())
	}
	return cmds
}

func (r *registry) parseInput(s string) {
	cmdList := strings.Fields(s)
	src := cmdList[4]
	cond := cmdList[5]
	target := cmdList[0]
	action := cmdList[1]
	amt, err := strconv.Atoi(cmdList[2])
	if err != nil {
		panic(err)
	}
	val, err := strconv.Atoi(cmdList[6])
	if err != nil {
		panic(err)
	}
	if testCond(r, src, cond, val) {
		trigger(r, target, action, amt)
	}
}

func testCond(r *registry, src, cond string, val int) bool {
	reg := *r
	switch cond {
	case ">":
		return reg[src] > val
	case "<":
		return reg[src] < val
	case ">=":
		return reg[src] >= val
	case "<=":
		return reg[src] <= val
	case "==":
		return reg[src] == val
	case "!=":
		return reg[src] != val
	}
	return false
}

func trigger(r *registry, target, action string, val int) {
	reg := *r
	switch action {
	case "inc":
		reg[target] += val
	case "dec":
		reg[target] -= val
	}
}

// PartA finds the largest value in any register after all commands are done
func PartA(filename string) int {
	reg := registry{}
	cmds := readData(filename)
	for _, cmd := range cmds {
		reg.parseInput(cmd)
	}
	max := math.Inf(-1)
	for _, value := range reg {
		max = math.Max(max, float64(value))
	}
	return int(max)
}
