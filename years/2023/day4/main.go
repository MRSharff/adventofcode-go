package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/MRSharff/adventofcode-go/inputs"
)

var testInput = `
Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
`

//go:embed input.txt
var input []byte

func main() {
	testGot := part1(inputs.NewBytesReader([]byte(testInput)))
	if testGot != 13 {
		log.Fatal("expected 13, got", testGot)
	}

	log.Println(part1(bytes.NewReader(input)))
}

func part1(r io.Reader) int {
	scanner := bufio.NewScanner(r)

	var points int
	for scanner.Scan() {
		tokens := bufio.NewScanner(strings.NewReader(scanner.Text()))
		tokens.Split(bufio.ScanWords)
		tokens.Scan()
		tokens.Scan()

		winningNumbers := make(map[int]struct{})
		for tokens.Scan() {
			s := tokens.Text()
			if s == "|" {
				break
			}
			var n int
			fmt.Sscanf(s, "%d", &n)
			winningNumbers[n] = struct{}{}
		}
		var matches int
		for tokens.Scan() {
			s := tokens.Text()
			var n int
			fmt.Sscanf(s, "%d", &n)
			if _, ok := winningNumbers[n]; ok {
				matches++
			}
		}
		if matches > 0 {
			points += 1 << (matches - 1)
		}
	}
	return points
}
