package circus

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

// Program is a program in the system
type Program struct {
	Name     string
	Weight   int
	Parent   *Program
	Children []*Program
}

// ReadData reads data from a txt file and returns an array of Program
func ReadData(filename string) []*Program {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	programs := []*Program{}
	for scanner.Scan() {
		lst := strings.Fields(scanner.Text())
		weight, err := strconv.Atoi(lst[1][1 : len(lst[1])-1])
		if err != nil {
			panic(err)
		}
		programs = append(programs, &Program{
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
			for _, prog := range programs {
				if prog.Name == childName {
					prog.AddParent(programs[idx])
					programs[idx].AddChild(prog)
				}
			}
		}
	}
	return programs
}

// AddParent adds a parent to the program
func (p *Program) AddParent(parent *Program) *Program {
	p.Parent = parent
	return p
}

// AddChild adds a child program
func (p *Program) AddChild(child *Program) *Program {
	p.Children = append(p.Children, child)
	return p
}

// TotalWeight gives the weight of a program and its' stack
func (p *Program) TotalWeight() int {
	total := 0
	for _, c := range p.Children {
		total += c.TotalWeight()
	}
	return p.Weight + total
}

func (p *Program) String() string {
	return p.Name
}

// PartA finds the parent program
func PartA(filename string) string {
	programs := ReadData(filename)
	for _, prog := range programs {
		if prog.Parent == nil {
			return prog.Name
		}
	}
	return ""
}

// PartB finds the oddly weighted thingy
func PartB(filename string) int {
	programs := ReadData(filename)
	rootNodeName := PartA(filename)
	var rootNode *Program
	for _, p := range programs {
		if rootNodeName == p.Name {
			rootNode = p
		}
	}
	if balanced, val := rootNode.IsBalanced(); !balanced {
		return val
	}
	return 0
}

// IsBalanced checks if a tree is balanced
func (p *Program) IsBalanced() (bool, int) {
	weights := make(map[int]int)
	for _, c := range p.Children {
		weights[c.TotalWeight()] = weights[c.TotalWeight()] + 1
	}
	logrus.Info(weights)
	for key, value := range weights {
		if value != 1 && len(weights) != 1 {
			return false, key
		}
	}
	return len(weights) == 1, 0
}
