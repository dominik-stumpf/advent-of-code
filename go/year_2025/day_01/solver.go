package day_01

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input
var Input string

func SolvePartOne(input string) (result int) {
	// fmt.Printf("%v\n", (100+50-68)%100) // 82
	// fmt.Printf("%v\n", (100+82-30)%100) // 52
	// fmt.Printf("%v\n", (100+52+48)%100) // 0
	// fmt.Printf("%v\n", (100+0-5)%100)   // 95
	dial := 50
	for line := range strings.SplitSeq(input, "\n") {
		direction := line[0]
		rotation, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		if direction == 'L' {
			rotation = -rotation
		}
		dial = (100 + dial + rotation) % 100
		fmt.Printf("%q %d %d\n", direction, rotation, dial)
		if dial == 0 {
			result++
		}
	}
	_ = dial
	return
}

func SolvePartTwo(input string) (result int) {
	return
}
