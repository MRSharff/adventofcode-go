package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

var in = func() string {
	b, _ := ioutil.ReadFile("day15/input.txt")
	return string(b)
}()

var testInput = `1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`

var smallInput = `111
441
441`

func main() {
	// * motion restricted to 2 dimensions
	// * cavern is a square (or at least resembles one)
	// * The starting position is never entered, so its risk is not counted.

	// I think this is a search problem with edge weights being the risk levels
	// Dijkstras
	fmt.Println(part1(smallInput))

	if part1(testInput) != 40 {
		panic("Test failed")
	}

	fmt.Println(part1(in))
}

type point struct {
	x, y int
}

type node struct {
	p         point
	riskLevel int
}

type queue struct {
	q           chan node
	length      int
	containsMap map[node]bool
}

func (q *queue) add(n node) {
	// TODO: probably need to make sure we don't add to a full queue
	q.q <- n
	q.length++
	if q.containsMap == nil {
		q.containsMap = make(map[node]bool)
	}
	q.containsMap[n] = true
}

func (q *queue) remove() node {
	if q.isEmpty() {
		panic("remove called on empty queue")
	}
	n := <-q.q
	q.length--
	delete(q.containsMap, n)
	return n
}

func (q *queue) isEmpty() bool {
	return q.length == 0
}

func (q *queue) Len() int {
	return q.length
}

func (q *queue) contains(u node) bool {
	return q.containsMap[u]
}

type graph struct {
	nodes         map[point]node
	neighborCache map[node][]node
}

func (g *graph) addNode(n node) {
	if g.nodes == nil {
		g.nodes = make(map[point]node)
	}
	g.nodes[n.p] = n
}

func (g *graph) neighbors(u node) []node {
	var ns []node
	var isCached bool
	if g.neighborCache == nil {
		g.neighborCache = make(map[node][]node)
	}
	ns, isCached = g.neighborCache[u]
	if isCached {
		return ns
	}
	p := u.p
	topPoint := point{p.x, p.y - 1}
	rightPoint := point{p.x + 1, p.y}
	bottomPoint := point{p.x, p.y + 1}
	leftPoint := point{p.x - 1, p.y}
	if top, hasTop := g.nodes[topPoint]; hasTop {
		ns = append(ns, top)
	}
	if right, hasRight := g.nodes[rightPoint]; hasRight {
		ns = append(ns, right)
	}
	if bottom, hasBottom := g.nodes[bottomPoint]; hasBottom {
		ns = append(ns, bottom)
	}
	if left, hasLeft := g.nodes[leftPoint]; hasLeft {
		ns = append(ns, left)
	}
	g.neighborCache[u] = ns
	return ns
}

func part1(in string) int {
	g := graph{}
	lines := strings.Split(in, "\n")
	for y, line := range lines {
		for x, c := range line {
			riskLevel, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}
			g.addNode(node{point{x, y}, riskLevel})
		}
	}
	endPoint := point{len(lines[0]) - 1, len(lines) - 1}
	start := g.nodes[point{0, 0}]
	end := g.nodes[endPoint]
	return dijkstras(g, start, end)
}

func dijkstras(g graph, src node, dst node) int {
	distances := make(map[node]int)
	q := queue{q: make(chan node, len(g.nodes))}
	visited := make(map[node]bool)
	distances[src] = 0
	for _, v := range g.nodes {
		if v != src {
			distances[v] = math.MaxInt64
		}
		q.add(v)
	}

	for !q.isEmpty() {
		v := removeMinDistanceNodeNotAlreadyVisited(&q, distances, visited)
		visited[v] = true
		var neighbors []node
		neighbors = g.neighbors(v)
		for _, u := range neighbors {
			if visited[u] {
				continue
			}
			// the distance to any point will always be its risk level, so edges don't *really* matter
			d := distances[v] + u.riskLevel
			if d < distances[u] {
				distances[u] = d
			}
		}
	}
	return distances[dst]
}

func removeMinDistanceNodeNotAlreadyVisited(q *queue, distances map[node]int, visited map[node]bool) node {
	l := q.Len()
	var v node
	found := false
	min := math.MaxInt64
	for i := 0; i < l; i++ {
		n := q.remove()
		if !visited[n] && distances[n] < min {
			if found {
				q.add(v)
			}
			v = n
			min = distances[n]
			found = true
		} else {
			q.add(n)
		}
	}
	return v
}
