package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var day3Input = func() string {
	b, _ := ioutil.ReadFile("day03/input.txt")
	return string(b)
}()

// Binary Diagnostic
// https://adventofcode.com/2021/day/3
func main() {
	// input: diagnostic report is a list of binary numbers
	testInput := "00100\n11110\n10110\n10111\n10101\n01111\n00111\n11100\n10000\n11001\n00010\n01010"

	powerConsumption := func(input string) int {
		diagnosticReport := strings.Fields(input)
		size := len(diagnosticReport[0])
		sums := make([]int, size)
		for _, bits := range diagnosticReport {
			for i := 0; i < size; i++ {
				if bits[i] == '1' {
					sums[i]++
				}
			}
		}

		mostCommonBits := make([]int, size)
		half := len(diagnosticReport) / 2
		for i := 0; i < size; i++ {
			oneMostCommon := sums[i] > half
			if oneMostCommon {
				mostCommonBits[i] = 1
			}
		}

		gammaRate := 0
		for i, shiftAmount := 0, size-1; i < size; i, shiftAmount = i+1, shiftAmount-1 {
			bit := mostCommonBits[i]
			gammaRate |= bit << shiftAmount
		}

		// we don't have uint5 so let's just work with uint16 and shift x amount left and right to clear the first x unused bits.
		x := 16 - size
		epsilonRate := int(^uint16(gammaRate) << x >> x)

		return epsilonRate * gammaRate
	}

	expected := 198
	got := powerConsumption(testInput)
	if expected != got {
		fmt.Printf("expected %d, got %d", expected, got)
	} else {
		fmt.Println("test passed")
	}

	fmt.Println(powerConsumption(day3Input))

	part2 := func(input string) int {
		diagnosticReport := strings.Fields(input)
		bitLength := len(diagnosticReport[0])

		oxygenGeneratorList := make([]string, len(diagnosticReport))
		copy(oxygenGeneratorList, diagnosticReport)

		// filter loop
		for i := 0; i < bitLength && len(oxygenGeneratorList) > 1; i++ {
			half := len(oxygenGeneratorList) / 2
			zeroes := make([]string, 0, half)
			ones := make([]string, 0, half)

			for _, bits := range oxygenGeneratorList {
				if bits[i] == '0' {
					zeroes = append(zeroes, bits)
				} else {
					ones = append(ones, bits)
				}
			}
			if len(ones) >= len(zeroes) {
				oxygenGeneratorList = ones
			} else {
				oxygenGeneratorList = zeroes
			}
		}

		co2ScrubberList := make([]string, len(diagnosticReport))
		copy(co2ScrubberList, diagnosticReport)

		// filter loop
		for i := 0; i < bitLength && len(co2ScrubberList) > 1; i++ {
			half := len(co2ScrubberList) / 2
			zeroes := make([]string, 0, half)
			ones := make([]string, 0, half)

			for _, bits := range co2ScrubberList {
				if bits[i] == '0' {
					zeroes = append(zeroes, bits)
				} else {
					ones = append(ones, bits)
				}
			}
			if len(ones) < len(zeroes) {
				co2ScrubberList = ones
			} else {
				co2ScrubberList = zeroes
			}
		}

		oxygenGeneratorRating, err := strconv.ParseInt(oxygenGeneratorList[0], 2, 0)
		if err != nil {
			panic(err)
		}
		c02ScrubberRating, err := strconv.ParseInt(co2ScrubberList[0], 2, 0)
		if err != nil {
			panic(err)
		}
		return int(oxygenGeneratorRating) * int(c02ScrubberRating)
	}
	lifeSupportRating := part2(testInput)
	expected = 230
	if expected != lifeSupportRating {
		fmt.Printf("expected %d, got %d\n", expected, lifeSupportRating)
	} else {
		fmt.Println("test passed")
	}
	fmt.Println(part2(day3Input))
}
