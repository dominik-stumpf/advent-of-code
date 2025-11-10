package day_08

import (
	_ "embed"
	"fmt"
	"maps"
	"slices"
	"strconv"
	"strings"
)

//go:embed input
var Input string

type Digit string

func (left Digit) exclude(right string) (segments []rune) {
	for _, leftSegment := range left {
		if !slices.Contains([]rune(right), leftSegment) {
			segments = append(segments, leftSegment)
		}
	}
	return
}

func (digit Digit) translate(reversedLexicon map[rune]rune) Digit {
	var result []rune
	for _, segment := range digit {
		result = append(result, reversedLexicon[segment])
	}
	return Digit(result)
}

type NoteEntry struct {
	Patterns []Digit
	Output   []Digit
}

const (
	Zero  Digit = "abcefg"
	One   Digit = "cf"
	Two   Digit = "acdeg"
	Three Digit = "acdfg"
	Four  Digit = "bcdf"
	Five  Digit = "abdfg"
	Six   Digit = "abdefg"
	Seven Digit = "acf"
	Eight Digit = "abcdefg"
	Nine  Digit = "abcdfg"
)

// segments:
//
//  aaaa
// b    c
// b    c
//  dddd
// e    f
// e    f
//  gggg

var digits = []Digit{Zero, One, Two, Three, Four, Five, Six, Seven, Eight, Nine}
var segmentCount = map[rune]int{
	'a': 8,
	'b': 6,
	'c': 8,
	'd': 7,
	'e': 4,
	'f': 9,
	'g': 7,
}
var segmentsWithDistincCount = []rune{'b', 'e', 'f'}

func parseInput(input string) (result []NoteEntry) {
	for line := range strings.SplitSeq(input, "\n") {
		segments := strings.Split(line, " | ")
		convToDigits := func(Digits string) (result []Digit) {
			for digit := range strings.SplitSeq(Digits, " ") {
				result = append(result, Digit(digit))
			}
			return
		}
		Patterns := convToDigits(segments[0])
		Output := convToDigits(segments[1])
		result = append(result, NoteEntry{Patterns, Output})
	}
	return
}

func SolvePartOne(input string) (result int) {
	notes := parseInput(input)
	digits := make(map[Digit]int)
	for _, note := range notes {
		for _, digit := range []Digit{One, Four, Seven, Eight} {
			for _, Digit := range note.Output {
				if len(digit) == len(Digit) {
					result += 1
					digits[digit] += 1
				}
			}
		}
		fmt.Println(note.Output)
	}
	fmt.Println(digits)
	return
}

func SolvePartTwo(input string) (result int) {
	notes := parseInput(input)
	for _, note := range notes {
		segmentLexicon := map[rune]rune{}
		digitLexicon := map[Digit]Digit{}
		for _, digit := range []Digit{One, Four, Seven, Eight} {
			for _, pattern := range note.Patterns {
				if len(digit) == len(pattern) {
					digitLexicon[digit] = pattern
				}
			}
		}

		// find 'b' 'e' 'f'
		mixedSegmentCount := map[rune]int{}
		for _, pattern := range note.Patterns {
			for _, segment := range pattern {
				mixedSegmentCount[segment]++
			}
		}
		for _, segment := range segmentsWithDistincCount {
			count := segmentCount[segment]
			for mixedSegment, mixedCount := range mixedSegmentCount {
				if mixedCount == count {
					segmentLexicon[segment] = mixedSegment
				}
			}
		}

		// find 'c'
		exclusion := digitLexicon[One].exclude(string(segmentLexicon['f']))
		if len(exclusion) != 1 {
			panic("'c' should be the only difference")
		}
		segmentLexicon['c'] = exclusion[0]

		// find 'd'
		exclusion = digitLexicon[Four].exclude(string(digitLexicon[One]) + string(segmentLexicon['b']))
		if len(exclusion) != 1 {
			panic("'d' should be the only difference")
		}
		segmentLexicon['d'] = exclusion[0]

		// find 'a'
		exclusion = digitLexicon[Seven].exclude(string(digitLexicon[One]))
		if len(exclusion) != 1 {
			panic("'a' should be the only difference")
		}
		segmentLexicon['a'] = exclusion[0]

		// find 'g'
		exclusion = digitLexicon[Eight].exclude(string(slices.Collect(maps.Values(segmentLexicon))))
		if len(exclusion) != 1 {
			panic("'g' should be the only difference")
		}
		segmentLexicon['g'] = exclusion[0]

		reversedLexicon := make(map[rune]rune, len(segmentLexicon))
		for key, value := range segmentLexicon {
			reversedLexicon[value] = key
		}
		var decoded string
		for _, outputDigit := range note.Output {
			translatedRunes := []rune(outputDigit.translate(reversedLexicon))
			slices.Sort([]rune(translatedRunes))
			translatedDigit := Digit(translatedRunes)

			for i, digit := range digits {
				if translatedDigit == digit {
					decoded += strconv.Itoa(i)
					break
				}
			}
		}
		partial, err := strconv.Atoi(decoded)
		if err != nil {
			panic(err)
		}
		result += partial
	}

	return
}
