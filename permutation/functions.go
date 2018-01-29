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

func PartA(filename string) string {
	s := "abcdefghijklmnop"
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
