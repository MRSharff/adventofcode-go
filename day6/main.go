package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const testInput = `3,4,3,1,2`

var maxDays = 80

func main() {
	solve := func(in string) int {

		agesStrings := strings.Split(in, ",")
		agesCount := make(map[int]int, 0)
		for _, a := range agesStrings {
			n, err := strconv.Atoi(a)
			if err != nil {
				panic(err)
			}
			agesCount[n]++
		}

		// respawningFish holds the fish that have 6 days or less on their spawn timer
		var respawningFish = make(chan int, 7)

		// newFish hold fish with 7 and 8 days left, such as those fish that are brand new
		var newFish = make(chan int, 2)

		// set up the queues with the initial state
		for i := 0; i <= 6; i++ {
			respawningFish <- agesCount[i]
		}
		for i := 7; i <= 8; i++ {
			newFish <- 0
		}

		for d := 0; d < maxDays; d++ {
			// respawning fish create n amount of new fish, and restart their spawn journey
			n, transitioning := <-respawningFish, <-newFish
			// add 7 day fish and the restarting fish to the back of the respawningFish queue
			respawningFish <- n + transitioning
			newFish <- n
		}
		close(respawningFish)
		close(newFish)

		sum := 0
		for n := range respawningFish {
			sum += n
		}
		for n := range newFish {
			sum += n
		}
		return sum
	}

	// tests
	maxDays = 80
	got := solve(testInput)
	expected := 5934
	if expected != got {
		fmt.Printf("expected %d, got %d\n", expected, got)
	} else {
		fmt.Println("80 day test passed")
	}

	maxDays = 256
	got = solve(testInput)
	expected = 26984457539
	if expected != got {
		fmt.Printf("expected %d, got %d\n", expected, got)
	} else {
		fmt.Println("256 day test passed")
	}

	// solve
	b, _ := ioutil.ReadFile("day6/input.txt")
	maxDays = 80
	fmt.Println(solve(string(b)))
	maxDays = 256
	fmt.Println(solve(string(b)))
}
