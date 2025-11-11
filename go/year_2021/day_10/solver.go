package day_10

import (
	_ "embed"
	"errors"
	"fmt"
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

type SyntaxError struct {
	Expected rune
	Found    rune
}

func (err SyntaxError) Error() string {
	return fmt.Sprintf("Expected %v, but found %v instead.", string(err.Expected), string(err.Found))
}

func completePair(line string) (string, error) {
	buffer := []rune{}
	for _, char := range line {
		pair := pairs[char]
		isClosing := pair == 0
		if isClosing {
			latestOpen := buffer[len(buffer)-1]
			if pairs[latestOpen] != char {
				return "", SyntaxError{Expected: pairs[latestOpen], Found: char}
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
	slices.Reverse(buffer)
	invertedBuffer := strings.Map(func(open rune) rune { return pairs[open] }, string(buffer))
	return invertedBuffer, nil
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
	var scores []int
	for line := range strings.SplitSeq(input, "\n") {
		tail, err := completePair(line)
		if errors.As(err, &SyntaxError{}) {
			continue
		} else if err != nil {
			panic(err)
		}
		var score int
		for _, closing := range tail {
			score *= 5
			switch closing {
			case ')':
				score += 1
			case ']':
				score += 2
			case '}':
				score += 3
			case '>':
				score += 4
			}
		}
		scores = append(scores, score)
	}
	slices.Sort(scores)
	result = scores[len(scores)/2]
	return
}
