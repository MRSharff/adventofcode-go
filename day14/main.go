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

var testInput = `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`

// Day 14: Extended Polymerization
// https://adventofcode.com/2021/day/14
func main() {
	test()
	part1(input)
	part2(testInput)
	part2(input)
}

type element string
type pair string

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

func part2(in string) {
	// Think about this in terms of what each pair produces each step. Then we have to do at most 40 * len(rules) iterations.
	//
	// keep track of how many other pairs will be produced for each step,
	// this becomes the multiplier for how many new elements are seen in the next step.
	polymer, rules := getPolymer(in), getRules(in)
	totals := make(map[element]int)
	pairCounts := make(map[pair]int)
	for _, p := range pairs(polymer) {
		pairCounts[p]++
	}
	for _, c := range polymer {
		totals[c]++
	}

	for i := 0; i < 40; i++ {
		newPairCounts := make(map[pair]int)
		for p, count := range pairCounts {
			e := rules[p]
			totals[e] += count
			leftPair := pair(element(p[0]) + e)
			rightPair := pair(e + element(p[1]))
			newPairCounts[leftPair] += count
			newPairCounts[rightPair] += count
		}
		pairCounts = newPairCounts
	}

	fmt.Println(max(totals) - min(totals))
}

func pairs(polymer []element) []pair {
	var ps []pair
	for i := 0; i < len(polymer)-1; i++ {
		p := pair(polymer[i] + polymer[i+1])
		ps = append(ps, p)
	}
	return ps
}

func min(totals map[element]int) int {
	m := math.MaxInt64
	for _, total := range totals {
		if total < m {
			m = total
		}
	}
	return m
}

func max(totals map[element]int) int {
	m := math.MinInt64
	for _, total := range totals {
		if total > m {
			m = total
		}
	}
	return m
}

func test() {
	rules := map[pair]element{
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
	polymer := []element{"N", "N", "C", "B"}
	for i := 0; i < 10; i++ {
		polymer = react(polymer, rules)
	}

	counts := countElements(polymer)
	if max(counts)-min(counts) != 1588 {
		panic("Test failed")
	}
}

func react(polymer []element, rules map[pair]element) []element {
	for i := 0; i < len(polymer)-1; i += 2 {
		p := pair(polymer[i] + polymer[i+1])
		e := rules[p]
		polymer = append(polymer[:i+1], polymer[i:]...)
		polymer[i+1] = e
	}
	return polymer
}

func countElements(polymer []element) map[element]int {
	counts := make(map[element]int)
	for _, e := range polymer {
		counts[e]++
	}
	return counts
}

func getRules(in string) map[pair]element {
	rules := make(map[pair]element)
	split := strings.Split(in, "\n")
	ruleStrings := split[2:]
	for _, s := range ruleStrings {
		var p pair
		var e element
		_, err := fmt.Sscanf(s, "%s -> %s", &p, &e)
		if err != nil {
			panic(err)
		}
		rules[p] = e
	}
	return rules
}

func getPolymer(in string) []element {
	split := strings.Split(in, "\n")
	var polymer []element
	for _, c := range split[0] {
		polymer = append(polymer, element(c))
	}
	return polymer
}
