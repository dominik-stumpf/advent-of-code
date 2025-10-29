package day_03

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input
var Input string

// func countOnes(lines []string) []int {
// 	var width = len(lines[0])
// 	var ones = make([]int, width)
// 	for _, line := range lines {
// 		for i, bit := range strings.Split(line, "") {
// 			val, err := strconv.Atoi(bit)
// 			if err != nil {
// 				panic(err)
// 			}
// 			ones[i] += val
// 		}
// 	}
// 	return ones
// }

// func Solve(input string) (result int) {
// 	lines := strings.Split(input, "\n")
// 	if len(lines) == 0 {
// 		panic("no input")
// 	}
// 	var width = len(lines[0])
// 	var ones = countOnes(lines)
// 	// fmt.Println(ones)
// 	var height = len(lines)
// 	var bits string
// 	for _, one := range ones {
// 		if one > (height+1)/2 {
// 			bits += "1"
// 		} else {
// 			bits += "0"
// 		}
// 	}
// 	gammaRate, err := strconv.ParseUint(bits, 2, 0)
// 	if err != nil {
// 		panic(err)
// 	}
// 	mask := uint64(1<<width) - 1
// 	epsilonRate := ^gammaRate
// 	epsilonRate = epsilonRate & mask
// 	fmt.Println(strconv.FormatUint(gammaRate, 2), strconv.FormatUint(epsilonRate, 2))
// 	result = int(epsilonRate) * int(gammaRate)
// 	return
// }

func reduceBits(lines []string, inverse bool) int {
	var width = len(lines[0])
	for x := range width {
		var height = len(lines)
		var zeroCount int
		var zeroLines, oneLines []string
		for y := range height {
			bit := string(lines[y][x])
			if bit == "0" {
				zeroCount += 1
				zeroLines = append(zeroLines, lines[y])
			} else {
				oneLines = append(oneLines, lines[y])
			}
		}

		keepZeros := (zeroCount > height/2) != inverse
		if keepZeros {
			lines = zeroLines
		} else {
			lines = oneLines
		}

		if len(lines) == 1 {
			var result, err = strconv.ParseInt(lines[0], 2, 0)
			if err != nil {
				panic(err)
			}
			return int(result)
		}
	}
	panic("unexpected error")
}

func Solve(input string) (result int) {
	lines := strings.Split(input, "\n")
	if len(lines) == 0 {
		panic("no input")
	}

	oxygenRating := reduceBits(lines, false)
	co2Rating := reduceBits(lines, true)

	println(oxygenRating, co2Rating)

	result = oxygenRating * co2Rating
	return
}
