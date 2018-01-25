// Package pipes solves day 12 of AOC2017
package pipes

import (
	"bufio"
	"os"
	"strings"
)

type graph struct {
	Nodes []node
	Edges []edge
}

type node struct {
	Name string
}

type edge struct {
	From *node
	To   *node
}

func readInput(filename string) *graph {
	g := &graph{}
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parseLine(g, scanner.Text())
	}
	return g
}

func parseLine(g *graph, line string) *graph {
	items := strings.Fields(line)
	n := g.getNode(items[0])
	for _, j := range items[2:] {
		partner := g.getNode(strings.Trim(j, ", "))
		g.addEdge(n, partner)
	}
	return g
}

func (g *graph) getNode(name string) node {
	for _, n := range g.Nodes {
		if n.Name == name {
			return n
		}
	}
	n := node{Name: name}
	g.Nodes = append(g.Nodes, n)
	return n
}

func (g *graph) addEdge(src, dest node) {
	e := edge{From: &src, To: &dest}
	e2 := edge{From: &dest, To: &src}
	for _, edge := range g.Edges {
		if edge.From.Name == e.From.Name && edge.To.Name == e.To.Name {
			return
		}
	}
	g.Edges = append(g.Edges, e)
	g.Edges = append(g.Edges, e2)
}

// PartA returns the number of nodes connected to the node named 0
func PartA(filename string) int {
	g := readInput(filename)
	return g.countNodes("0")
}

func (g *graph) countNodes(name string) int {
	found := []node{}
	connected := []node{}
	for _, n := range g.Nodes {
		if n.Name == name {
			connected = append(connected, n)
		}
	}
	for len(connected) != 0 {
		toAdd := connected[0]
		connected = connected[1:]
		skip := false
		for _, f := range found {
			if f.Name == toAdd.Name {
				skip = true
			}
		}
		if !skip {
			found = append(found, toAdd)
			for _, e := range g.Edges {
				if e.From.Name == toAdd.Name {
					connected = append(connected, *e.To)
				}
			}
		}
		connected = unique(connected)
	}
	return len(found)
}

func (n node) String() string {
	return n.Name
}

func unique(nodes []node) []node {
	res := []node{}
	for _, n := range nodes {
		found := false
		for _, e := range res {
			if e.Name == n.Name {
				found = true
			}
		}
		if !found {
			res = append(res, n)
		}
	}
	return res
}
