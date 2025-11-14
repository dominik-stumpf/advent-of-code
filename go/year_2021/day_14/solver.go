package day_14

import (
	"cmp"
	_ "embed"
	"maps"
	"slices"
	"strings"
)

//go:embed input
var Input string

type PolymerTemplate []rune
type PairInsertionRules map[string]rune

func parseInput(input string) (template PolymerTemplate, rules PairInsertionRules) {
	segments := strings.Split(input, "\n\n")
	template = PolymerTemplate(segments[0])
	rules = map[string]rune{}
	for line := range strings.SplitSeq(segments[1], "\n") {
		segments = strings.Split(line, " -> ")
		rules[segments[0]] = rune(segments[1][0])
	}
	return
}

func findAllIndex(s string, substr string) (result []int) {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			result = append(result, i)
		}
	}
	return
}

func (template *PolymerTemplate) Step(rules PairInsertionRules) {
	inserts := map[int]rune{}
	for pattern, insert := range rules {
		indices := findAllIndex(string(*template), pattern)
		for _, i := range indices {
			inserts[i+1] = insert
		}
	}
	sortedKeys := slices.Collect(maps.Keys(inserts))
	slices.SortFunc(sortedKeys, func(a, b int) int {
		return cmp.Compare(b, a)
	})
	for _, key := range sortedKeys {
		*template = slices.Insert(*template, key, inserts[key])
	}
}

func (template *PolymerTemplate) countElements() map[rune]int {
	result := map[rune]int{}
	for _, r := range *template {
		result[r]++
	}
	return result
}

func SolvePartOne(input string) (result int) {
	template, rules := parseInput(input)
	for range 10 {
		template.Step(rules)
	}
	elements := template.countElements()
	keys := slices.Collect(maps.Keys(elements))
	max := slices.MaxFunc(keys, func(a, b rune) int {
		return cmp.Compare(elements[a], elements[b])
	})
	min := slices.MinFunc(keys, func(a, b rune) int {
		return cmp.Compare(elements[a], elements[b])
	})
	result = elements[max] - elements[min]
	return
}

func SolvePartTwo(input string) (result int) {
	return
}
