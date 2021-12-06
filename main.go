package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func getInput(day int, part int) string {
	input, err := ioutil.ReadFile(fmt.Sprintf("inputs/day%d-%d.txt", day, part))
	if err != nil {
		panic(err)
	}
	return string(input)
}

// Sonar Sweep
// https://adventofcode.com/2021/day/1
func day1() {
	reportToDepthList := func(sonarSweepReport string) []int {
		lines := strings.Fields(sonarSweepReport)
		depths := make([]int, len(lines))
		for i, line := range lines {
			var err error
			depths[i], err = strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
		}
		return depths
	}

	// part 1 solution
	countIncreasing := func(sonarSweepReport string) int {
		lastDepth := math.MaxInt64
		increaseCount := 0
		for _, depth := range reportToDepthList(sonarSweepReport) {
			if depth > lastDepth {
				increaseCount++
			}
			lastDepth = depth
		}
		return increaseCount
	}

	// part 1 tests
	testInput := "199\n200\n208\n210\n200\n207\n240\n269\n260\n263"
	expected := 7
	got := countIncreasing(testInput)
	if expected != got {
		fmt.Printf("expected %d, got %d\n", expected, got)
	} else {
		fmt.Println("test passed")
	}

	// part 1 answers
	adventInput := getInput(1, 1)
	answer := countIncreasing(adventInput)
	fmt.Println(answer)

	// part 2 solution
	threeMeasurementSlidingWindow := func(sonarSweepReport string) int {
		depths := reportToDepthList(sonarSweepReport)
		windowSize := 3
		increaseCount := 0
		for i := 0; i < len(depths)-windowSize; i++ {
			outgoing := depths[i]
			incoming := depths[i+windowSize]
			if incoming > outgoing {
				increaseCount++
			}
		}
		return increaseCount
	}

	// part 2 tests
	expected = 5
	got = threeMeasurementSlidingWindow(testInput)
	if expected != got {
		fmt.Printf("expected %d, got %d\n", expected, got)
	} else {
		fmt.Println("test passed")
	}

	// part 2 result
	part2Answer := threeMeasurementSlidingWindow(adventInput)
	fmt.Println(part2Answer)
}

func day2() {
	type command struct {
		direction string
		units     int
	}
	readCommands := func(input string) []command {
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

	positionAfterCommands := func(input string) (horizontal int, depth int) {
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

	horizontal, depth := positionAfterCommands(getInput(2, 1))
	answer := horizontal * depth
	fmt.Println(answer)

	positionAfterCommandsIncludingAim := func(input string) (horizontal int, depth int) {
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

	expectedHorizontal, expectedDepth = 15, 10
	gotHorizontal, gotDepth = positionAfterCommandsIncludingAim("forward 5\ndown 5\nforward 8\nup 3\ndown 8\nforward 2")

	expectedAnswer := 15 * 60 // = 900
	gotAnswer := gotHorizontal * gotDepth
	if expectedAnswer != gotAnswer {
		fmt.Printf("expected %d, got %d\n", expectedAnswer, gotAnswer)
	} else {
		fmt.Println("tests passed")
	}

	h, d := positionAfterCommandsIncludingAim(getInput(2, 1))
	fmt.Println(h * d)
}

func main() {
	//day1()
	day2()
}
