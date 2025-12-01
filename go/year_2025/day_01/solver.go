package day_01

import (
	_ "embed"
	"iter"
	"strconv"
	"strings"
)

//go:embed input
var Input string

func ParseInput(input string) iter.Seq[int] {
	return func(yield func(int) bool) {
		for line := range strings.SplitSeq(input, "\n") {
			direction := line[0]
			rotation, err := strconv.Atoi(line[1:])
			if err != nil {
				panic(err)
			}
			switch direction {
			case 'L':
				rotation = -rotation
			case 'R':
			default:
				panic("failed to parse direction")
			}
			if yield(rotation) == false {
				return
			}
		}
	}
}

func SolvePartOne(input string) (result int) {
	dial := 50
	for rotation := range ParseInput(input) {
		dial = (100 + dial + rotation) % 100
		if dial == 0 {
			result++
		}
	}
	return
}

func SolvePartTwo(input string) (result int) {
	dial := 50
	for rotation := range ParseInput(input) {
		if rotation < 0 {
			result += rotation / 100 * -1
		} else {
			result += rotation / 100
		}
		nextDial := dial + rotation%100
		if nextDial <= 0 {
			if nextDial < 0 {
				nextDial += 100
			}
			if dial != 0 {
				result++
			}
		} else if nextDial >= 100 {
			nextDial -= 100
			result++
		}
		dial = nextDial
	}
	return
}
