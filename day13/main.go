package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var testInput = `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`

var input = func() string {
	b, _ := ioutil.ReadFile("day13/input.txt")
	return string(b)
}()

// Day 13: Transparent Origami
// https://adventofcode.com/2021/day/13
func main() {
	part1(testInput)
	part1(input)
	part2(testInput)
	part2(input)
}

func part1(in string) {
	p, folds := getPaperAndInstructions(in)
	foldedPaper := foldPaper(p, folds[0])
	fmt.Println("Total dots,", len(foldedPaper))
}

func part2(in string) {
	p, folds := getPaperAndInstructions(in)
	foldedPaper := p
	for _, f := range folds {
		foldedPaper = foldPaper(foldedPaper, f)
	}
	printPaper(foldedPaper)
	fmt.Println("Total dots,", len(foldedPaper))
}

type point struct {
	x, y int
}
type paper map[point]bool

func add(p paper, x, y int) {
	p[point{x, y}] = true
}

type fold func(x, y int) (int, int)

func up(pos int) fold {
	return func(x, y int) (int, int) {
		if y < pos {
			return x, y
		}
		return x, pos - (y - pos)
	}
}

func left(pos int) fold {
	return func(x, y int) (int, int) {
		if x < pos {
			return x, y
		}
		return pos - (x - pos), y
	}
}

func foldPaper(p paper, f fold) paper {
	foldedPaper := make(map[point]bool)
	for pt := range p {
		x, y := pt.x, pt.y
		x2, y2 := f(x, y)
		add(foldedPaper, x2, y2)
	}
	return foldedPaper
}

func getPaperAndInstructions(in string) (map[point]bool, []fold) {
	split := strings.Split(in, "\n\n")
	pointStrings := strings.Split(split[0], "\n")
	foldStrings := strings.Split(split[1], "\n")
	var folds []fold
	for _, fs := range foldStrings {
		var axis rune
		var position int

		sc, err := fmt.Sscanf(fs, "fold along %c=%d", &axis, &position)
		if err != nil {
			panic(err)
		}
		if sc != 2 {
			panic("did not scan all," + fs)
		}

		if axis == 'y' {
			folds = append(folds, up(position))
		} else {
			folds = append(folds, left(position))
		}
	}

	ppr := make(map[point]bool)
	for _, p := range pointStrings {
		var x, y int
		scanned, err := fmt.Sscanf(p, "%d,%d", &x, &y)
		if err != nil {
			panic(err)
		}
		if scanned != 2 {
			panic("Did not find x and y:" + p)
		}
		ppr[point{x, y}] = true
	}
	return ppr, folds
}

func printPaper(p paper) {
	maxX, maxY := 0, 0
	for pt := range p {
		x, y := pt.x, pt.y
		if y > maxY {
			maxY = y
		}
		if x > maxX {
			maxX = x
		}
	}

	sb := strings.Builder{}
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			_, dotted := p[point{x, y}]
			if dotted {
				sb.WriteRune('#')
			} else {
				sb.WriteRune('.')
			}
		}
		if y < maxY {
			sb.WriteRune('\n')
		}
	}
	fmt.Println(sb.String())
}
