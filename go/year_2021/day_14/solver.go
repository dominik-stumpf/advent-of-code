package day_14

import (
	"cmp"
	_ "embed"
	"fmt"
	"maps"
	"slices"
	"strings"
)

//go:embed input
var input string
var Input = strings.TrimSpace(input)

type PolymerTemplate []byte
type PairInsertionRules map[string]byte

func parseInput(input string) (template PolymerTemplate, rules PairInsertionRules) {
	segments := strings.Split(input, "\n\n")
	template = PolymerTemplate(segments[0])
	rules = map[string]byte{}
	for line := range strings.SplitSeq(segments[1], "\n") {
		segments = strings.Split(line, " -> ")
		rules[segments[0]] = byte(segments[1][0])
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
	inserts := map[int]byte{}
	for pattern, insert := range rules {
		indices := findAllIndex(string(*template), pattern)
		for _, i := range indices {
			inserts[i] = insert
		}
	}
	sortedKeys := slices.Collect(maps.Keys(inserts))
	slices.SortFunc(sortedKeys, func(a, b int) int {
		return cmp.Compare(b, a)
	})
	for _, key := range sortedKeys {
		*template = slices.Insert(*template, key+1, inserts[key])
	}
}

func stepFast(rules PairInsertionRules, state *map[string]int, diff *map[byte]int) {
	nextState := map[string]int{}
	for pattern, count := range *state {
		insert := rules[pattern]
		(*diff)[insert] += 1 * count
		variants := [2][]byte{{pattern[0], insert}, {insert, pattern[1]}}
		for _, variant := range variants {
			nextState[string(variant)] += 1 * count
		}
	}
	*state = nextState
}

func (template *PolymerTemplate) ConvertToPatterns(rules PairInsertionRules) map[string]int {
	result := map[string]int{}
	for pattern := range rules {
		indices := findAllIndex(string(*template), pattern)
		if len(indices) > 0 {
			result[pattern] = len(indices)
		}
	}
	return result
}

func (template *PolymerTemplate) CountElements() map[byte]int {
	result := map[byte]int{}
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
	elements := template.CountElements()
	keys := slices.Collect(maps.Keys(elements))
	max := slices.MaxFunc(keys, func(a, b byte) int {
		return cmp.Compare(elements[a], elements[b])
	})
	min := slices.MinFunc(keys, func(a, b byte) int {
		return cmp.Compare(elements[a], elements[b])
	})
	fmt.Println(len(template), elements)
	result = elements[max] - elements[min]
	return
}

func SolvePartTwo(input string) (result int) {
	template, rules := parseInput(input)
	step := 40
	patterns := template.ConvertToPatterns(rules)
	elements := map[byte]int{}
	for _, element := range template {
		elements[element]++
	}
	for range step {
		stepFast(rules, &patterns, &elements)
	}
	keys := slices.Collect(maps.Keys(elements))
	max := slices.MaxFunc(keys, func(a, b byte) int {
		return cmp.Compare(elements[a], elements[b])
	})
	min := slices.MinFunc(keys, func(a, b byte) int {
		return cmp.Compare(elements[a], elements[b])
	})
	result = elements[max] - elements[min]
	return
}
