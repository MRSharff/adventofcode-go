package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var testInput = `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`

var smallTest = `11111
19991
19191
19991
11111`

type octopus struct {
	i, j int
}

var octopuses [][]int
var flashed [][]bool
var flashCount int
var energizer chan int

// Probably a little over-engineered, but I've been wanting to try
// using channels more since they work decently as a queue for BFS
func main() {
	part1(testInput)
	part1(smallTest)

	b, _ := ioutil.ReadFile("day11/input.txt")
	part1(string(b))
}

func part1(in string) {
	setOctopuses(in)
	fmt.Println("Before any steps:")
	printOctopuses()
	for i := 0; i < 100; i++ {
		increaseEnergyLevels()
		flash()
		resetFlashedEnergyLevels()
		fmt.Printf("After step %d:\n", i+1)
		printOctopuses()
	}
	fmt.Println("Total Flashes: ", flashCount)
}

func printOctopuses() {
	for i := 0; i < len(octopuses); i++ {
		for j := 0; j < len(octopuses[i]); j++ {
			fmt.Print(octopuses[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func setOctopuses(in string) {
	octopuses = make([][]int, 0)
	flashed = make([][]bool, 0)
	flashCount = 0
	energizer = make(chan int, 500)
	for i, line := range strings.Split(in, "\n") {
		octopuses = append(octopuses, []int{})
		flashed = append(flashed, []bool{})
		for _, c := range line {
			energy, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}
			octopuses[i] = append(octopuses[i], energy)
			flashed[i] = append(flashed[i], false)
		}
	}

	// make sure the neighborCache is reset when we move to a different
	// sized input, like from testInput to smallTest.
	if len(octopuses)*len(octopuses[0]) != len(neighborCache) {
		neighborCache = make(map[octopus][]octopus)
	}
}

func increaseEnergyLevels() {
	for i := 0; i < len(octopuses); i++ {
		for j := 0; j < len(octopuses[i]); j++ {
			energizer <- i
			energizer <- j
		}
	}
}

func resetFlashedEnergyLevels() {
	for i := 0; i < len(octopuses); i++ {
		for j := 0; j < len(octopuses[i]); j++ {
			if flashed[i][j] {
				octopuses[i][j] = 0
				flashed[i][j] = false
			}
		}
	}
	energizer = make(chan int, 500)
}

func flash() {
	doneFlashing := false
	for !doneFlashing {
		select {
		case i := <-energizer:
			j := <-energizer
			octopuses[i][j]++
			energy := octopuses[i][j]
			if flashes := energy > 9 && !flashed[i][j]; flashes {
				flashOctopus(i, j)
			}
		default:
			close(energizer)
			doneFlashing = true
		}
	}
}

func flashOctopus(i, j int) {
	flashed[i][j] = true
	flashCount++
	energizeNeighbors(i, j)
}

func energizeNeighbors(i, j int) {
	for _, n := range neighbors(i, j) {
		if flashed[n.i][n.j] {
			continue
		}
		energizer <- n.i
		energizer <- n.j
	}
}

var neighborCache map[octopus][]octopus

func neighbors(i int, j int) []octopus {
	if ns, isCached := neighborCache[octopus{i, j}]; isCached {
		return ns
	}
	var ns []octopus
	top, bottom, left, right := i-1, i+1, j-1, j+1
	hasTop := top >= 0
	hasRight := right < len(octopuses[i])
	hasBottom := bottom < len(octopuses)
	hasLeft := left >= 0
	if hasTop {
		ns = append(ns, octopus{top, j})
	}
	if hasTop && hasRight {
		ns = append(ns, octopus{top, right})
	}
	if hasRight {
		ns = append(ns, octopus{i, right})
	}
	if hasBottom && hasRight {
		ns = append(ns, octopus{bottom, right})
	}
	if hasBottom {
		ns = append(ns, octopus{bottom, j})
	}
	if hasBottom && hasLeft {
		ns = append(ns, octopus{bottom, left})
	}
	if hasLeft {
		ns = append(ns, octopus{i, left})
	}
	if hasTop && hasLeft {
		ns = append(ns, octopus{top, left})
	}
	neighborCache[octopus{i, j}] = ns
	return ns
}
