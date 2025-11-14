package main

import (
	active_day "aoc/year_2021/day_14"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func printDuration(duration time.Duration) {
	fmt.Printf("\nduration: %v\n", duration)
}

func main() {
	input := strings.TrimSpace(active_day.Input)
	start := time.Now()
	resultPartTwo := active_day.SolvePartTwo(input)
	duration := time.Since(start)
	if resultPartTwo == 0 {
		start := time.Now()
		resultPartOne := active_day.SolvePartOne(input)
		duration := time.Since(start)
		printDuration(duration)
		handleResult(resultPartOne)
	} else {
		printDuration(duration)
		handleResult(resultPartTwo)
	}
}

func handleResult(result int) {
	fmt.Printf("result: %d\n", result)
	if result == 0 {
		return
	}
	cmd := exec.Command("wl-copy", strconv.Itoa(result))
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("failed to copy to clipboard", err)
	}
}
