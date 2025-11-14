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
	// newTemplate := make(PolymerTemplate, len(*template)+len(inserts))
	// for key, value := range inserts {
	// 	newTemplate[key] = value
	// }
	// var offset int
	// for i, value := range newTemplate {
	// 	if value == 0 {
	// 		newTemplate[i] = (*template)[i-offset]
	// 	} else {
	// 		offset++
	// 	}
	// }
	// fmt.Printf("%s\n%s\n", template, newTemplate)
	// *template = newTemplate
}

func (template *PolymerTemplate) countElements() map[byte]int {
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
	elements := template.countElements()
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

func SolvePartTwo(input string) (result int) {
	template, rules := parseInput(input)
	for range 10 {
		template.Step(rules)
		fmt.Println(len(template))
	}
	elements := template.countElements()
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
