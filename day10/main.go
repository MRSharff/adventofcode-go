package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

var testInput = `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`

func main() {
	analyse(strings.Split(testInput, "\n"))

	b, _ := ioutil.ReadFile("day10/input.txt")
	analyse(strings.Split(string(b), "\n"))
}

func analyse(navSubsystem []string) {
	syntaxErrorScore := 0
	var completionScores []int

	for _, line := range navSubsystem {
		var stack []rune
		var isCorrupt = false
		for _, c := range line {
			closer, isOpener := closing[c]
			if isOpener {
				stack = push(stack, closer)
			} else {
				var expectedCloser rune
				stack, expectedCloser = pop(stack)
				isCorrupt = expectedCloser != c
				if isCorrupt {
					syntaxErrorScore += corruptionPoints[c]
					break
				}
			}
		}

		if isCorrupt {
			continue
		}
		score := 0
		var closer rune
		l := len(stack)
		for i := 0; i < l; i++ {
			stack, closer = pop(stack)
			score *= 5
			score += autocompletePoints[closer]
		}
		completionScores = append(completionScores, score)
	}

	fmt.Println("Syntax Error Score:", syntaxErrorScore)

	sort.Ints(completionScores)
	middle := len(completionScores) / 2
	fmt.Println("Middle Completion String Score:", completionScores[middle])
}

func push(stack []rune, r rune) []rune {
	stack = append(stack, r)
	return stack
}

func pop(stack []rune) ([]rune, rune) {
	var c rune
	if len(stack) > 0 {
		end := len(stack) - 1
		return stack[:end], stack[end]
	}
	return stack, c
}

var closing = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

var corruptionPoints = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var autocompletePoints = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}
