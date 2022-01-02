package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

const SIZE = 5

type Board struct {
	grid            [SIZE][SIZE]int
	unmarkedNumbers []int
}

func CreateBoard(numbers []int) Board {
	b := Board{}
	b.unmarkedNumbers = make([]int, SIZE*SIZE)
	if len(numbers) != 25 {
		panic("Must choose 25 numbers to fill board with")
	}
	for i, n := range numbers {
		row := i / SIZE
		col := i % SIZE
		b.grid[row][col] = n
		b.unmarkedNumbers[i] = n
	}
	return b
}

func NewBingo(filename string) (pool []int, boards []Board) {
	inputContents, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	bingoInput := string(inputContents)
	lines := strings.Split(bingoInput, "\n\n")
	poolString := lines[0]
	poolStrings := strings.Split(poolString, ",")
	pool = make([]int, len(poolStrings))
	for i, s := range poolStrings {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		pool[i] = n
	}

	boardStrings := lines[1:]
	totalBoards := len(boardStrings)
	boards = make([]Board, totalBoards)
	for i := 0; i < totalBoards; i++ {
		ns := make([]int, SIZE*SIZE)
		for j, s := range strings.Fields(boardStrings[i]) {
			n, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			ns[j] = n
		}
		boards[i] = CreateBoard(ns)
	}
	return pool, boards
}

func PlayBoard(b Board, pool []int) (winningTurn, sum, winningNumber int) {
	rowSums := make([]int, SIZE)
	colSums := make([]int, SIZE)
	for turn, n := range pool {
		for row := 0; row < SIZE; row++ {
			for col := 0; col < SIZE; col++ {
				if b.grid[row][col] == n {
					rowSums[row]++
					colSums[col]++
					removeIndex := row*SIZE + col
					b.unmarkedNumbers[removeIndex] = 0
				}
				isWinningTurn := rowSums[row] == SIZE || colSums[col] == SIZE
				if isWinningTurn {
					winningTurn = turn
					winningNumber = n
					for i := 0; i < len(b.unmarkedNumbers); i++ {
						sum += b.unmarkedNumbers[i]
					}
					return
				}
			}
		}
	}
	return 0, 0, 0
}

// Returns the turn and sum*winningNumber of the first
// and last boards to finish
func solve(filename string) (first, last []int) {
	pool, boards := NewBingo(filename)

	boardRanks := make([][]int, len(boards))
	for i := 0; i < len(boards); i++ {
		turn, sum, n := PlayBoard(boards[i], pool)
		total := sum * n
		boardRanks[i] = []int{turn, total}
	}

	sort.Slice(boardRanks, func(i, j int) bool {
		return boardRanks[i][0] < boardRanks[j][0]
	})

	return boardRanks[0], boardRanks[len(boardRanks)-1]
}

func main() {
	first, last := solve("day04/test_input.txt")
	expectedFirst := 4512
	if expectedFirst != first[1] {
		fmt.Printf("expected %d, got %d\n", expectedFirst, first[1])
		return
	} else {
		fmt.Println("test passed")
	}

	first, last = solve("day04/input.txt")
	fmt.Printf("First: turn %d, total: %d\n", first[0], first[1])
	fmt.Printf("Last: turn %d, total: %d\n", last[0], last[1])
}
