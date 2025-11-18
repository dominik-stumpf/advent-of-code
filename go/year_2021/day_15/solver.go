package day_15

import (
	"container/heap"
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed example
var Input string

func ParseInput(input string) (result [][]int) {
	for line := range strings.SplitSeq(input, "\n") {
		row := make([]int, len(line))
		for x, char := range strings.Split(line, "") {
			cell, err := strconv.Atoi(char)
			if err != nil {
				panic(err)
			}
			row[x] = cell
		}
		result = append(result, row)
	}
	return
}

type Point struct {
	X, Y int
}

type Node struct {
	Point
	Cost  int
	Index int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Cost < pq[j].Cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	node := x.(*Node)
	node.Index = n
	*pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil
	node.Index = -1
	*pq = old[0 : n-1]
	return node
}

func Dijkstra(grid [][]int, start, end Point) ([]Point, int) {
	yMax, xMax := len(grid), len(grid[0])
	if !inBounds(start, yMax, xMax) || !inBounds(end, yMax, xMax) {
		return nil, math.MaxInt
	}

	dist := make([][]int, yMax)
	parent := make([][]Point, yMax)
	for i := range dist {
		dist[i] = make([]int, xMax)
		parent[i] = make([]Point, xMax)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt
		}
	}

	dist[start.Y][start.X] = grid[start.Y][start.X]

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Node{Point: start, Cost: dist[start.Y][start.X]})

	directions := []Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(*Node)
		x, y := curr.X, curr.Y

		if curr.Cost != dist[y][x] {
			continue
		}

		if x == end.X && y == end.Y {
			path := []Point{}
			cur := end
			for cur != start {
				path = append(path, cur)
				cur = parent[cur.Y][cur.X]
			}
			path = append(path, start)

			for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
				path[i], path[j] = path[j], path[i]
			}
			return path, dist[y][x]
		}

		for _, d := range directions {
			nx, ny := x+d.X, y+d.Y
			if !inBounds(Point{nx, ny}, yMax, xMax) {
				continue
			}

			newCost := dist[y][x] + grid[ny][nx]
			if newCost < dist[ny][nx] {
				dist[ny][nx] = newCost
				parent[ny][nx] = Point{x, y}
				heap.Push(pq, &Node{Point: Point{nx, ny}, Cost: newCost})
			}
		}
	}

	return nil, math.MaxInt
}

func inBounds(p Point, yMax, xMax int) bool {
	return p.X >= 0 && p.X < xMax && p.Y >= 0 && p.Y < yMax
}

func printWithPath(grid [][]int, path []Point) {
	pathSet := make(map[Point]bool)
	for _, p := range path {
		pathSet[p] = true
	}

	for y := range grid {
		for x := 0; x < len(grid[y]); x++ {
			if pathSet[Point{x, y}] {
				fmt.Printf("[%d]", grid[y][x])
			} else {
				fmt.Printf(" %d ", grid[y][x])
			}
		}
		fmt.Println()
	}
}

func SolvePartOne(input string) (result int) {
	grid := ParseInput(input)

	start := Point{0, 0}
	end := Point{len(grid[0]) - 1, len(grid) - 1}

	_, totalCost := Dijkstra(grid, start, end)
	result = totalCost - grid[0][0]
	return
}

func SolvePartTwo(input string) (result int) {
	return
}
