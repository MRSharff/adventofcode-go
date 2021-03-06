package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

var testInput = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

type point struct {
	x, y int
}
type segment [2]point

func solve(input string, doDiagonals bool) int {
	entries := strings.Split(input, "\n")
	segments := make([]segment, len(entries))
	var horizontals []*segment
	var verticals []*segment
	var diagonals []*segment
	for i, entry := range entries {
		s := segment{}
		_, _ = fmt.Sscanf(entry, "%d,%d -> %d,%d", &s[0].x, &s[0].y, &s[1].x, &s[1].y)
		segments[i] = s
		if s[0].y == s[1].y {
			horizontals = append(horizontals, &segments[i])
		} else if s[0].x == s[1].x {
			verticals = append(verticals, &segments[i])
		} else {
			diagonals = append(diagonals, &segments[i])
		}
	}

	// reorder the points in a segment so that p1.x < p2.x for horizontals
	// and p1.y < p2.y for verticals, so we don't have to check later (seems to improve time by about 1ms, 12ms vs 13ms)
	reorder := func(segs []*segment, doReorder func(seg *segment) bool) {
		for i := 0; i < len(segs); i++ {
			s := segs[i]
			if doReorder(s) {
				temp := s[0]
				s[0] = s[1]
				s[1] = temp
			}
		}
	}
	reorder(horizontals, func(s *segment) bool { return s[0].x > s[1].x })
	reorder(verticals, func(s *segment) bool { return s[0].y > s[1].y })

	// reorder diagonals, so we only have to check ys for which direction to go
	reorder(diagonals, func(s *segment) bool { return s[0].x > s[1].x })

	seen := map[point]int{}

	var start, end point
	for i := 0; i < len(horizontals); i++ {
		s := horizontals[i]
		start = s[0]
		end = s[1]
		y := start.y
		for x := start.x; x <= end.x; x++ {
			p := point{x, y}
			seen[p]++
		}
	}

	for i := 0; i < len(verticals); i++ {
		s := verticals[i]
		start = s[0]
		end = s[1]
		x := start.x
		for y := start.y; y <= end.y; y++ {
			p := point{x, y}
			seen[p]++
		}
	}

	if doDiagonals {
		for i := 0; i < len(diagonals); i++ {
			s := diagonals[i]
			start = s[0]
			end = s[1]
			increase := func(p *point) { p.x++; p.y++ }
			decrease := func(p *point) { p.x++; p.y-- }
			var move func(p *point)
			if start.y < end.y {
				move = increase
			} else {
				move = decrease
			}
			for p := start; p != end; move(&p) {
				seen[p]++
			}
			seen[end]++
		}
	}

	overlappedPointCount := 0
	for _, v := range seen {
		if v > 1 {
			overlappedPointCount++
		}
	}
	return overlappedPointCount
}

// Day 5: Hydrothermal Venture
func main() {
	got := solve(testInput, false)
	expected := 5
	if expected != got {
		fmt.Printf("expected %d, got %d\n", expected, got)
	} else {
		fmt.Println("Test Passed!")
	}

	b, _ := ioutil.ReadFile("day05/input.txt")
	in := string(b)

	answer := func() int {
		start := time.Now()
		a := solve(in, false)
		fmt.Println(time.Since(start).String())
		return a
	}
	fmt.Println(answer())

	got = solve(testInput, true)
	expected = 12
	if expected != got {
		fmt.Printf("expected %d, got %d\n", expected, got)
	} else {
		fmt.Println("Test Passed!")
	}

	part2 := func() int {
		start := time.Now()
		a := solve(in, true)
		fmt.Println(time.Since(start).String())
		return a
	}
	fmt.Println(part2())
}
