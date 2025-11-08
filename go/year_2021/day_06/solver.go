package day_06

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input
var Input string

func parseInput(input string) []int {
	lines := strings.SplitSeq(input, "\n")
	for line := range lines {
		var fishes []int
		for fish := range strings.SplitSeq(line, ",") {
			num, err := strconv.Atoi(fish)
			if err != nil {
				panic(err)
			}
			fishes = append(fishes, num)
		}

		return fishes
	}
	return []int{}
}

// func calculateSpawn(progress int, day int) int {
// 	growthTime := 7
// 	rate := 1.0 / float64(growthTime)
// 	delta := progress
// 	if progress > growthTime {
// 		delta = growthTime - progress - 1
// 	}

// 	spawn := int(math.Floor(rate*float64(day) + (rate * float64(delta))))
// 	var result = spawn
// 	for i := range spawn {
// 		remainingDay := day - (progress + (i)*growthTime)
// 		println(result, progress, day, remainingDay)
// 		if remainingDay >= 0 {
// 			result += calculateSpawn(8, remainingDay)
// 		}
// 	}
// 	return int(math.Max(float64(result), 1))
// }

func SolvePartOne(input string) (result int) {
	fishes := parseInput(input)
	for range 80 {
		for i, fish := range fishes {
			if fish > 0 {
				fishes[i] -= 1
			}
			if fish == 0 {
				fishes[i] = 6
				fishes = append(fishes, 8)
			}
		}
	}
	result = len(fishes)

	// first := fishes[0]
	// spawn := calculateSpawn(first, 12)
	// fmt.Println(spawn)

	// for _, fish := range fishes {
	// 	result += calculateSpawn(fish, 18)
	// }

	return
}

func SolvePartTwo(input string) (result int) {
	return
}
