package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"io"
	"log"
	"strings"
)

//go:embed input.txt
var input []byte

func main() {
	var part1TestInput = strings.NewReader(`1abc2
	pqr3stu8vwx
	a1b2c3d4e5f
	treb7uchet`)

	if part1(part1TestInput) != 142 {
		log.Fatal("expected test input to be 142")
	}

	log.Println(part1(bytes.NewReader(input)))

	var part2TestInput = strings.NewReader(`two1nine
	eightwothree
	abcone2threexyz
	xtwone3four
	4nineeightseven2
	zoneight234
	7pqrstsixteen`)

	if part2(part2TestInput) != 281 {
		log.Fatal("expected 281 for part 2 test")
	} else {
		log.Println("part 2 test succeeded")
	}

	log.Println(part2(bytes.NewReader(input)))

}

func part1(calibrationDocument io.Reader) int {
	var sum int
	scanner := bufio.NewScanner(calibrationDocument)
	for scanner.Scan() {
		line := scanner.Text()
		start := 0
		for line[start] < '0' || '9' < line[start] {
			start++
		}
		end := len(line) - 1
		for line[end] < '0' || '9' < line[end] {
			end--
		}
		sum += 10*int(line[start]-'0') + int(line[end]-'0')
	}
	return sum
}

func part2(calibrationDocument io.Reader) int {

	var sum int
	scanner := bufio.NewScanner(calibrationDocument)
	for scanner.Scan() {
		line := scanner.Text()

		var startDigit int
		i := 0
		for ; i < len(line); i++ {
			if '0' <= line[i] && line[i] <= '9' {
				startDigit = int(line[i] - '0')
				break
			}
			switch {
			case i+len("zero") < len(line) && line[i:i+len("zero")] == "zero":
				startDigit = 0
			case i+len("one") < len(line) && line[i:i+len("one")] == "one":
				startDigit = 1
			case i+len("two") < len(line) && line[i:i+len("two")] == "two":
				startDigit = 2
			case i+len("three") < len(line) && line[i:i+len("three")] == "three":
				startDigit = 3
			case i+len("four") < len(line) && line[i:i+len("four")] == "four":
				startDigit = 4
			case i+len("five") < len(line) && line[i:i+len("five")] == "five":
				startDigit = 5
			case i+len("six") < len(line) && line[i:i+len("six")] == "six":
				startDigit = 6
			case i+len("seven") < len(line) && line[i:i+len("seven")] == "seven":
				startDigit = 7
			case i+len("eight") < len(line) && line[i:i+len("eight")] == "eight":
				startDigit = 8
			case i+len("nine") < len(line) && line[i:i+len("nine")] == "nine":
				startDigit = 9
			default:
				continue
			}
			break
		}
		var endDigit int
		i = len(line) - 1
		for ; i >= 0; i-- {
			if '0' <= line[i] && line[i] <= '9' {
				endDigit = int(line[i] - '0')
				break
			}
			switch {
			case i-len("zero")+1 >= 0 && line[i-len("zero")+1:i+1] == "zero":
				endDigit = 0
			case i-len("one")+1 >= 0 && line[i-len("one")+1:i+1] == "one":
				endDigit = 1
			case i-len("two")+1 >= 0 && line[i-len("two")+1:i+1] == "two":
				endDigit = 2
			case i-len("three")+1 >= 0 && line[i-len("three")+1:i+1] == "three":
				endDigit = 3
			case i-len("four")+1 >= 0 && line[i-len("four")+1:i+1] == "four":
				endDigit = 4
			case i-len("five")+1 >= 0 && line[i-len("five")+1:i+1] == "five":
				endDigit = 5
			case i-len("six")+1 >= 0 && line[i-len("six")+1:i+1] == "six":
				endDigit = 6
			case i-len("seven")+1 >= 0 && line[i-len("seven")+1:i+1] == "seven":
				endDigit = 7
			case i-len("eight")+1 >= 0 && line[i-len("eight")+1:i+1] == "eight":
				endDigit = 8
			case i-len("nine")+1 >= 0 && line[i-len("nine")+1:i+1] == "nine":
				endDigit = 9
			default:
				continue
			}
			break
		}
		sum += 10*startDigit + endDigit
	}
	return sum
}
