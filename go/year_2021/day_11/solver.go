package day_11

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input
var Input string

type Grid [][]int

func addAnsiEscape(sub string) string {
	highlight, reset := "\033[33;1m", "\033[0m"
	return highlight + sub + reset
}

func (grid Grid) print() {
	for y, row := range grid {
		for x, _ := range row {
			cell := strconv.Itoa(grid[y][x])
			if grid[y][x] == 0 {
				fmt.Print(addAnsiEscape(cell))
			} else {
				fmt.Print(cell)
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func checkIsInbound(predicate int, length int) bool {
	return predicate >= 0 && predicate < length
}

func (grid Grid) incrementAdjacent(xTarget, yTarget int) {
	if !checkIsInbound(xTarget, len(grid[0])) || !checkIsInbound(yTarget, len(grid)) {
		panic("target out of bound")
	}

	for iy := range 3 {
		for ix := range 3 {
			if ix == 1 && iy == 1 {
				continue
			}
			x := xTarget - 1 + ix
			y := yTarget - 1 + iy
			if !checkIsInbound(x, len(grid[0])) || !checkIsInbound(y, len(grid)) {
				continue
			}
			grid[y][x] += 1
			if grid[y][x] == 10 {
				grid.incrementAdjacent(x, y)
			}
		}
	}
}

func (grid Grid) step() int {
	var flashes int
	for y, row := range grid {
		for x, _ := range row {
			grid[y][x] += 1
			if grid[y][x] == 10 {
				grid.incrementAdjacent(x, y)
			}
		}
	}

	for y, row := range grid {
		for x, _ := range row {
			if grid[y][x] > 9 {
				flashes += 1
				grid[y][x] = 0
			}
		}
	}
	return flashes
}

func parseInput(input string) Grid {
	var result [][]int
	for line := range strings.SplitSeq(input, "\n") {
		row := make([]int, len(line))
		for i, item := range strings.Split(line, "") {
			num, _ := strconv.Atoi(item)
			row[i] = num
		}
		result = append(result, row)
	}
	return Grid(result)
}

func SolvePartOne(input string) (result int) {
	grid := parseInput(input)
	// grid.print()
	// grid.step()
	// fmt.Println()
	// grid.print()
	// grid.step()
	// fmt.Println()
	// grid.print()

	for range 100 {
		result += grid.step()
	}

	return
}

func SolvePartTwo(input string) (result int) {
	return
}
