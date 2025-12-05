package day_05

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

type Ingredients struct {
	IDRanges []IDRange
	IDs      []int
}

func MustParseInput(input *string) (result Ingredients) {
	ranges, ids, found := strings.Cut(*input, "\n\n")
	if !found {
		panic("failed to parse input")
	}
	for line := range strings.SplitSeq(ranges, "\n") {
		var left, right int
		fmt.Sscanf(line, "%d-%d", &left, &right)
		result.IDRanges = append(result.IDRanges, IDRange{left, right})
	}
	for line := range strings.SplitSeq(ids, "\n") {
		id, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		result.IDs = append(result.IDs, id)
	}
	return
}

func (rang IDRange) CheckIsInRange(id int) bool {
	return rang.Left <= id && rang.Right >= id
}

func SolvePartOne(input string) (result int) {
	ingredients := MustParseInput(&input)
	for _, id := range ingredients.IDs {
		isSpoiled := true
		for _, rang := range ingredients.IDRanges {
			if rang.CheckIsInRange(id) {
				isSpoiled = false
				break
			}
		}
		if !isSpoiled {
			result++
		}
	}
	return
}

func SolvePartTwo(input string) (result int) {
	return
}
