package main

import (
	"fmt"
	"github.com/MRSharff/algo/graph"
	"io/ioutil"
	"strconv"
	"strings"
)

var in = func() string {
	b, _ := ioutil.ReadFile("day15/input.txt")
	return string(b)
}()

func main() {
	// We use Dijkstras to find the shortest path from top left, to bottom right.
	// Each risk level is the edge weight.
	fmt.Println(part1(in))
	fmt.Println(part2(in))
}

type riskGraph struct {
	// nodes is a list of risks, we reference them by their index in this slice.
	nodes []int

	// the cave is technically an nxn matrix, m = height = width
	width, height int
}

// I really want to pull all this graph stuff out since it's required for AoC so often.
// I'm at least starting with Dijkstras in github.com/MRSharff/algo
func findNeighbors(index, width, height int) []int {
	row := index / width
	col := index % width
	hasTop, hasRight, hasBottom, hasLeft := row != 0, col != width-1, row != height-1, col != 0

	var neighbors []int
	if hasTop {
		top := index - width
		neighbors = append(neighbors, top)
	}

	if hasRight {
		right := index + 1
		neighbors = append(neighbors, right)
	}

	if hasBottom {
		bottom := index + width
		neighbors = append(neighbors, bottom)
	}

	if hasLeft {
		left := index - 1
		neighbors = append(neighbors, left)
	}

	return neighbors
}

func (r riskGraph) Neighbors(f graph.Node) []graph.Node {
	var neighbors []graph.Node
	for _, n := range findNeighbors(int(f), r.width, r.height) {
		neighbors = append(neighbors, graph.Node(n))
	}
	return neighbors
}

func (r riskGraph) Weight(f graph.Node, neighbor graph.Node) int {
	return r.nodes[neighbor]
}

func (r riskGraph) getStart() graph.Node {
	return graph.Node(0)
}

func (r riskGraph) getEnd() graph.Node {
	return graph.Node((r.width * r.height) - 1)
}

type repeatRisk struct {
	riskGraph
	repeat int // This is 5 in the AoC prompt, but I want to test with some smaller inputs.
}

func (r repeatRisk) Weight(f graph.Node, neighbor graph.Node) int {
	// This would be much easier with a 2d array instead.
	n := int(neighbor)
	col := n % (r.width * r.repeat)
	row := n / (r.height * r.repeat)
	smallCol := col % r.width
	smallRow := row % r.height
	smallNode := r.width*smallRow + smallCol

	rowMod := row / r.height
	colMod := col / r.width
	weight := r.riskGraph.Weight(f, graph.Node(smallNode)) + rowMod + colMod
	if weight > 9 {
		weight -= 9
	}
	return weight
}

func (r repeatRisk) Neighbors(f graph.Node) []graph.Node {
	var neighbors []graph.Node
	for _, n := range findNeighbors(int(f), r.width*r.repeat, r.height*r.repeat) {
		neighbors = append(neighbors, graph.Node(n))
	}
	return neighbors
}

func (r repeatRisk) getEnd() graph.Node {
	width := r.width * r.repeat
	height := r.height * r.repeat
	return graph.Node(width*height - 1)
}

func part2(input string) int {
	g := buildRiskGraph(input)
	rg := repeatRisk{
		riskGraph: g,
		repeat:    5,
	}
	start := rg.getStart()
	end := rg.getEnd()

	return graph.Dijkstras(rg, start, end)
}

func part1(in string) int {
	rg := buildRiskGraph(in)
	start, end := rg.getStart(), rg.getEnd()
	return graph.Dijkstras(rg, start, end)
}

func buildRiskGraph(in string) riskGraph {
	var nodes []int
	lines := strings.Split(in, "\n")
	for _, line := range lines {
		for _, c := range line {
			riskLevel, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}
			nodes = append(nodes, riskLevel)
		}
	}
	height, width := len(lines), len(lines[0])

	rg := riskGraph{
		nodes:  nodes,
		width:  width,
		height: height,
	}
	return rg
}
