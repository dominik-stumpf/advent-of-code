package day_06

import (
	_ "embed"
	"fmt"
	"runtime"
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

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", m.Alloc/1024/1024)
	fmt.Printf("\tSys = %v MiB", m.Sys/1024/1024)
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func SolvePartOne(input string) (result int) {
	fishes := parseInput(input)
	cache := make(map[int]int)
	for _, fish := range fishes {
		if cache[fish] > 0 {
			result += cache[fish]
			continue
		}
		partial := []uint8{uint8(fish)}
		for range 80 {
			for i, fish := range partial {
				if fish > 0 {
					partial[i] -= 1
				}
				if fish == 0 {
					partial[i] = 6
					partial = append(partial, 8)
				}
			}
		}
		cache[fish] = len(partial)
		result += len(partial)
	}

	return
}

func SolvePartTwo(input string) (result int) {
	fishes := parseInput(input)
	cache := make(map[int]int)
	for _, fish := range fishes {
		if cache[fish] > 0 {
			result += cache[fish]
			continue
		}
		partial := []uint8{uint8(fish)}
		for range 256 {
			PrintMemUsage()
			for i, fish := range partial {
				if fish > 0 {
					partial[i] -= 1
				}
				if fish == 0 {
					partial[i] = 6
					partial = append(partial, 8)
				}
			}
		}
		cache[fish] = len(partial)
		result += len(partial)
	}

	return
}
