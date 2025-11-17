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

// 3 map[CB:1 NC:1 NN:1]
// CH HB | NB BC | NC CN
// 6 map[BC:1 CH:1 CN:1 HB:1 NB:1 NC:1]
func stepFast(rules PairInsertionRules, state map[string]int) map[string]int {
	result := map[string]int{}
	for pattern, count := range state {
		insert := rules[pattern]
		variants := [2][]byte{{pattern[0], insert}, {insert, pattern[1]}}
		for _, variant := range variants {
			result[string(variant)] += 1 * count
		}
	}

	return result
}
func castElements(statec map[string]int, length int) map[byte]int {
	var result map[byte]int

	var target int
	for range 1000000 {
		if length == target {
			// fmt.Println("castElements:", string(result), target, length)
			break
		}
		var state = map[string]int{}
		for key, value := range statec {
			state[key] = value
		}
		var next byte
		target = 0
		result = map[byte]int{}
		var notFound bool
		for {
			if len(state) == 0 {
				// result = append(result, next)
				result[next]++
				target += 1
				// fmt.Println("exhaust")
				break
			}
			if notFound {
				// fmt.Println("not found")
				break
			}
			// fmt.Println(string(next), state)
			notFound = true
			for key, _ := range state {
				if key[0] != next && next != 0 {
					continue
				}
				notFound = false
				next = key[1]
				result[key[0]]++
				// result = append(result, key[0])
				target += 1
				// for key, _ := range state {
				// 	if key[0] == next {
				// 	}
				// }
				state[key]--
				if state[key] <= 0 {
					delete(state, key)
				}
				break
			}
		}
		fmt.Println(target, length)
	}

	return result
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
	step := 20
	expectedLength := len(template)
	patterns := template.ConvertToPatterns(rules)
	for range step {
		expectedLength = (expectedLength * 2) - 1
		// fmt.Println(string(template), patterns)
		// patternsOld := template.ConvertToPatterns(rules)
		// fmt.Println(patternsOld, patterns)
		// fmt.Println(maps.Equal(patternsOld, patterns))
		patterns = stepFast(rules, patterns)
		// for key := range patterns {
		// 	fmt.Print(string(key[0])+string(rules[key])+string(key[1]), ",")
		// }
		// fmt.Println(len(template))
		// template.Step(rules)
	}
	// elements := map[byte]int{}
	// for pattern, count := range patterns {
	// 	for _, char := range []byte(pattern) {
	// 		elements[char] += (1 * count)
	// 	}
	// }
	// var elementsSum int
	// for _, count := range elements {
	// 	elementsSum += count
	// }
	// elementsOld := template.CountElements()
	// // fmt.Println(string(template), patterns)
	// fmt.Println(elementsSum, len(template), expectedLength)
	// fmt.Print("\n")
	// for key, value := range elements {
	// 	fmt.Print(string(key), " ", value, elementsOld[key], "\n")
	// }
	// // fmt.Println(elementsOld)
	// keys := slices.Collect(maps.Keys(elements))
	// max := slices.MaxFunc(keys, func(a, b byte) int {
	// 	return cmp.Compare(elements[a], elements[b])
	// })
	// min := slices.MinFunc(keys, func(a, b byte) int {
	// 	return cmp.Compare(elements[a], elements[b])
	// })
	// result = elements[max] - elements[min]

	elements := castElements(patterns, expectedLength)
	fmt.Println(elements)
	var elementsSum int
	for _, count := range elements {
		elementsSum += count
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
