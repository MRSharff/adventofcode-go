package main

import (
	"fmt"
	"io/ioutil"
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

var uniqueSignalCounts = map[int]int{
	len(digits[1]): 1,
	len(digits[4]): 4,
	len(digits[7]): 7,
	len(digits[8]): 8,
}

func main() {
	b, _ := ioutil.ReadFile("day8/input.txt")
	entries := strings.Split(string(b), "\n")
	occurrences := 0
	for _, entry := range entries {
		split := strings.Split(entry, " ")
		_, output := split[:10], split[11:]
		for i := 0; i < len(output); i++ {
			_, isUnique := uniqueSignalCounts[len(output[i])]
			if isUnique {
				occurrences++
			}
		}
	}
	fmt.Println(occurrences)
}
