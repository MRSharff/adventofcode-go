package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

var input = func() string {
	b, _ := ioutil.ReadFile("day14/input.txt")
	return string(b)
}()

// Day 14: Extended Polymerization
// https://adventofcode.com/2021/day/14
func main() {
	test()
	part1(input)
}

func part1(in string) {
	polymer, rules := getPolymer(in), getRules(in)
	for i := 0; i < 10; i++ {
		polymer = react(polymer, rules)
	}

	counts := countElements(polymer)
	leastCommon, mostCommon := math.MaxInt32, 0
	for _, count := range counts {
		if count > mostCommon {
			mostCommon = count
		}
		if count < leastCommon {
			leastCommon = count
		}
	}
	fmt.Println(mostCommon - leastCommon)
}

func test() {
	rules := map[string]string{
		"CH": "B",
		"HH": "N",
		"CB": "H",
		"NH": "C",
		"HB": "C",
		"HC": "B",
		"HN": "C",
		"NN": "C",
		"BH": "H",
		"NC": "B",
		"NB": "B",
		"BN": "B",
		"BB": "N",
		"BC": "B",
		"CC": "N",
		"CN": "C",
	}
	polymer := []string{"N", "N", "C", "B"}
	for i := 0; i < 10; i++ {
		polymer = react(polymer, rules)
	}

	counts := countElements(polymer)
	min, max := math.MaxInt32, 0
	for _, count := range counts {
		if count > max {
			max = count
		}
		if count < min {
			min = count
		}
	}
	if max-min != 1588 {
		panic("Test failed")
	}
}

func react(polymer []string, rules map[string]string) []string {
	for i := 0; i < len(polymer)-1; i += 2 {
		pair := polymer[i] + polymer[i+1]
		element := rules[pair]
		polymer = append(polymer[:i+1], polymer[i:]...)
		polymer[i+1] = element
	}
	return polymer
}

func countElements(polymer []string) map[string]int {
	counts := make(map[string]int)
	for _, element := range polymer {
		counts[element]++
	}
	return counts
}

func getRules(in string) map[string]string {
	rules := make(map[string]string)
	split := strings.Split(in, "\n\n")
	ruleStrings := strings.Split(split[1], "\n")
	for _, s := range ruleStrings {
		var pair, element string
		_, err := fmt.Sscanf(s, "%s -> %s", &pair, &element)
		if err != nil {
			panic(err)
		}
		rules[pair] = element
	}
	return rules
}

func getPolymer(in string) []string {
	split := strings.Split(in, "\n\n")
	var polymer []string
	for _, c := range split[0] {
		polymer = append(polymer, string(c))
	}
	return polymer
}
