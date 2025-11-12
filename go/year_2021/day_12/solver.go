package day_12

import (
	_ "embed"
	"fmt"
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
				paths = append(paths, append(path, "end"))
				continue
			}
			if node == strings.ToLower(node) {
				for _, p := range path {
					if p != "start" && p == node {
						continue neighbor
					}
				}
			}
			tracer(node, append(path, node))
		}
	}

	tracer("start", []string{"start"})
	return paths
}

func SolvePartOne(input string) (result int) {
	nodes := parseInput(input)
	paths := traverse(nodes)
	fmt.Println(len(paths))

	return
}

func SolvePartTwo(input string) (result int) {
	return
}
