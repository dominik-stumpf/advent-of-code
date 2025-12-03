package day_02

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input
var Input string

type Range struct {
	Left  int
	Right int
}

func (r Range) getInvalidIds() (result []int) {
	for i := r.Left; i <= r.Right; i++ {
		id := strconv.Itoa(i)
		if id[:len(id)/2] == id[len(id)/2:] {
			result = append(result, i)
		}
	}
	return
}

func ParseInput(input *string) (result []Range) {
	for idRange := range strings.SplitSeq(*input, ",") {
		var left, right int
		fmt.Sscanf(idRange, "%d-%d", &left, &right)
		result = append(result, Range{Left: left, Right: right})
	}
	return
}

func SolvePartOne(input string) (result int) {
	ranges := ParseInput(&input)
	for _, r := range ranges {
		for _, invalidId := range r.getInvalidIds() {
			result += invalidId
		}
	}
	return
}

func SolvePartTwo(input string) (result int) {
	return
}
