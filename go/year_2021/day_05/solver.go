package day_05

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed example
var Input string

type Point [2]int
type Line [2]Point
type Diagram []Line

func (diagram Diagram) plot() {
	maxX, maxY := 0, 0
	for _, line := range diagram {
		for _, point := range line {
			if point[0] > maxX {
				maxX = point[0]
			}
			if point[1] > maxY {
				maxY = point[1]
			}
		}
	}

	grid := make([][]rune, maxY+1)
	for i := range grid {
		grid[i] = make([]rune, maxX+1)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	for _, line := range diagram {
		if !line.isDiagonal() {
			x1, y1 := line[0][0], line[0][1]
			x2, y2 := line[1][0], line[1][1]

			if x1 == x2 {
				for y := min(y1, y2); y <= max(y1, y2); y++ {
					grid[y][x1] = '#'
				}
			} else if y1 == y2 {
				for x := min(x1, x2); x <= max(x1, x2); x++ {
					grid[y1][x] = '#'
				}
			}
		} else {
			x1, y1 := line[0][0], line[0][1]
			x2, y2 := line[1][0], line[1][1]

			if x1 < x2 && y1 < y2 {
				for i := 0; i <= abs(x2-x1); i++ {
					grid[y1+i][x1+i] = '#'
				}
			} else if x1 > x2 && y1 < y2 {
				for i := 0; i <= abs(x2-x1); i++ {
					grid[y1+i][x1-i] = '#'
				}
			} else if x1 < x2 && y1 > y2 {
				for i := 0; i <= abs(x2-x1); i++ {
					grid[y1-i][x1+i] = '#'
				}
			} else if x1 > x2 && y1 > y2 {
				for i := 0; i <= abs(x2-x1); i++ {
					grid[y1-i][x1-i] = '#'
				}
			}
		}
	}

	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func (diagram Diagram) print() {
	for _, line := range diagram {
		line.print()
		fmt.Println()
	}
}

func (line Line) print() {
	fmt.Printf("%2d %2d", line[0], line[1])
}

func (line Line) isDiagonal() bool {
	return line[0][0] != line[1][0] && line[0][1] != line[1][1]
}

func (firstLine Line) checkOverlap(secondLine Line) (result []Point) {
	a := firstLine[0]
	b := firstLine[1]
	c := secondLine[0]
	d := secondLine[1]

	denominator := float64((b[0]-a[0])*(d[1]-c[1]) - (b[1]-a[1])*(d[0]-c[0]))
	numerator1 := float64((a[1]-c[1])*(d[0]-c[0]) - (a[0]-c[0])*(d[1]-c[1]))
	numerator2 := float64((a[1]-c[1])*(b[0]-a[0]) - (a[0]-c[0])*(b[1]-a[1]))

	if denominator == 0 && numerator1 == 0 && numerator2 == 0 {
		minX1, maxX1 := min(a[0], b[0]), max(a[0], b[0])
		minY1, maxY1 := min(a[1], b[1]), max(a[1], b[1])
		minX2, maxX2 := min(c[0], d[0]), max(c[0], d[0])
		minY2, maxY2 := min(c[1], d[1]), max(c[1], d[1])

		if max(minX1, minX2) <= min(maxX1, maxX2) && max(minY1, minY2) <= min(maxY1, maxY2) {
			startX := max(minX1, minX2)
			endX := min(maxX1, maxX2)
			startY := max(minY1, minY2)
			endY := min(maxY1, maxY2)

			if a[0] == b[0] {
				for y := startY; y <= endY; y++ {
					result = append(result, Point{a[0], y})
				}
			} else {
				for x := startX; x <= endX; x++ {
					result = append(result, Point{x, a[1]})
				}
			}
		}
		return result
	}

	if denominator != 0 {
		r := numerator1 / denominator
		s := numerator2 / denominator
		if r >= 0 && r <= 1 && s >= 0 && s <= 1 {
			ix := a[0] + int(r*float64(b[0]-a[0]))
			iy := a[1] + int(r*float64(b[1]-a[1]))
			result = append(result, Point{ix, iy})
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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

func parseDiagram(diagram string) (result Diagram) {
	lines := strings.SplitSeq(diagram, "\n")

	for line := range lines {
		var diagramLine Line
		for i, point := range strings.Split(line, " -> ") {
			parts := strings.Split(point, ",")
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])
			diagramLine[i] = Point{x, y}
		}
		result = append(result, diagramLine)
	}
	return
}

func Solve(input string) (result int) {
	fullDiagram := parseDiagram(input)
	var diagram Diagram
	for _, line := range fullDiagram {
		if !line.isDiagonal() {
			diagram = append(diagram, line)
		}
	}
	// diagram.print()

	var overlappingPoints []Point
	for i, firstLine := range diagram {
		for j := i; j < len(diagram); j += 1 {
			secondLine := diagram[j]
			if firstLine == secondLine {
				continue
			}
			firstLine.print()
			fmt.Print(" - ")
			secondLine.print()
			fmt.Printf(" %d\n", len(overlappingPoints))
			for _, point := range firstLine.checkOverlap(secondLine) {
				overlappingPoints = append(overlappingPoints, point)
				// if !slices.Contains(overlappingPoints, point) {
				// 	overlappingPoints = append(overlappingPoints, point)
				// }
			}
		}
	}

	diagram.plot()
	for _, line := range diagram {
		line.print()
		fmt.Println()
	}

	fmt.Println(overlappingPoints)
	result = len(overlappingPoints)

	return
	// less than 7067
}
