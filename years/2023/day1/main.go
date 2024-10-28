package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"io"
	"log"
)

//go:embed input.txt
var input []byte

var testInput = []byte(`1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`)

func main() {
	if part1(bytes.NewReader(testInput)) != 142 {
		log.Fatal("expected test input to be 142")
	}

	log.Println(part1(bytes.NewReader(input)))
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
