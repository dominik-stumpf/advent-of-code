package day_05

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input
var Input string

type Point struct {
	X, Y int
}

type Line struct {
	Start, End Point
}

func (line Line) isDiagonal() bool {
	return line.Start.X != line.End.X && line.Start.Y != line.End.Y
}

type Diagram map[Point]int

func parseInput(input string) (result []Line) {
	for line := range strings.SplitSeq(input, "\n") {
		var x1, y1, x2, y2 int
		fmt.Sscanf(line, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		result = append(result, Line{Start: Point{X: x1, Y: y1}, End: Point{X: x2, Y: y2}})
	}
	return
}

func (diagram Diagram) addLine(line Line) {
	x1, y1 := line.Start.X, line.Start.Y
	x2, y2 := line.End.X, line.End.Y

	xStep := 0
	if x2 > x1 {
		xStep = 1
	} else if x2 < x1 {
		xStep = -1
	}

	yStep := 0
	if y2 > y1 {
		yStep = 1
	} else if y2 < y1 {
		yStep = -1
	}

	iterationCount := int(max(abs(x2-x1), abs(y2-y1))) + 1

	x, y := x1, y1
	for range iterationCount {
		point := Point{X: x, Y: y}
		diagram[point]++
		x += xStep
		y += yStep
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func (diagram Diagram) countOverlaps() (count int) {
	for _, overlapCount := range diagram {
		if overlapCount >= 2 {
			count++
		}
	}
	return
}

func SolvePartOne(input string) int {
	lines := parseInput(input)
	diagram := make(Diagram)
	for _, line := range lines {
		if line.isDiagonal() {
			continue
		}
		diagram.addLine(line)
	}

	return diagram.countOverlaps()
}

func SolvePartTwo(input string) int {
	lines := parseInput(input)
	diagram := make(Diagram)
	for _, line := range lines {
		diagram.addLine(line)
	}

	return diagram.countOverlaps()
}
