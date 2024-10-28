package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var input = func() string {
	b, _ := ioutil.ReadFile("day02/input.txt")
	return string(b)
}()

// Dive
// https://adventofcode.com/2021/day/2
func main() {
	part1test()
	answer := part1(input)
	fmt.Println(answer)

	part2test()
	fmt.Println(part2(input))
}

func part1(in string) int {
	horizontal, depth := positionAfterCommands(in)
	return horizontal * depth
}

func part1test() {
	expectedHorizontal, expectedDepth := 15, 10
	gotHorizontal, gotDepth := positionAfterCommands("forward 5\ndown 5\nforward 8\nup 3\ndown 8\nforward 2")

	if expectedHorizontal != gotHorizontal {
		fmt.Printf("expected horizontal %d, got %d\n", expectedHorizontal, gotHorizontal)
		panic("test failed")
	}
	if expectedDepth != gotDepth {
		fmt.Printf("expected depth %d, got %d\n", expectedDepth, gotDepth)
		panic("test failed")
	}
	fmt.Println("tests passed")
}

func part2(in string) int {
	horizontal, depth := positionAfterCommandsIncludingAim(in)
	return horizontal * depth
}

func part2test() {
	gotHorizontal, gotDepth := positionAfterCommandsIncludingAim("forward 5\ndown 5\nforward 8\nup 3\ndown 8\nforward 2")

	expectedAnswer := 15 * 60 // = 900
	gotAnswer := gotHorizontal * gotDepth
	if expectedAnswer != gotAnswer {
		fmt.Printf("expected %d, got %d\n", expectedAnswer, gotAnswer)
	} else {
		fmt.Println("tests passed")
	}
}

func positionAfterCommandsIncludingAim(input string) (horizontal int, depth int) {
	aim := 0
	commands := readCommands(input)
	for _, c := range commands {
		x := c.units
		switch c.direction {
		case "forward":
			horizontal += x
			depth += aim * x
		case "down":
			aim += x
		case "up":
			aim -= x
		}
	}
	return
}

type command struct {
	direction string
	units     int
}

func readCommands(input string) []command {
	lines := strings.Fields(input)
	commands := make([]command, len(lines)/2)
	for i := 0; i < len(lines)-1; i += 2 {
		direction := lines[i]
		units, err := strconv.Atoi(lines[i+1])
		if err != nil {
			panic(err)
		}
		commands[i/2] = command{direction, units}
	}
	return commands
}

func positionAfterCommands(input string) (horizontal int, depth int) {
	commands := readCommands(input)
	for _, c := range commands {
		switch c.direction {
		case "forward":
			horizontal += c.units
		case "down":
			depth += c.units
		case "up":
			depth -= c.units
		}
	}
	return
}
