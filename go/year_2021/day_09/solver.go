package day_xx

import (
	_ "embed"
	"math"
	"strconv"
	"strings"
)

//go:embed input
var Input string

type HeightMap [][]int

func checkIsInbound(predicate int, length int) bool {
	return predicate >= 0 && predicate < length
}

func (heightMap HeightMap) findLowPoints() []int {
	possibleLowPoints := [][]int{}
	for _, row := range heightMap {
		prevHeight := math.MaxInt
		var lowPoint int
		var lowPoints []int
		for x, height := range row {
			if prevHeight > height {
				lowPoint = height
				if len(row)-1 == x {
					lowPoints = append(lowPoints, x)
				}
			} else {
				if lowPoint != -1 {
					lowPoints = append(lowPoints, x-1)
					lowPoint = -1
				}
			}
			prevHeight = height
		}

		possibleLowPoints = append(possibleLowPoints, lowPoints)
	}

	confirmedLowPoints := []int{}
	for y, lowPoints := range possibleLowPoints {
		for _, x := range lowPoints {
			lowPoint := heightMap[y][x]

			var isLowPoint bool
			for i := range 2 {
				yOffset := i/2 + 1
				sign := 1 - 2*(i%2)
				yNext := y + (yOffset * sign)
				if !checkIsInbound(yNext, len(heightMap)) {
					continue
				}
				predicate := heightMap[yNext][x]
				if predicate < lowPoint {
					isLowPoint = false
					break
				} else if predicate == lowPoint {
					continue
				} else {
					isLowPoint = true
				}
			}
			if isLowPoint {
				confirmedLowPoints = append(confirmedLowPoints, lowPoint)
			}
		}
	}

	return confirmedLowPoints
}

func parseInput(input string) (result HeightMap) {
	for line := range strings.SplitSeq(input, "\n") {
		row := []int{}
		for value := range strings.SplitSeq(line, "") {
			height, err := strconv.Atoi(value)
			if err != nil {
				panic(err)
			}
			row = append(row, height)
		}
		result = append(result, row)
	}
	return
}

func SolvePartOne(input string) (result int) {
	heightMap := parseInput(input)
	for _, height := range heightMap.findLowPoints() {
		result += height + 1
	}
	return
}

func SolvePartTwo(input string) (result int) {
	return
}
