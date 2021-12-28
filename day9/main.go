package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

var testInput = `2199943210
3987894921
9856789892
8767896789
9899965678`

type heightmap [][]int

func createHeightmap(in string) heightmap {
	lines := strings.Split(in, "\n")
	rows := len(lines)
	cols := len(lines[0])
	h := make(heightmap, rows)
	for y := 0; y < rows; y++ {
		h[y] = make([]int, cols)
		for x := 0; x < cols; x++ {
			var err error
			h[y][x], err = strconv.Atoi(string(lines[y][x]))
			if err != nil {
				panic(err)
			}
		}
	}
	return h
}

type point struct {
	x, y int
}

func left(h heightmap, x, y int) int {
	if x == 0 {
		return math.MaxInt64
	}
	return h[y][x-1]
}

func right(h heightmap, x, y int) int {
	if x == len(h[y])-1 {
		return math.MaxInt64
	}
	return h[y][x+1]
}

func top(h heightmap, x, y int) int {
	if y == 0 {
		return math.MaxInt64
	}
	return h[y-1][x]
}

func bottom(h heightmap, x, y int) int {
	if y == len(h)-1 {
		return math.MaxInt64
	}
	return h[y+1][x]
}

func lowPoints(h heightmap) []point {
	var lowpoints []point
	for y := 0; y < len(h); y++ {
		for x := 0; x < len(h[y]); x++ {
			if isLowpoint(h, x, y) {
				lowpoints = append(lowpoints, point{x, y})
			}
		}
	}
	return lowpoints
}

func isLowpoint(h heightmap, x int, y int) bool {
	n := h[y][x]
	return n < left(h, x, y) && n < bottom(h, x, y) && n < right(h, x, y) && n < top(h, x, y)
}

func riskLevels(h heightmap, points []point) []int {
	sums := make([]int, len(points))
	for i, p := range points {
		sums[i] = 1 + h[p.y][p.x]
	}
	return sums
}

func sum(nums ...int) int {
	s := 0
	for _, n := range nums {
		s += n
	}
	return s
}

func main() {
	part1 := func(in string) {
		h := createHeightmap(in)
		fmt.Println(sum(riskLevels(h, lowPoints(h))...))
	}
	part1(testInput)
	b, _ := ioutil.ReadFile("day9/input.txt")
	part1(string(b))
}
