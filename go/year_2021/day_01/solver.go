package day_01

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input
var Input string

// func Solve(input string) (result int) {
// 	var parse []string
// 	parse = strings.Split(input, "\n")

// 	var prev int = -1
// 	for _, s := range parse {
// 		v, err := strconv.Atoi(s)
// 		if err != nil {
// 			panic(err)
// 		}
// 		if v > prev && prev != -1 {
// 			result++
// 		}
// 		prev = v
// 	}
// 	return
// }

func Solve(input string) (result int) {
	var parse []string
	parse = strings.Split(input, "\n")

	var prev int = -1
	var window [3]int
	for i, s := range parse {
		v, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		window[i%3] = v
		if i <= 2 {
			continue
		}
		windowSum := window[0] + window[1] + window[2]
		if windowSum > prev {
			result++
		}
		prev = windowSum
	}
	return
}
