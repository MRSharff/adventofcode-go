package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const testInput = `3,4,3,1,2`

func main() {

	// count the occurrences of each number, then
	// run the algorithm for the set of numbers that appear, and then
	// multiply the outcome for each number by how often it occurs
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

		results := make(map[int]int, len(agesCount))
		for age := range agesCount {
			// start with a population of 1 of the initial fish
			fish := []int{age}
			for d := 0; d < 80; d++ {
				for i := range fish {
					fish[i]--
				}
				var newFishOfTheDay []int
				for i, a := range fish {
					if a == -1 {
						newFishOfTheDay = append(newFishOfTheDay, 8)
						fish[i] = 6
					}
				}
				fish = append(fish, newFishOfTheDay...)
				//fmt.Printf("After %d day: %v\n", d+1, ages)
			}
			results[age] = len(fish)
		}
		total := 0
		for age, fishCount := range results {
			fmt.Printf("age %d produces %d fish\n", age, fishCount)
			totalForAge := fishCount * agesCount[age]
			total += totalForAge
		}
		return total
	}
	fmt.Println(solve(testInput))

	b, _ := ioutil.ReadFile("day6/input.txt")
	fmt.Println(solve(string(b)))
}
