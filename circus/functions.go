package circus

import (
	"bufio"
	"errors"
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
				}
			}
		}
	}
	return programs
}

// AddParent adds a parent to the program
func (p *Program) AddParent(parent *Program) *Program {
	p.Parent = parent
	parent.Children = append(parent.Children, p)
	return p
}

// TotalWeight gives the weight of a program and its' stack
func (p *Program) TotalWeight() int {
	return p.Weight + p.childrenWeight()
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
	leaves := []*Program{}
	nextTier := []*Program{}
	for _, p := range programs {
		if len(p.Children) == 0 {
			leaves = append(leaves, p)
		}
	}
	for i := 1; len(leaves) != 0; i++ {
		logrus.Warnf("Starting tier %d", i)
		for _, l := range leaves {
			if l.Parent != nil {
				nextTier = append(nextTier, l.Parent)
			}
			w, err := l.expectedWeight()
			if err != nil {
				logrus.Errorf("Error! %v", err)
			} else if w != l.TotalWeight() {
				return w - l.childrenWeight()
			}
		}
		leaves = nextTier
		nextTier = []*Program{}
	}
	return 0
}

func (p *Program) childrenWeight() int {
	total := 0
	for _, child := range p.Children {
		total += child.TotalWeight()
	}
	return total
}

func (p *Program) expectedWeight() (int, error) {
	weights := map[int]int{} // count of weights to count
	if p.Parent == nil {
		return 0, errors.New("There's no parent for this node")
	}
	for _, child := range p.Parent.Children {
		weights[child.TotalWeight()]++
	}
	logrus.Info(weights)
	for weight, count := range weights {
		if count != 1 {
			return weight, nil
		}
	}
	return p.TotalWeight(), nil
}
