package day_04

import (
	_ "embed"
	"strings"
)

//go:embed input
var Input string

type Grid[T any] [][]T
type Point struct {
	X int
	Y int
}

func checkIsInbound(predicate int, length int) bool {
	return predicate >= 0 && predicate < length
}

func (grid Grid[T]) CheckIsInbound(predicate Point) bool {
	return checkIsInbound(predicate.X, len(grid[0])) && checkIsInbound(predicate.Y, len(grid))
}

func (grid Grid[T]) GetNeighborIndices(start Point) (indices []Point) {
	if !grid.CheckIsInbound(start) {
		panic("start point is out of bounds")
	}
	for i := range 2 {
		offset := i/2 + 1
		sign := 1 - 2*(i%2)
		yNext := start.Y + (offset * sign)
		if checkIsInbound(yNext, len(grid)) {
			indices = append(indices, Point{start.X, yNext})
		}
		xNext := start.X + (offset * sign)
		if checkIsInbound(xNext, len(grid[0])) {
			indices = append(indices, Point{xNext, start.Y})
		}
	}
	return
}

func (grid Grid[T]) GetNeighborIndicesWithCorners(start Point) (indices []Point) {
	if !grid.CheckIsInbound(start) {
		panic("start point is out of bounds")
	}
	for yOffset := range 3 {
		for xOffset := range 3 {
			if xOffset == 1 && yOffset == 1 {
				continue
			}
			x := start.X - 1 + xOffset
			y := start.Y - 1 + yOffset
			if !checkIsInbound(x, len(grid[0])) || !checkIsInbound(y, len(grid)) {
				continue
			}
			indices = append(indices, Point{x, y})
		}
	}
	return
}

// external code above

func (grid Grid[rune]) CheckIsForkliftable(rollPoint Point) bool {
	var rollCount int
	for _, point := range grid.GetNeighborIndicesWithCorners(rollPoint) {
		cell := any(grid[point.Y][point.X]).(byte)
		if cell == '@' {
			rollCount += 1
		}
		if rollCount >= 4 {
			return false
		}
	}
	return true
}

func ParseInput(input *string) (grid Grid[byte]) {
	for line := range strings.SplitSeq(*input, "\n") {
		cells := strings.Split(line, "")
		row := make([]byte, len(line))
		for _, cell := range cells {
			row = append(row, byte(cell[0]))
		}
		grid = append(grid, row)
	}
	return
}

func SolvePartOne(input string) (result int) {
	grid := ParseInput(&input)

	for y, row := range grid {
		for x := range row {
			if grid[y][x] == '@' && grid.CheckIsForkliftable(Point{x, y}) {
				result += 1
			}
		}
	}
	return
}

func SolvePartTwo(input string) (result int) {
	grid := ParseInput(&input)

	couldForkLift := true
	for couldForkLift {
		couldForkLift = false
		forkLiftablePositions := []Point{}
		for y, row := range grid {
			for x := range row {
				if grid[y][x] == '@' && grid.CheckIsForkliftable(Point{x, y}) {
					forkLiftablePositions = append(forkLiftablePositions, Point{x, y})
					couldForkLift = true
					result += 1
				}
			}
		}
		for _, pos := range forkLiftablePositions {
			grid[pos.Y][pos.X] = 'x'
		}
	}

	return
}
