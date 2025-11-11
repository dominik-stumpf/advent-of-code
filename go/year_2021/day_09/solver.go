package day_xx

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed example
var Input string

type HeightMap [][]int

func checkIsInbound(predicate int, length int) bool {
	return predicate >= 0 && predicate < length
}

func (heightMap HeightMap) findPossibleLowPointIndices() [][]int {
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
	return possibleLowPoints
}

func (heightMap HeightMap) findLowPoints() []int {
	possibleLowPoints := heightMap.findPossibleLowPointIndices()

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

func (heightMap HeightMap) traverse(startX, startY int, positions [][]int) {
	if !checkIsInbound(startX, len(heightMap[0])) || !checkIsInbound(startY, len(heightMap)) {
		panic("start point is out of bounds")
	}

	traverser := func(x, y int) {
		height := heightMap[y][x]
		if height == 9 {
			return
		}
		for _, position := range positions {
			if position[0] == x && position[1] == y {
				return
			}
		}
		fmt.Println(height)
		positions := append(positions, []int{x, y})
		heightMap.traverse(x, y, positions)
	}

	for i := range 2 {
		offset := i/2 + 1
		sign := 1 - 2*(i%2)
		yNext := startY + (offset * sign)
		if checkIsInbound(yNext, len(heightMap)) {
			// traverser(startX, yNext)
			// fmt.Println(heightMap[yNext][startX])
		}
		xNext := startY + (offset * sign)
		if checkIsInbound(xNext, len(heightMap[0])) {
			traverser(xNext, startY)
			// fmt.Println(xNext, heightMap[startY][xNext])
		}
	}
}

func SolvePartOne(input string) (result int) {
	heightMap := parseInput(input)
	for _, height := range heightMap.findLowPoints() {
		result += height + 1
	}
	return
}

func SolvePartTwo(input string) (result int) {
	heightMap := parseInput(input)
	heightMap.traverse(3, 3, [][]int{})
	// lowPoints := heightMap.findPossibleLowPointIndices()
	// fmt.Println(lowPoints)
	result = 1
	return
}
