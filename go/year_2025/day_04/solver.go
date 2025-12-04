package day_04

import (
	gridtl "aoc/year_2021/standalone/gridtl"
	_ "embed"
	"strings"
)

//go:embed input
var Input string

type Diagram struct{ gridtl.Grid[byte] }

func (diagram Diagram) CheckIsForkliftable(rollPoint gridtl.Point) bool {
	var rollCount int
	for _, point := range diagram.GetNeighborIndicesWithCorners(rollPoint) {
		if diagram.Grid[point.Y][point.X] == '@' {
			rollCount += 1
		}
		if rollCount >= 4 {
			return false
		}
	}
	return true
}

func ParseInput(input *string) (diagram Diagram) {
	for line := range strings.SplitSeq(*input, "\n") {
		cells := strings.Split(line, "")
		row := make([]byte, len(line))
		for _, cell := range cells {
			row = append(row, byte(cell[0]))
		}
		diagram.Grid = append(diagram.Grid, row)
	}
	return
}

func SolvePartOne(input string) (result int) {
	diagram := ParseInput(&input)
	for y, row := range diagram.Grid {
		for x := range row {
			if diagram.Grid[y][x] == '@' && diagram.CheckIsForkliftable(gridtl.Point{X: x, Y: y}) {
				result += 1
			}
		}
	}
	return
}

func SolvePartTwo(input string) (result int) {
	diagram := ParseInput(&input)
	couldForkLift := true
	for couldForkLift {
		couldForkLift = false
		forkLiftablePositions := []gridtl.Point{}
		for y, row := range diagram.Grid {
			for x := range row {
				if diagram.Grid[y][x] == '@' && diagram.CheckIsForkliftable(gridtl.Point{X: x, Y: y}) {
					forkLiftablePositions = append(forkLiftablePositions, gridtl.Point{X: x, Y: y})
					couldForkLift = true
					result += 1
				}
			}
		}
		for _, pos := range forkLiftablePositions {
			diagram.Grid[pos.Y][pos.X] = 'x'
		}
	}
	return
}
