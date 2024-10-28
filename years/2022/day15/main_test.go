package main

import (
	"github.com/MRSharff/algo/graph"
	"io/ioutil"
	"sort"
	"testing"
)

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

var shortestPathGoesUp = `19111
19191
11191
99991`

var shortestPathGoesLeft = `11111
99991
11111
19999
11111`

var bigTestInput = func() string {
	b, err := ioutil.ReadFile("test_input.txt")
	if err != nil {
		panic(err)
	}
	return string(b)
}()

func TestRiskGraph(t *testing.T) {
	small := "11" + "\n" +
		"21"
	g := buildRiskGraph(small)
	if len(g.nodes) != 4 {
		t.Errorf("bad len")
	}

	if g.getStart() != graph.Node(0) {
		t.Errorf("bad start")
	}

	if g.getEnd() != graph.Node(3) {
		t.Errorf("bad end")
	}

	if g.Weight(0, 2) != 2 {
		t.Errorf("bad weight")
	}

	expectedNeighbors := []graph.Node{1, 2}
	neighbors := g.Neighbors(0)
	if len(neighbors) != len(expectedNeighbors) {
		t.Errorf("bad neighbors")
	}

	sort.Slice(neighbors, func(i, j int) bool {
		return int(neighbors[i]) < int(neighbors[j])
	})

	for i := 0; i < len(neighbors); i++ {
		if neighbors[i] != expectedNeighbors[i] {
			t.Errorf("bad neighbors")
		}
	}
}

func TestPart1(t *testing.T) {
	smallPathLength := part1(smallInput)
	if smallPathLength != 4 {
		t.Error("Should work on a 3x3 cave")
	}

	got := part1(testInput)
	if got != 40 {
		t.Errorf("Expected 40, got %d\n", got)
	}

	if part1(shortestPathGoesUp) != 11 {
		t.Error("Should work on caves where the shortest path can go up")
	}

	leftMovingPathLength := part1(shortestPathGoesLeft)
	if leftMovingPathLength != 16 {
		t.Error("Should work on caves where the shortest path can go left")
	}
}

func TestRepeatRisk(t *testing.T) {
	small := "11" + "\n" +
		"21"
	g := buildRiskGraph(small)
	rg := repeatRisk{
		riskGraph: g,
		repeat:    2,
	}

	if rg.getStart() != graph.Node(0) {
		t.Errorf("bad start")
	}

	if rg.getEnd() != graph.Node(15) {
		t.Errorf("Bad end")
	}

	expectedNeighbors := []graph.Node{2, 5, 7, 10}
	neighbors := rg.Neighbors(6)
	sort.Slice(neighbors, func(i, j int) bool {
		return neighbors[i] < neighbors[j]
	})

	if len(neighbors) != len(expectedNeighbors) {
		t.Errorf("Different amount of neighbors than expected")
	}
	for i := 0; i < len(neighbors); i++ {
		if neighbors[i] != expectedNeighbors[i] {
			t.Errorf("Expected neighbor %d, got %d\n", expectedNeighbors[i], neighbors[i])
		}
	}

	got := graph.Dijkstras(rg, rg.getStart(), rg.getEnd())
	expected := 13
	if got != expected {
		t.Errorf("Expected %d, got %d\n", expected, got)
	}
}

func TestRepeatRisk_Neighbors(t *testing.T) {
	g := buildRiskGraph(testInput)
	rg := repeatRisk{
		riskGraph: g,
		repeat:    5,
	}

	// TODO: Come up with more pragmatic cases, like testing edge cases.

	expectedNeighbors := []graph.Node{8, 10, 59}
	neighbors := rg.Neighbors(9)
	sort.Slice(neighbors, func(i, j int) bool {
		return neighbors[i] < neighbors[j]
	})
	if len(expectedNeighbors) != len(neighbors) {
		t.Fatal("Neighbors Not the same size")
	}

	var risks []int
	for i := 0; i < len(neighbors); i++ {
		risks = append(risks, rg.Weight(0, neighbors[i]))
	}

	expectedRisk := []int{4, 2, 2}
	if len(expectedRisk) != len(expectedNeighbors) {
		t.Fatal("expectedRisk should be the same length as your expectedNeighbors")
	}

	for i := 0; i < len(expectedRisk); i++ {
		if expectedNeighbors[i] != neighbors[i] {
			t.Errorf("Expected %d, got %d\n", expectedNeighbors[i], neighbors[i])
		}
		if expectedRisk[i] != risks[i] {
			t.Errorf("Expected %d risk, got %d risk\n", expectedRisk[i], risks[i])
		}
	}

	node := graph.Node(510)
	risk := rg.Weight(0, node)
	if risk != 3 {
		t.Errorf("Expected 3 risk, got %d\n", risk)
	}
}

func TestPart2(t *testing.T) {

	totalRisk := part1(bigTestInput)
	expected := 315
	if totalRisk != expected {
		t.Errorf("Expected %d, got %d\n", expected, totalRisk)
	}

	p2risk := part2(testInput)
	if p2risk != expected {
		t.Errorf("Expected %d, got %d\n", expected, p2risk)
	}
}
