package day_08

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input
var Input string

type Signal string

type NoteEntry struct {
	Patterns []Signal
	Output   []Signal
}

type Digit string

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

func parseInput(input string) (result []NoteEntry) {
	for line := range strings.SplitSeq(input, "\n") {
		segments := strings.Split(line, " | ")
		convToSignals := func(signals string) (result []Signal) {
			for signal := range strings.SplitSeq(signals, " ") {
				result = append(result, Signal(signal))
			}
			return
		}
		Patterns := convToSignals(segments[0])
		Output := convToSignals(segments[1])
		result = append(result, NoteEntry{Patterns, Output})
	}
	return
}

func SolvePartOne(input string) (result int) {
	notes := parseInput(input)
	digits := make(map[Digit]int)
	for _, note := range notes {
		for _, digit := range []Digit{One, Four, Seven, Eight} {
			for _, signal := range note.Output {
				if len(digit) == len(signal) {
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
	return
}
