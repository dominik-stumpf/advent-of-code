package day_03

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input
var Input string

type Bank []byte

func ParseInput(input *string) (result []Bank) {
	for line := range strings.SplitSeq(*input, "\n") {
		result = append(result, Bank(line))
	}
	return
}

func (b Bank) FindLargestJoltage() int {
	max := make([]byte, 2)
	for i, battery := range b {
		if battery > max[0] {
			if i != len(b)-1 {
				max[0] = battery
				max[1] = 0
			} else {
				max[1] = battery
			}
		} else if battery > max[1] {
			max[1] = battery
		}
	}
	largestJoltage, err := strconv.Atoi(string(max))
	if err != nil {
		panic(err)
	}
	return largestJoltage
}

func SolvePartOne(input string) (result int) {
	banks := ParseInput(&input)
	for _, bank := range banks {
		result += bank.FindLargestJoltage()
	}
	return
}

func SolvePartTwo(input string) (result int) {
	return
}
