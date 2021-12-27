package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

var digits = []string{
	"abcefg",
	"cf",
	"acdeg",
	"acdfg",
	"bcdf",
	"abdfg",
	"abdefg",
	"acf",
	"abcdefg",
	"abcdfg",
}

var segmentsToDigit = map[string]int{
	"abcefg":  0,
	"cf":      1,
	"acdeg":   2,
	"acdfg":   3,
	"bcdf":    4,
	"abdfg":   5,
	"abdefg":  6,
	"acf":     7,
	"abcdefg": 8,
	"abcdfg":  9,
}

var testEntry = "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf"

func diff(d1, d2 string) string {
	longer, shorter := d1, d2
	if len(d2) > len(d1) {
		longer, shorter = d2, d1
	}
	s := ""
	for _, c := range longer {
		if !strings.Contains(shorter, string(c)) {
			s += string(c)
		}
	}
	return s
}

func sortLetters(s string) string {
	l := make([]string, len(s))
	for i, r := range s {
		l[i] = string(r)
	}
	sort.Slice(l, func(i, j int) bool {
		return l[i] < l[j]
	})
	return strings.Join(l, "")
}

func values(m map[string]string) string {
	s := ""
	for _, v := range m {
		s += v
	}
	return s
}

func intListToInt(ints []int) int {
	s := ""
	for _, n := range ints {
		s += strconv.Itoa(n)
	}
	n, _ := strconv.Atoi(s)
	return n
}

func solveEntry(entry string) (int, int) {
	known := make(map[int]string, len("abcdefg"))
	fivesCounts := map[string]int{
		"a": 0,
		"b": 0,
		"c": 0,
		"d": 0,
		"e": 0,
		"f": 0,
		"g": 0,
	}
	sixCounts := map[string]int{
		"a": 0,
		"b": 0,
		"c": 0,
		"d": 0,
		"e": 0,
		"f": 0,
		"g": 0,
	}

	split := strings.Split(entry, " ")
	patterns, output := split[:10], split[11:]

	all := make([]string, 0, len(patterns)+len(output))
	all = append(all, patterns...)
	all = append(all, output...)
	seen := map[string]bool{}
	allSet := make([]string, 0)
	for _, digit := range all {
		sortedDigit := sortLetters(digit)
		if !seen[sortedDigit] {
			allSet = append(allSet, digit)
			seen[sortedDigit] = true
		}
	}

	lengths := map[int][]string{}

	for _, digit := range allSet {
		lengths[len(digit)] = append(lengths[len(digit)], digit)
		switch len(digit) {
		case 2:
			known[1] = digit
		case 3:
			known[7] = digit
		case 4:
			known[4] = digit
		case 5:
			for _, c := range digit {
				fivesCounts[string(c)]++
			}
		case 6:
			for _, c := range digit {
				sixCounts[string(c)]++
			}
		case 7:
			known[8] = digit
		}
	}

	digitMap := map[string]string{}

	digitMap["a"] = diff(known[7], known[1])

	// b
	fourNotInOne := diff(known[4], known[1])
	for _, c := range fourNotInOne {
		s := string(c)
		if fivesCounts[s] == 1 {
			digitMap["b"] = s
			break
		}
	}

	// c
	for _, c := range known[1] {
		s := string(c)
		if sixCounts[s] == 2 {
			digitMap["c"] = s
			break
		}
	}

	// d
	digitMap["d"] = diff(diff(known[4], known[1]), digitMap["b"])

	// e
	for c, count := range fivesCounts {
		if c != digitMap["b"] && count == 1 {
			digitMap["e"] = c
			break
		}
	}

	// f
	digitMap["f"] = strings.Replace(known[1], digitMap["c"], "", 1)

	digitMap["g"] = diff(values(digitMap), "abcdefg")

	wiresToSegments := make(map[string]string)
	for k, v := range digitMap {
		wiresToSegments[v] = k
	}
	patternNumbers := make([]int, len(patterns))
	for i, pattern := range patterns {
		s := ""
		for _, c := range pattern {
			s += wiresToSegments[string(c)]
		}
		patternNumbers[i] = segmentsToDigit[sortLetters(s)]
	}

	outputNumbers := make([]int, len(output))
	for i, pattern := range output {
		s := ""
		for _, c := range pattern {
			s += wiresToSegments[string(c)]
		}
		outputNumbers[i] = segmentsToDigit[sortLetters(s)]
	}

	return intListToInt(patternNumbers), intListToInt(outputNumbers)
}

func main() {
	b, _ := ioutil.ReadFile("day8/input.txt")
	entries := strings.Split(string(b), "\n")

	sum := 0
	for _, entry := range entries {
		_, output := solveEntry(entry)
		fmt.Println(output)
		sum += output
	}
	fmt.Println(sum)
}
