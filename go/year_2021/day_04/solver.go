package day_04

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input
var Input string

type Board [][]int

func (board *Board) mark(guess int) {
	for x, row := range *board {
		for y, col := range row {
			if col == guess {
				(*board)[x][y] = -1
			}
		}
	}
}

func hasCompleteSequence(cells []int) bool {
	for _, cell := range cells {
		if cell != -1 {
			return false
		}
	}
	return true
}

func (board Board) checkSequence() (hasSequence bool) {
	width := len(board)
	height := len(board[0])

	for x := range width {
		column := make([]int, height)
		for y := range height {
			column[y] = board[y][x]
		}
		if hasCompleteSequence(column) {
			return true
		}
	}

	for y := range height {
		if hasCompleteSequence(board[y]) {
			return true
		}
	}

	return
}

func (board Board) sum() (result int) {
	for _, row := range board {
		for _, col := range row {
			if col != -1 {
				result += col
			}
		}
	}
	return
}

func (board Board) print() {
	for _, row := range board {
		fmt.Printf("%2d\n", row)
	}
}

type Puzzle struct {
	Guesses []int
	Boards  []Board
}

func parsePuzzle(puzzle string) (result Puzzle) {
	lines := strings.Split(puzzle, "\n\n")

	for chars := range strings.SplitSeq(lines[0], ",") {
		n, err := strconv.Atoi(chars)
		if err != nil {
			panic(err)
		}
		result.Guesses = append(result.Guesses, n)
	}

	for _, board := range lines[1:] {
		boardRows := strings.Split(board, "\n")
		board := make(Board, len(boardRows))
		for i, boardRow := range boardRows {
			for chars := range strings.SplitSeq(boardRow, " ") {
				if chars == "" {
					continue
				}
				n, err := strconv.Atoi(chars)
				if err != nil {
					panic(err)
				}
				board[i] = append(board[i], n)
			}
		}
		result.Boards = append(result.Boards, board)
	}

	return
}

func Solve(input string) (result int) {
	puzzle := parsePuzzle(input)

	var lastGuess int
	var winningBoard Board
guessing:
	for _, guess := range puzzle.Guesses {
		lastGuess = guess
		for _, board := range puzzle.Boards {
			board.mark(guess)
			if board.checkSequence() {
				winningBoard = board
				break guessing
			}
		}
	}

	for _, board := range puzzle.Boards {
		board.print()
		fmt.Printf("\n")
	}

	result = lastGuess * winningBoard.sum()

	return
}
