package permutation

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func spin(s string, i int) string {
	return s[len(s)-i:] + s[:len(s)-i]
}

func exchange(s string, a, b int) string {
	if a > b {
		tmp := a
		a = b
		b = tmp
	}
	return s[:a] + string(s[b]) + s[a+1:b] + string(s[a]) + s[b+1:]
}

func partner(s string, a, b rune) string {
	var i, j int
	for idx, c := range s {
		if c == a {
			i = idx
		} else if c == b {
			j = idx
		}
	}
	return exchange(s, i, j)
}

func splitBySlash(r rune) bool {
	return r == '/'
}

// PartA finds the results of spinning through a certain number of times
func PartA(filename, s string) string {
	cmds := strings.FieldsFunc(readFile(filename), func(r rune) bool { return r == ',' })
	for _, cmd := range cmds {
		switch cmd[0] {
		case 's':
			count, err := strconv.Atoi(cmd[1:])
			if err != nil {
				panic(err)
			}
			s = spin(s, count)
		case 'x':
			fields := strings.FieldsFunc(cmd[1:], splitBySlash)
			a, err := strconv.Atoi(fields[0])
			b, err2 := strconv.Atoi(fields[1])
			if err != nil || err2 != nil {
				panic(err2)
			}
			s = exchange(s, a, b)
		case 'p':
			s = partner(s, rune(cmd[1]), rune(cmd[3]))
		}
	}
	return s
}

func readFile(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines[0]
}

func effectivePartA(s string) string {
	return string([]byte{s[4], s[1], s[9], s[15], s[5], s[3], s[6], s[12], s[8], s[7], s[14], s[13], s[0], s[2], s[10], s[11]})
}

// PartB finds the results of performing PartA one billion times
func PartB(filename string) string {
	found := map[string]int{}
	s := "abcdefghijklmnop"
	for i := 0; i < 1000000000; i++ {
		s = effectivePartA(s)
		val, exists := found[s]
		if exists {
			i = fastForward(val, i)
		}
		found[s] = i
	}
	return s
}

func fastForward(originalIndex, currIndex int) int {
	diff := originalIndex - currIndex
	remainder := (1000000000 - currIndex) % diff
	return 1000000000 - remainder
}
