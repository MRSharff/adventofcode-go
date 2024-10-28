package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"log"
	"strings"
)

//go:embed input.txt
var input []byte

// we got games, sets, cubes

func main() {
	part1TestInput := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\nGame 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue\nGame 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red\nGame 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red\nGame 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"

	got := part1(strings.NewReader(part1TestInput))
	if got != 8 {
		log.Fatal("expected part 1 test to be 8, got ", got)
	}

	log.Println("part 1:", part1(bytes.NewReader(input)))

	got = part2(strings.NewReader(part1TestInput))
	if got != 2286 {
		log.Fatal("expected part 2 test to be 2286, got ", got)
	}

	log.Println("part 2:", part2(bytes.NewReader(input)))
}

func part1(input io.Reader) int {
	totalCubes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	var sum int

	scanner := bufio.NewScanner(input)
	for id := 1; scanner.Scan(); id++ {
		game := scanner.Text()
		split := strings.Split(game, ": ")
		sets := strings.Split(split[1], "; ")
		possible := true
		for _, set := range sets {
			if !possible {
				break
			}
			reveals := strings.Split(set, ", ")
			for _, reveal := range reveals {
				var n int
				var color string
				_, _ = fmt.Sscanf(reveal, "%d %s", &n, &color)
				if n > totalCubes[color] {
					possible = false
					break
				}
			}
		}
		if possible {
			sum += id
		}
	}

	return sum
}

func part2(input io.Reader) int {
	var sum int

	scanner := bufio.NewScanner(input)
	for id := 1; scanner.Scan(); id++ {
		maxCubes := map[string]int{}
		game := scanner.Text()
		split := strings.Split(game, ": ")
		sets := strings.Split(split[1], "; ")
		for _, set := range sets {
			reveals := strings.Split(set, ", ")
			for _, reveal := range reveals {
				var n int
				var color string
				_, _ = fmt.Sscanf(reveal, "%d %s", &n, &color)
				if max, ok := maxCubes[color]; !ok || n > max {
					maxCubes[color] = n
				}
			}
		}
		power := 1
		for _, v := range maxCubes {
			power *= v
		}
		sum += power
	}

	return sum
}
