package day_02

import (
	_ "embed"
	"fmt"
	"iter"
	"slices"
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

func GetSlidingWindow(field []byte, windowSize int) iter.Seq[[]byte] {
	return func(yield func([]byte) bool) {
		for i := 0; i+windowSize <= len(field); i += windowSize {
			if !yield(field[i : i+windowSize]) {
				return
			}
		}
	}
}

func (r IDRange) GetIdsWithRepeatedDigits() (result []int) {
	for i := r.Left; i <= r.Right; i++ {
		id := []byte(strconv.Itoa(i))

	nextWindow:
		for windowSize := len(id); windowSize > 0; windowSize-- {
			if len(id)%windowSize != 0 || len(id)/windowSize <= 1 {
				continue
			}
			var prev []byte
			for window := range GetSlidingWindow(id, windowSize) {
				if len(prev) != 0 && !slices.Equal(window, prev) {
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
