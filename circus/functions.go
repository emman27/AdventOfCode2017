package circus

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
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
	unbalanced := rootNode.findUnbalancedNode()
	expected, _ := unbalanced.expectedWeight()
	return expected - unbalanced.childrenWeight()
}

func (p *Program) findUnbalancedNode() *Program {
	if len(p.Children) == 0 {
		return nil
	}
	weights := map[int]int{} // Hash of weight to count
	indexes := map[int]int{} // Hash of weight to index in children
	for _, tower := range p.Children {
		weights[tower.TotalWeight()]++
	}
	var oddNode *Program
	for weight, count := range weights {
		if count == 1 && len(weights) != 1 {
			oddNode = p.Children[indexes[weight]]
		}
	}
	for _, node := range p.Children {
		if res := node.findUnbalancedNode(); res != nil {
			return res
		}
	}
	if oddNode != nil {
		return oddNode
	}
	return nil
}

func (p *Program) childrenWeight() int {
	total := 0
	for _, child := range p.Children {
		total += child.Weight
	}
	return total
}

func (p *Program) expectedWeight() (int, error) {
	weights := map[int]int{} // count of weights to count
	for _, child := range p.Parent.Children {
		weights[child.TotalWeight()]++
	}
	for weight, count := range weights {
		if count != 1 {
			return weight, nil
		}
	}
	return 0, errors.New("Cannot find expected parent")
}
