package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

var adventInput = func() string {
	b, _ := ioutil.ReadFile("day01/input.txt")
	return string(b)
}()

// Sonar Sweep
// https://adventofcode.com/2021/day/1
func main() {
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
