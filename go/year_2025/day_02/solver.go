package day_02

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input
var Input string

type IDRange struct {
	Left  int
	Right int
}

func (r IDRange) GetIdsWithDuplicateDigits() (result []int) {
	for i := r.Left; i <= r.Right; i++ {
		id := strconv.Itoa(i)
		if id[:len(id)/2] == id[len(id)/2:] {
			result = append(result, i)
		}
	}
	return
}

// func GetOverlappedSlidingWindow(field string, windowSize int) []string {
// 	var result []string
// 	for i := 0; i+windowSize <= len(field); i += 1 {
// 		result = append(result, field[i:i+windowSize])
// 	}
// 	return result
// }

func GetSlidingWindow(field string, windowSize int) []string {
	var result []string
	for i := 0; i+windowSize <= len(field); i += windowSize {
		result = append(result, field[i:i+windowSize])
	}
	return result
}

func (r IDRange) GetIdsWithRepeatedDigits() (result []int) {
	for i := r.Left; i <= r.Right; i++ {
		id := strconv.Itoa(i)
	nextWindow:
		for windowSize := len(id); windowSize > 0; windowSize-- {
			if len(id)%windowSize != 0 || len(id)/windowSize <= 1 {
				continue
			}
			var prev string
			for _, window := range GetSlidingWindow(id, windowSize) {
				if prev != "" && window != prev {
					continue nextWindow
				}
				prev = window
			}
			result = append(result, i)
			break
		}
	}
	return
}

func ParseInput(input *string) (result []IDRange) {
	for idRange := range strings.SplitSeq(*input, ",") {
		var left, right int
		fmt.Sscanf(idRange, "%d-%d", &left, &right)
		result = append(result, IDRange{Left: left, Right: right})
	}
	return
}

func SolvePartOne(input string) (result int) {
	idRanges := ParseInput(&input)
	for _, r := range idRanges {
		for _, invalidId := range r.GetIdsWithDuplicateDigits() {
			result += invalidId
		}
	}
	return
}

func SolvePartTwo(input string) (result int) {
	idRanges := ParseInput(&input)
	for _, idRange := range idRanges {
		for _, invalidId := range idRange.GetIdsWithRepeatedDigits() {
			result += invalidId
		}
	}
	return
}
