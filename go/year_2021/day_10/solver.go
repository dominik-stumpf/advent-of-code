package day_10

import (
	_ "embed"
	"slices"
	"strings"
)

//go:embed input
var Input string

var pairs = map[rune]rune{
	'{': '}',
	'(': ')',
	'<': '>',
	'[': ']',
}

func findCorruptedClosingPair(line string) rune {
	buffer := []rune{}
	for _, char := range line {
		pair := pairs[char]
		isClosing := pair == 0
		if isClosing {
			latestOpen := buffer[len(buffer)-1]
			if pairs[latestOpen] != char {
				// fmt.Printf("Expected %v, but found %v instead.\n", string(pairs[latestOpen]), string(char))
				return char
			}
			for j := len(buffer) - 1; j >= 0; j -= 1 {
				opening := buffer[j]
				if pairs[opening] == char {
					buffer = slices.Delete(buffer, j, j+1)
					break
				}
			}
		} else {
			buffer = append(buffer, char)
		}
	}
	return 0
}

func SolvePartOne(input string) (result int) {
	for line := range strings.SplitSeq(input, "\n") {
		closing := findCorruptedClosingPair(line)
		switch closing {
		case ')':
			result += 3
		case ']':
			result += 57
		case '}':
			result += 1197
		case '>':
			result += 25137
		}
	}
	return
}

func SolvePartTwo(input string) (result int) {
	return
}
