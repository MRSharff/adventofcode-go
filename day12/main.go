package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const testInput = `start-A
start-b
A-c
A-b
b-d
A-end
b-end`

var paths [][]string
var adjacencyLists map[string][]string

func main() {
	part1(testInput)
	b, _ := ioutil.ReadFile("day12/input.txt")
	part1(string(b))
}

func part1(in string) {
	buildAdjacencyList(in)
	printAdjacencyList()
	paths = make([][]string, 0)
	var path []string
	dfs(path, "start")
	printPaths()
}

func printPaths() {
	for _, path := range paths {
		fmt.Println(strings.Join(path, ","))
	}
	fmt.Println("Total paths", len(paths))
	fmt.Println()
}

func printAdjacencyList() {
	for k, v := range adjacencyLists {
		fmt.Println(k, ":", v)
	}
	fmt.Println()
}

func canVisitMultipleTimes(cave string) bool {
	if len(cave) == 0 {
		panic("caves must be represented with at least one letter")
	}
	isBigCave := 'A' <= cave[0] && cave[0] <= 'Z'
	return cave != "start" && isBigCave
}

func buildAdjacencyList(input string) {
	adjacencyLists = make(map[string][]string)
	for _, edge := range strings.Split(input, "\n") {
		nodes := strings.Split(edge, "-")
		for _, n := range nodes {
			_, nodeExists := adjacencyLists[n]
			if !nodeExists {
				adjacencyLists[n] = make([]string, 0)
			}
		}
		u, v := nodes[0], nodes[1]
		adjacencyLists[u] = append(adjacencyLists[u], v)
		adjacencyLists[v] = append(adjacencyLists[v], u)
	}
}

func dfs(visitedPath []string, start string) {
	// It seems like we can remove the slice copying, but this seems more safe.
	// I need to look into why this is.
	path := make([]string, len(visitedPath))
	copy(path, visitedPath)
	path = append(path, start)
	if start == "end" {
		paths = append(paths, path)
		return
	}
	options := adjacencyLists[start]
	for _, cave := range options {
		if !visited(path, cave) || canVisitMultipleTimes(cave) {
			dfs(path, cave)
		}
	}
}

func visited(path []string, cave string) bool {
	for _, c := range path {
		if c == cave {
			return true
		}
	}
	return false
}
