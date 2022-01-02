package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

const testInput = `start-A
start-b
A-c
A-b
b-d
A-end
b-end`

var in = func() string {
	b, _ := ioutil.ReadFile("day12/input.txt")
	return string(b)
}()

var paths [][]string
var adjacencyLists map[string][]string
var caveAllowed2Visits = ""
var pathSeen map[string]bool

func main() {
	part1(testInput)
	part1(in)
	testPart1()
	part2(testInput)
	testPart2()
	part2(in)
	testPart2()
}

func part1(in string) {
	buildAdjacencyList(in)
	printAdjacencyList()
	paths = make([][]string, 0)
	pathSeen = make(map[string]bool)
	dfs(nil, "start")
	printPaths()
}

func part2(in string) {
	caveAllowed2Visits = ""
	buildAdjacencyList(in)
	printAdjacencyList()
	paths = make([][]string, 0)
	pathSeen = make(map[string]bool)

	var smallCaves []string
	for cave := range adjacencyLists {
		isSmallCave := 'a' <= cave[0] && cave[0] <= 'z'
		if isSmallCave && cave != "start" && cave != "end" {
			smallCaves = append(smallCaves, cave)
		}
	}
	sort.Strings(smallCaves)

	for _, cave := range smallCaves {
		caveAllowed2Visits = cave
		dfs(nil, "start")
	}
	printPaths()
}

func testPart1() {
	for _, path := range paths {
		if !isDistinct(path) {
			panic(fmt.Sprintln("Path not distinct:", path))
		}

		if !startToEnd(path) {
			panic(fmt.Sprintln("Path not start to end:", path))
		}

		if !visitsSmallCaveAtMostOnce(path) {
			panic(fmt.Sprintln("Small cave visited twice: ", path))
		}
	}
	fmt.Println("Part 1 Tests passed")
}

func testPart2() {
	pathCounts = nil
	for _, path := range paths {
		if !isDistinct(path) {
			panic(fmt.Sprintln("Path not distinct:", path))
		}

		if !startToEnd(path) {
			panic(fmt.Sprintln("Path not start to end:", path))
		}

		if !visitsSingleSmallCaveAtMostTwice(path) {
			panic(fmt.Sprintln("More than one small cave was visited twice: ", path))
		}
	}
	fmt.Println("Part 2 Tests passed")
}

func visitsSingleSmallCaveAtMostTwice(path []string) bool {
	smallCaveVisitCount := make(map[string]int)
	seenTwice := false
	for _, cave := range path {
		if isBigCave(cave) {
			continue
		}
		smallCaveVisitCount[cave]++
		if smallCaveVisitCount[cave] == 2 {
			if !seenTwice {
				seenTwice = true
			} else {
				return false
			}
		}
	}
	return true
}

func visitsSmallCaveAtMostOnce(path []string) bool {
	alreadyVisited := make(map[string]bool)
	for _, cave := range path {
		if alreadyVisited[cave] {
			if !isBigCave(cave) {
				return false
			}
		}
	}
	return true
}

var pathCounts map[string]int

func isDistinct(path []string) bool {
	if pathCounts == nil {
		pathCounts = make(map[string]int)
		for _, p := range paths {
			s := strings.Join(p, "")
			pathCounts[s]++
		}
	}
	return pathCounts[strings.Join(path, "")] == 1
}

func startToEnd(path []string) bool {
	// I actually should make sure that a path with a start or end in the middle should fail this
	return path[0] == "start" && path[len(path)-1] == "end"
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
	for _, list := range adjacencyLists {
		sort.Strings(list)
	}
}

func dfs(visitedPath []string, cave string) {
	// It seems like we can remove the slice copying, but this seems more safe.
	// I need to look into why this is.
	path := copyPath(visitedPath)
	path = append(path, cave)
	if cave == "end" {
		pString := strings.Join(path, "")
		if pathSeen[pString] {
			// This is a hack to not add the path more than once.
			// I don't really like it, but it allowed me to solve the problem
			return
		}
		paths = append(paths, path)
		pathSeen[strings.Join(path, "")] = true
		return
	}
	neighbors := adjacencyLists[cave]
	for _, neighbor := range neighbors {
		if !visited(path, neighbor) || isBigCave(neighbor) || (neighbor == caveAllowed2Visits && visitCount(path, neighbor) < 2) {
			dfs(path, neighbor)
		}
	}
}

func copyPath(pathToCopy []string) []string {
	path := make([]string, len(pathToCopy))
	copy(path, pathToCopy)
	return path
}

func visitCount(path []string, cave string) int {
	count := 0
	for _, c := range path {
		if cave == c {
			count++
		}
	}
	return count
}

func isBigCave(cave string) bool {
	return 'A' <= cave[0] && cave[0] <= 'Z'
}

func visited(path []string, cave string) bool {
	for _, c := range path {
		if c == cave {
			return true
		}
	}
	return false
}
