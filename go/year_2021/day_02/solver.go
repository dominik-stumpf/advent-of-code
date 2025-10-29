package day_02

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input
var Input string

func parseLine(line string) (direction string, value int) {
	parts := strings.Split(line, " ")
	direction = parts[0]
	value, _ = strconv.Atoi(parts[1])
	return
}

// func Solve(input string) (result int) {
// 	lines := strings.Split(input, "\n")
// 	var horizontal int
// 	var vertical int
// 	for _, line := range lines {
// 		direction, value := parseLine(line)
// 		switch direction {
// 		case "forward":
// 			horizontal += value
// 		case "down":
// 			vertical += value
// 		case "up":
// 			vertical -= value
// 		}
// 	}
// 	result = horizontal * vertical
// 	return
// }

func Solve(input string) (result int) {
	lines := strings.Split(input, "\n")
	var horizontal, aim, depth int
	for _, line := range lines {
		direction, value := parseLine(line)
		switch direction {
		case "forward":
			horizontal += value
			depth += value * aim
		case "down":
			aim += value
		case "up":
			aim -= value
		}
	}
	result = horizontal * depth
	return
}
