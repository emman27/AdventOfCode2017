package circus

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Program is a program in the system
type Program struct {
	Name    string
	Weight  int
	Parents []*Program
}

// ReadData reads data from a txt file and returns an array of Program
func ReadData(filename string) []Program {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	programs := []Program{}
	for scanner.Scan() {
		lst := strings.Fields(scanner.Text())
		weight, err := strconv.Atoi(lst[1][1 : len(lst[1])-1])
		if err != nil {
			panic(err)
		}
		programs = append(programs, Program{
			Name:   lst[0],
			Weight: weight,
		})
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	file.Seek(0, 0)
	scanner = bufio.NewScanner(file)
	for idx := 0; scanner.Scan(); idx++ {
		lst := strings.Fields(scanner.Text())
		for i := 3; i < len(lst); i++ {
			childName := lst[i]
			childName = strings.Trim(childName, ",")
			for childIdx, prog := range programs {
				if prog.Name == childName {
					programs[childIdx] = *prog.AddParent(&programs[idx])
				}
			}
		}
	}

	return programs
}

// AddParent adds a parent to the program
func (p *Program) AddParent(parent *Program) *Program {
	p.Parents = append(p.Parents, parent)
	return p
}

func (p *Program) String() string {
	return p.Name
}

// PartA finds the parent program
func PartA(filename string) string {
	programs := ReadData(filename)
	for _, prog := range programs {
		if len(prog.Parents) == 0 {
			return prog.Name
		}
	}
	return ""
}
