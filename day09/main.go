package main

import (
	"fmt"
	"io/ioutil"
	"sort"
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
	height := h[y][x]
	for _, p := range neighbors(h, point{x, y}) {
		if height >= h[p.y][p.x] {
			return false
		}
	}
	return true
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

func product(nums ...int) int {
	p := 1
	for _, n := range nums {
		p *= n
	}
	return p
}

type basin []point

func sizes(basins []basin) []int {
	s := make([]int, len(basins))
	for i, b := range basins {
		s[i] = len(b)
	}
	return s
}

func threeLargest(basins []basin) []basin {
	sort.Slice(basins, func(i, j int) bool {
		return len(basins[i]) > len(basins[j])
	})
	return []basin{
		basins[0],
		basins[1],
		basins[2],
	}
}

func neighbors(h heightmap, p point) []point {
	var points []point
	x, y := p.x, p.y
	hasTopNeighbor, hasBottomNeighbor, hasLeftNeighbor, hasRightNeighbor := y != 0, y != len(h)-1, x != 0, x != len(h[y])-1
	if hasTopNeighbor {
		topNeighbor := point{x, y - 1}
		points = append(points, topNeighbor)
	}
	if hasBottomNeighbor {
		bottomNeighbor := point{x, y + 1}
		points = append(points, bottomNeighbor)
	}
	if hasLeftNeighbor {
		leftNeighbor := point{x - 1, y}
		points = append(points, leftNeighbor)
	}
	if hasRightNeighbor {
		rightNeighbor := point{x + 1, y}
		points = append(points, rightNeighbor)
	}
	return points
}

func basinsOf(h heightmap) []basin {
	lps := lowPoints(h)
	var basins []basin
	for _, lp := range lps {
		var b basin
		visited := map[point]bool{
			lp: true,
		}
		q := make(chan point, 100)
		q <- lp
		done := false
		for !done {
			select {
			case p := <-q:
				b = append(b, p)
				for _, n := range neighbors(h, p) {
					height := h[n.y][n.x]
					if !visited[n] && height != 9 {
						q <- n
						visited[n] = true
					}
				}
			default:
				close(q)
				done = true
			}
		}
		basins = append(basins, b)
	}
	return basins
}

func main() {
	part1 := func(in string) {
		h := createHeightmap(in)
		fmt.Println(sum(riskLevels(h, lowPoints(h))...))
	}
	part1(testInput)
	b, _ := ioutil.ReadFile("day09/input.txt")
	part1(string(b))

	part2 := func(in string) {
		h := createHeightmap(in)
		fmt.Println(product(sizes(threeLargest(basinsOf(h)))...))
	}
	part2(testInput)
	part2(string(b))
}
