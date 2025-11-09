package day_07

import (
	_ "embed"
	"math"
	"slices"
	"strconv"
	"strings"
)

//go:embed input
var Input string

type Position int
type Positions []Position

func (position Position) align(target int) (cost int) {
	cost = abs(int(position) - target)
	return
}

func (positions Positions) unwrap() (result []int) {
	for _, position := range positions {
		result = append(result, int(position))
	}
	return
}

func (positions Positions) getMedian() (median int) {
	pos := positions.unwrap()
	slices.Sort(pos)
	if len(pos)%2 == 0 {
		median = (pos[len(pos)/2] + pos[len(pos)/2-1]) / 2
	} else {
		median = pos[len(pos)/2]
	}
	return
}

func (positions Positions) findCheapestAlignment() (cheapestAlignment int) {
	minPos, maxPos := slices.Min(positions), slices.Max(positions)
	lowestCost := math.Inf(1)
	for align := int(minPos); align < int(maxPos)+1; align += 1 {
		cost := float64(positions.getIncrementalCost(align))
		if cost < lowestCost {
			lowestCost = cost
			cheapestAlignment = align
		}
	}
	return
}

func (positions Positions) getCost(target int) (result int) {
	for position := range positions {
		cost := positions[position].align(target)
		result += cost
	}
	return
}

func (positions Positions) getIncrementalCost(target int) (result int) {
	for position := range positions {
		cost := positions[position].align(target)
		result += (cost * (cost + 1)) / 2
	}
	return
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func parseInput(input string) (positions Positions) {
	for position := range strings.SplitSeq(input, ",") {
		num, err := strconv.Atoi(position)
		if err != nil {
			panic(err)
		}
		positions = append(positions, Position(num))
	}
	return
}

func SolvePartOne(input string) (result int) {
	positions := parseInput(input)
	result = positions.getCost(positions.getMedian())
	return
}

func SolvePartTwo(input string) (result int) {
	positions := parseInput(input)
	cheapest := positions.findCheapestAlignment()
	result = positions.getIncrementalCost(cheapest)
	return
}
