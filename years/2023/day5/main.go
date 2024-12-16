package main

import (
	"bytes"
	_ "embed"
	"io"
	"log"

	"github.com/MRSharff/adventofcode-go/inputs"
)

var testInput = `
seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
`

//go:embed input.txt
var input []byte

type almanac struct {
	seeds            []int
	seedToSoil       []M
	soilToFertilizer []M
}

type M struct {
	destinationRangeStart []int
	sourceRangeStart      []int
	rangeLength           []int
}

func convert(source M, n int) int {
	k := len(source.sourceRangeStart)
	for i := 0; i < k; i++ {
		sourceStart := source.sourceRangeStart[i]
		destStart := source.destinationRangeStart[i]
		length := source.rangeLength[i]

		if n < sourceStart || sourceStart+length < n {
			continue
		}

		diff := n - sourceStart
		return destStart + diff
	}
	return n
}

func main() {

	sourceTest := M{
		destinationRangeStart: []int{50, 52},
		sourceRangeStart:      []int{98, 50},
		rangeLength:           []int{2, 48},
	}
	testConvert := convert(sourceTest, 53)
	if testConvert != 55 {
		log.Fatal("expected 55, got", testConvert)
	}

	testGot := part1(inputs.NewBytesReader([]byte(testInput)))
	if testGot != 13 {
		log.Fatal("expected 13, got", testGot)
	}

	log.Println(part1(bytes.NewReader(input)))

	testGot = part2(inputs.NewBytesReader([]byte(testInput)))
	if testGot != 30 {
		log.Fatal("expected 30, got", testGot)
	}

	log.Println(part2(bytes.NewReader(input)))
}

/*
The idea here is to have a multiplier for each card. Card 1 technically does not have a multiplier.

When a card has winning matches you multiply
*/
func part1(r io.Reader) int {
	return 0
}

func part2(r io.Reader) int {
	return 0
}
