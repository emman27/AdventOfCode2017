package scanners

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// PartA calculates the severity score
func PartA(filename string) int {
	cameras := readData(filename)
	return calculateScore(cameras, 0)
}

func calculateScore(cameras map[int]int, startTime int) int {
	score := 0
	for depth, rang := range cameras {
		if positionAtTime(depth+startTime, rang) == 0 {
			score += depth * rang
		}
	}
	return score
}

func checkCaught(cameras map[int]int, startTime int) bool {
	for depth, rang := range cameras {
		if positionAtTime(depth+startTime, rang) == 0 {
			return true
		}
	}
	return false
}

// PartB calculates the minimal delay before entering the system
func PartB(filename string) int {
	cameras := readData(filename)
	delay := 0
	for {
		if !checkCaught(cameras, delay) {
			return delay
		}
		delay++
	}
	return delay
}

func positionAtTime(t, rang int) int {
	return t % (2 * (rang - 1))
}

func readData(filename string) map[int]int {
	cameras := map[int]int{}
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cameras = parse(cameras, scanner.Text())
	}
	return cameras
}

func parse(cameras map[int]int, line string) map[int]int {
	fields := strings.FieldsFunc(line, func(r rune) bool { return r == ':' })
	depth, err := strconv.Atoi(fields[0])
	if err != nil {
		panic(err)
	}
	rang, err := strconv.Atoi(strings.Trim(fields[1], " "))
	if err != nil {
		panic(err)
	}
	cameras[depth] = rang
	return cameras
}
