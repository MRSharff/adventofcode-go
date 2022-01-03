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

type fold func(paper) paper

type dot struct {
	x, y int
}

type paper struct {
	width, height int
	dots          map[dot]bool
}

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
	foldedPaper := folds[0](p)
	fmt.Println(foldedPaper)
	fmt.Println("Total dots,", len(foldedPaper.dots))
}

func part2(in string) {
	p, folds := getPaperAndInstructions(in)
	foldedPaper := p
	for _, f := range folds {
		foldedPaper = f(foldedPaper)
	}
	fmt.Println(foldedPaper)
	fmt.Println("Total dots,", len(foldedPaper.dots))
}

func (p *paper) addDot(x, y int) {
	if p.width < x {
		p.width = x
	}
	if p.height < y {
		p.height = y
	}
	if p.dots == nil {
		p.dots = make(map[dot]bool)
	}
	p.dots[dot{x, y}] = true
}

func (p paper) foldUp(pos int) paper {
	var foldedPaper paper
	for d := range p.dots {
		x, y := d.x, d.y
		if y < pos {
			foldedPaper.addDot(x, y)
		} else {
			foldedPaper.addDot(x, pos-(y-pos))
		}
	}
	foldedPaper.height = pos - 1
	return foldedPaper
}

func (p paper) foldLeft(pos int) paper {
	var foldedPaper paper
	for d := range p.dots {
		x, y := d.x, d.y
		if x < pos {
			foldedPaper.addDot(x, y)
		} else {
			foldedPaper.addDot(pos-(x-pos), y)
		}
	}
	foldedPaper.width = pos - 1
	return foldedPaper
}

func (p paper) String() string {
	sb := strings.Builder{}
	for y := 0; y <= p.height; y++ {
		for x := 0; x <= p.width; x++ {
			if p.dots[dot{x, y}] {
				sb.WriteRune('#')
			} else {
				sb.WriteRune('.')
			}
		}
		if y < p.height {
			sb.WriteRune('\n')
		}
	}
	return sb.String()
}

func getPaperAndInstructions(in string) (paper, []fold) {
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
			folds = append(folds, func(p paper) paper {
				return p.foldUp(position)
			})
		} else {
			folds = append(folds, func(p paper) paper {
				return p.foldLeft(position)
			})
		}
	}

	var ppr paper
	for _, p := range pointStrings {
		var x, y int
		scanned, err := fmt.Sscanf(p, "%d,%d", &x, &y)
		if err != nil {
			panic(err)
		}
		if scanned != 2 {
			panic("Did not find x and y:" + p)
		}
		ppr.addDot(x, y)
	}
	return ppr, folds
}
