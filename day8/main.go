package main

import (
	"fmt"
	"io/ioutil"
	"sort"
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
	len(digits[1]): 1, // len = 2
	len(digits[4]): 4, // len = 4
	len(digits[7]): 7, // len = 3
	len(digits[8]): 8, // len = 7
}

var testEntry = "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf"

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

func main() {
	b, _ := ioutil.ReadFile("day8/test_input.txt")
	entries := strings.Split(string(b), "\n")
	occurrences := 0
	for _, entry := range entries[:1] {
		split := strings.Split(entry, " ")
		patterns, output := split[:10], split[11:]

		all := append(patterns, output...)

		possible := map[string]map[string]bool{
			"a": {"a": true, "b": true, "c": true, "d": true, "e": true, "f": true, "g": true},
			"b": {"a": true, "b": true, "c": true, "d": true, "e": true, "f": true, "g": true},
			"c": {"a": true, "b": true, "c": true, "d": true, "e": true, "f": true, "g": true},
			"d": {"a": true, "b": true, "c": true, "d": true, "e": true, "f": true, "g": true},
			"e": {"a": true, "b": true, "c": true, "d": true, "e": true, "f": true, "g": true},
			"f": {"a": true, "b": true, "c": true, "d": true, "e": true, "f": true, "g": true},
			"g": {"a": true, "b": true, "c": true, "d": true, "e": true, "f": true, "g": true},
		}

		seen := make(map[string]bool)
		digitSet := make([]string, 0, len(digits))
		easies := make(map[int]string)
		for i := 0; i < len(all); i++ {
			p := sortLetters(all[i])
			if !seen[p] {
				seen[p] = true
				//digitSet = append(digitSet, p)
				if d, easy := uniqueSignalCounts[len(p)]; easy {
					easies[d] = p
				} else {
					digitSet = append(digitSet, p)
				}
			}
		}

		sort.Slice(digitSet, func(i, j int) bool {
			return len(digitSet[i]) < len(digitSet[j])
		})

		connections := make(map[string]string)
		for _, i := range []int{1, 7, 4, 8} {
			digit := easies[i]
			correct := digits[uniqueSignalCounts[len(digit)]]
			wires := strings.Split("a b c d e f g", " ")
			wiresOfDigit, wiresNotOfDigit := make([]string, 0), make([]string, 0)
			for _, wire := range wires {
				if strings.Contains(digit, wire) {
					// wires of the digit
					wiresOfDigit = append(wiresOfDigit, wire)
				} else {
					// wire not in the digit
					wiresNotOfDigit = append(wiresNotOfDigit, wire)
				}
			}
			for _, wire := range wiresOfDigit {
				connection := possible[wire]
				for seg := range connection {
					doRemove := !strings.Contains(correct, seg)
					if doRemove {
						delete(connection, seg)
						if len(connection) == 1 {
							for k := range connection {
								connections[wire] = k
							}
						}
					}
				}
			}
			for _, wire := range wiresNotOfDigit {
				connection := possible[wire]
				for seg := range connection {
					doRemove := strings.Contains(correct, seg)
					if doRemove {
						delete(connection, seg)
						if len(connection) == 1 {
							for k := range connection {
								connections[wire] = k
							}
						}
					}
				}
			}
			fmt.Println("Possible Connections")
			for _, wire := range "abcdefg" {
				fmt.Printf("%s : ", string(wire))
				possibles := []string{}
				for seg := range possible[string(wire)] {
					possibles = append(possibles, seg)
					//fmt.Printf("%s ", seg)
				}
				sort.Strings(possibles)
				fmt.Printf("%v\n", possibles)

				fmt.Println()
			}
		}

		fmt.Println("Onto the rest?")
	}
	fmt.Println(occurrences)
}
