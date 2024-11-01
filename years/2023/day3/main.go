package main

import (
	"bytes"
	_ "embed"
	"errors"
	"io"
	"log"
	"strconv"
	"strings"
	"unicode"
)

//go:embed input.txt
var input []byte

func main() {
	var exampleEngineSchematic = `
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

	r := strings.NewReader(exampleEngineSchematic)
	newlinebuf := []byte{'\n'}
	r.Read(newlinebuf) // consume the added newline

	got := part1(r)
	if got != 4361 {
		log.Fatal("expected part 1 test to be 4361, got: ", got)
	}

	log.Println(part1(bytes.NewReader(input)))

	r = strings.NewReader(exampleEngineSchematic)
	r.Read(newlinebuf) // consume the added newline

	got2 := part2(r)
	if got2 != 467835 {
		log.Fatal("expected part 1 test to be 4361, got: ", got2)
	}

	log.Println(part2(bytes.NewReader(input)))
}

func part1(input io.Reader) int {
	var sum int

	var lines [][]byte
	buf := []byte{'.'}
	var buf2 []byte

	for {
		_, err := input.Read(buf)
		if err != nil {
			panic(err)
		}
		buf2 = append(buf2, buf[0])
		if buf[0] == '\n' {
			break
		}
	}
	l := make([]byte, len(buf2))
	for i := 0; i < 3; i++ {
		lines = append(lines, make([]byte, len(l)))
	}

	// fake a top line so we don't have to do bounds checking, we'll just start at line 1
	for i := 0; i < len(lines[0]); i++ {
		lines[0][i] = '.'
	}

	for i, c := range buf2 {
		lines[1][i] = c
	}

	run := true
	for run {
		if _, err := input.Read(lines[2]); err != nil {
			if errors.Is(err, io.EOF) {
				for i := 0; i < len(lines[0]); i++ {
					lines[0][i] = '.'
				}
				run = false
			}
		}
		for i, c := range lines[1] {
			if unicode.IsDigit(rune(c)) || c == '.' || c == '\n' {
				continue
			}

			var parts []part
			for _, neighbor := range neighbors(i, lines) {
				if unicode.IsDigit(rune(neighbor.b)) {
					p := getPartNumber(neighbor, lines)
					parts = append(parts, p)
				}
			}
			type key struct {
				line, start int
			}
			dedupedParts := make(map[key]int)
			for _, p := range parts {
				dedupedParts[key{p.line, p.start}] = p.number
			}
			for _, v := range dedupedParts {
				sum += v
			}
		}
		lines = append(lines[1:], lines[0])
	}

	return sum
}

type neighb struct {
	b    byte
	x, y int
}

func neighbors(k int, lines [][]byte) []neighb {
	var neibs []neighb
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			x := 1 + i
			y := k + j
			neibs = append(neibs, neighb{lines[x][y], x, y})
		}
	}

	return neibs
}

type part struct {
	start, end, line int
	number           int
}

func getPartNumber(n neighb, lines [][]byte) part {
	line := lines[n.x]
	start := n.y
	for start-1 >= 0 && unicode.IsDigit(rune(line[start-1])) {
		start--
	}
	end := n.y
	for end < len(line) && unicode.IsDigit(rune(line[end])) {
		end++
	}
	pn, err := strconv.Atoi(string(line[start:end]))
	if err != nil {
		panic(err)
	}
	return part{start, end, n.x, pn}
}

func part2(input io.Reader) int {
	var sum int

	var lines [][]byte
	buf := []byte{'.'}
	var buf2 []byte

	for {
		_, err := input.Read(buf)
		if err != nil {
			panic(err)
		}
		buf2 = append(buf2, buf[0])
		if buf[0] == '\n' {
			break
		}
	}
	l := make([]byte, len(buf2))
	for i := 0; i < 3; i++ {
		lines = append(lines, make([]byte, len(l)))
	}

	// fake a top line so we don't have to do bounds checking, we'll just start at line 1
	for i := 0; i < len(lines[0]); i++ {
		lines[0][i] = '.'
	}

	for i, c := range buf2 {
		lines[1][i] = c
	}

	run := true
	for run {
		if _, err := input.Read(lines[2]); err != nil {
			if errors.Is(err, io.EOF) {
				for i := 0; i < len(lines[0]); i++ {
					lines[0][i] = '.'
				}
				run = false
			}
		}
		for i, c := range lines[1] {
			if unicode.IsDigit(rune(c)) || c == '.' || c == '\n' {
				continue
			}

			var parts []part
			for _, neighbor := range neighbors(i, lines) {
				if unicode.IsDigit(rune(neighbor.b)) {
					p := getPartNumber(neighbor, lines)
					parts = append(parts, p)
				}
			}
			type key struct {
				line, start int
			}
			dedupedParts := make(map[key]int)
			for _, p := range parts {
				dedupedParts[key{p.line, p.start}] = p.number
			}

			if len(dedupedParts) == 2 {
				gearRatio := 1
				for _, v := range dedupedParts {
					gearRatio *= v
				}
				sum += gearRatio
			}

		}
		lines = append(lines[1:], lines[0])
	}

	return sum
}
