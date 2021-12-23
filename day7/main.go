package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

const testInput = `16,1,2,0,4,2,7,1,2,14`

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return n * -1
}

func naive(in string) {
	var xs []int
	for _, s := range strings.Split(in, ",") {
		x, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		xs = append(xs, x)
	}

	sort.Slice(xs, func(i, j int) bool {
		return xs[i] < xs[j]
	})

	fuelCosts := make(map[int]int)
	min, max := xs[0], xs[len(xs)-1]
	for i := min; i < max+1; i++ {
		for _, x := range xs {
			cost := abs(x - i)
			cost = (cost * (cost + 1)) / 2
			fuelCosts[i] += cost
		}
	}

	min = math.MaxInt64
	for _, cost := range fuelCosts {
		if cost < min {
			min = cost
		}
	}
	fmt.Println(min)
}

// Day 7: The Treachery of Whales
func main() {
	b, _ := ioutil.ReadFile("day7/input.txt")
	naive(testInput)
	naive(string(b))
}
