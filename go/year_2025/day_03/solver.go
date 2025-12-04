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

func (bank Bank) FindLargestJoltageByLength(length int) int {
	digits := make([]byte, length)
	for i, battery := range bank {
		for j := min(max(length-(len(bank)-i-1)-1, 0), length-1); j < length; j++ {
			if battery > digits[j] {
				digits[j] = battery
				for l := j + 1; l < length; l++ {
					digits[l] = 0
				}
				break
			}
		}
	}
	largestJoltage, err := strconv.Atoi(string(digits))
	if err != nil {
		panic(err)
	}
	return largestJoltage
}

func SolvePartOne(input string) (result int) {
	banks := ParseInput(&input)
	for _, bank := range banks {
		result += bank.FindLargestJoltageByLength(2)
	}
	return
}

func SolvePartTwo(input string) (result int) {
	banks := ParseInput(&input)
	for _, bank := range banks {
		result += bank.FindLargestJoltageByLength(12)
	}
	return
}
