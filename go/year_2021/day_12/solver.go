package day_12

import (
	_ "embed"
	"maps"
	"slices"
	"strings"
)

//go:embed input
var Input string

// type Node string
type NodeMap = map[string][]string

func parseInput(input string) NodeMap {
	nodes := NodeMap{}
	for line := range strings.SplitSeq(input, "\n") {
		segments := strings.Split(line, "-")
		nodes[segments[0]] = append(nodes[segments[0]], segments[1])
		nodes[segments[1]] = append(nodes[segments[1]], segments[0])

	}
	return nodes
}

func traverse(nodes NodeMap) [][]string {
	var tracer func(startNode string, path []string)
	paths := [][]string{}

	tracer = func(startNode string, path []string) {
		neighbors := nodes[startNode]
	neighbor:
		for _, node := range neighbors {
			if node == "start" {
				continue
			}
			if node == "end" {
				paths = append(paths, path)
				continue
			}
			if node == strings.ToLower(node) {
				for _, p := range path {
					if p == node {
						continue neighbor
					}
				}
			}
			tracer(node, append(path, node))
		}
	}

	tracer("start", []string{})
	return paths
}

func traverseWithRevisit(nodes NodeMap) [][]string {
	var tracer func(startNode string, path []string)
	paths := [][]string{}

	tracer = func(startNode string, path []string) {
		neighbors := nodes[startNode]
	neighbor:
		for _, node := range neighbors {
			if node == "start" {
				continue
			}
			if node == "end" {
				paths = append(paths, path)
				continue
			}
			if len(path) > 0 && node == strings.ToLower(node) {
				visitCounts := map[string]int{}
				for _, prevNode := range path {
					if prevNode == strings.ToLower(prevNode) {
						visitCounts[prevNode]++
					}
				}
				var maxVisitCount int
				if len(visitCounts) > 0 {
					maxVisitCount = slices.Max(slices.Collect(maps.Values(visitCounts)))
				}
				if visitCounts[node] == 1 && maxVisitCount >= 2 || visitCounts[node] >= 2 {
					continue neighbor
				}
			}
			tracer(node, append(path, node))
		}
	}

	tracer("start", []string{})
	return paths
}

func SolvePartOne(input string) (result int) {
	nodes := parseInput(input)
	paths := traverse(nodes)
	result = len(paths)

	return
}

func SolvePartTwo(input string) (result int) {
	nodes := parseInput(input)
	paths := traverseWithRevisit(nodes)
	// for _, path := range paths {
	// 	fmt.Println(path)
	// 	// fmt.Println(strings.Join(slices.Insert(append(path, "end"), 0, "start"), ","))
	// }
	result = len(paths) // 144603

	return
}
