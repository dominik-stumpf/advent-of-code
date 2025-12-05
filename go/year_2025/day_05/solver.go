package day_05

import (
	_ "embed"
	"fmt"
	"maps"
	"math"
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
		result.IDRanges = append(result.IDRanges, ParseIDRangeFromKey(line))
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

func ParseIDRangeFromKey(key string) IDRange {
	var left, right int
	fmt.Sscanf(key, "%d-%d", &left, &right)
	return IDRange{left, right}
}

func (rang IDRange) CheckIsInRange(id int) bool {
	return rang.Left <= id && rang.Right >= id
}

func (rang IDRange) GetLength() int {
	return rang.Right - rang.Left + 1
}
func (rang IDRange) GetKey() string {
	return fmt.Sprintf("%d-%d", rang.Left, rang.Right)
}

func (ing *Ingredients) RemoveDuplicates() {
	rangMap := map[string]IDRange{}
	for _, rang := range ing.IDRanges {
		rangMap[rang.GetKey()] = rang
	}
	ing.IDRanges = slices.Collect(maps.Values(rangMap))
}

func (ing *Ingredients) ReduceIDRanges() (result []IDRange) {
	ing.RemoveDuplicates()
	slices.SortFunc(ing.IDRanges, func(a, b IDRange) int {
		return a.Left - b.Left
	})
	groups := map[string]int{}
	var groupId int
	for _, rang := range ing.IDRanges {
		couldMerge := false
		for _, rangEnd := range ing.IDRanges {
			if rang == rangEnd {
				continue
			}
			if rang.Left <= rangEnd.Right && rangEnd.Left <= rang.Right {
				rangGroupId, ok := groups[rang.GetKey()]
				rangEndGroupId, endOk := groups[rangEnd.GetKey()]
				if ok && !endOk {
					groups[rangEnd.GetKey()] = rangGroupId
				} else if endOk && !ok {
					groups[rang.GetKey()] = rangEndGroupId
				} else if !ok && !endOk {
					groups[rang.GetKey()] = groupId
					groups[rangEnd.GetKey()] = groupId
					groupId += 1
				}
				couldMerge = true
				continue
			}
		}
		if !couldMerge {
			result = append(result, rang)
		}
	}
	for i := range groupId {
		groupOverlap := IDRange{Left: math.MaxInt, Right: 0}
		for key, value := range groups {
			if value == i {
				rang := ParseIDRangeFromKey(key)
				if rang.Left < groupOverlap.Left {
					groupOverlap.Left = rang.Left
				}
				if rang.Right > groupOverlap.Right {
					groupOverlap.Right = rang.Right
				}
			}
		}
		result = append(result, groupOverlap)
	}
	return
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
	ingredients := MustParseInput(&input)
	for _, rang := range ingredients.ReduceIDRanges() {
		result += rang.GetLength()
	}
	return
}
