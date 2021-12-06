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
	countIncreasing := func(sonarSweepReport string) int {
		lastDepth := math.MaxInt64
		increaseCount := 0
		for _, line := range strings.Fields(sonarSweepReport) {
			n, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			if n > lastDepth {
				increaseCount++
			}
			lastDepth = n
		}
		return increaseCount
	}

	testInput := "199\n200\n208\n210\n200\n207\n240\n269\n260\n263"
	expected := 7
	got := countIncreasing(testInput)
	if expected != got {
		fmt.Printf("expected %d, got %d\n", expected, got)
	} else {
		fmt.Println("test passed")
	}

	adventInput := getInput(1, 1)

	answer := countIncreasing(adventInput)
	fmt.Println(answer)
}

func main() {
	day1()
}
