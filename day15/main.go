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

var testInput = `1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`

var smallInput = `111
441
441`

var pathShouldGoUpInput = `19111
19191
11191
99991`

func main() {
	// * motion restricted to 2 dimensions
	// * cavern is a square (or at least resembles one)
	// * The starting position is never entered, so its risk is not counted.

	// I think this is a search problem with edge weights being the risk levels
	// Dijkstras
	fmt.Println(part1(smallInput))

	if part1(testInput) != 40 {
		panic("Test failed")
	}

	fmt.Println(part1(pathShouldGoUpInput))

	fmt.Println(part1(in))

	//totalRisk := part2(testInput)
	//if totalRisk != 315 {
	//	panic("Part 2 test failed. Expected 315, got, " + strconv.Itoa(totalRisk))
	//} else {
	//	fmt.Println("Part 2 test passed!")
	//}
}

type riskGraph struct {
	nodes         []int
	width, height int
}

func (r riskGraph) Neighbors(f graph.Node) []graph.Node {
	n := int(f)
	height := r.height
	width := r.width
	row := n / width
	col := n % width
	// can we skip left nodes, also perhaps top nodes?
	hasTop, hasRight, hasBottom := row != 0, col != width-1, row != height-1
	_ = hasTop

	var neighbors []graph.Node
	if hasTop {
		top := n - width
		neighbors = append(neighbors, graph.Node(top))
	}

	if hasRight {
		right := f + 1
		neighbors = append(neighbors, right)
	}

	if hasBottom {
		bottom := n + width
		neighbors = append(neighbors, graph.Node(bottom))
	}

	return neighbors
}

func (r riskGraph) Weight(f graph.Node, neighbor graph.Node) int {
	return r.nodes[neighbor]
}

func part2(input string) int {

	return 0
}

func part1(in string) int {
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
	start := graph.Node(0)
	end := graph.Node((width * height) - 1)

	rg := riskGraph{
		nodes:  nodes,
		width:  width,
		height: height,
	}
	return graph.Dijkstras(rg, start, end)
}
