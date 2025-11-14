package day_13

import (
	"cmp"
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input
var Input string

type Position struct {
	X int
	Y int
}

type FoldInstruction struct {
	Position int
	FoldAxis rune
}

type Manual struct {
	Dots  []*Position
	Folds []FoldInstruction
}

func (manual Manual) Fold(fold FoldInstruction) {
	if fold.FoldAxis == 'y' {
		for _, dot := range manual.Dots {
			if dot.Y <= fold.Position {
				continue
			}
			dot.Y = abs(dot.Y - fold.Position*2)
		}
	} else {
		for _, dot := range manual.Dots {
			if dot.X <= fold.Position {
				continue
			}
			dot.X = abs(dot.X - fold.Position*2)
		}
	}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func (manual Manual) String() string {
	xMax := slices.MaxFunc(manual.Dots, func(a, b *Position) int {
		return cmp.Compare(a.X, b.X)
	}).X + 1
	yMax := slices.MaxFunc(manual.Dots, func(a, b *Position) int {
		return cmp.Compare(a.Y, b.Y)
	}).Y + 1
	matrix := make([][]rune, yMax)
	for y := range yMax {
		row := make([]rune, xMax)
		for x := range row {
			row[x] = '.'
		}
		matrix[y] = row
	}
	for _, dot := range manual.Dots {
		matrix[dot.Y][dot.X] = '#'
	}
	result := make([]string, yMax)
	for y, row := range matrix {
		result[y] = string(row)
	}
	return strings.Join(result, "\n")
}

func parseInput(input string) (result Manual) {
	parts := strings.Split(input, "\n\n")
	dots, folds := parts[0], parts[1]
	for dot := range strings.SplitSeq(dots, "\n") {
		var X, Y int
		fmt.Sscanf(dot, "%d,%d", &X, &Y)
		result.Dots = append(result.Dots, &Position{X, Y})
	}
	for fold := range strings.SplitSeq(folds, "\n") {
		parts = strings.Split(fold, "=")
		Position, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		FoldAxis := rune(parts[0][len(parts[0])-1])
		result.Folds = append(result.Folds, FoldInstruction{Position, FoldAxis})
	}
	return
}

func SolvePartOne(input string) (result int) {
	manual := parseInput(input)
	fmt.Print(manual, "\n\n\n")
	manual.Fold(manual.Folds[0])
	fmt.Println(manual)

	dots := map[string]struct{}{}
	for _, dot := range manual.Dots {
		key := fmt.Sprintf("%d-%d", dot.X, dot.Y)
		dots[key] = struct{}{}
	}
	result = len(dots)
	return
}

func SolvePartTwo(input string) (result int) {
	manual := parseInput(input)
	for _, fold := range manual.Folds {
		manual.Fold(fold)
	}
	fmt.Printf("%v\n", manual)
	result = -1
	return
}
