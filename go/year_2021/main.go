package main

import (
	active_day "aoc/year_2021/day_06"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	input := strings.TrimSpace(active_day.Input)
	resultPartTwo := active_day.SolvePartTwo(input)
	if resultPartTwo == 0 {
		resultPartOne := active_day.SolvePartOne(input)
		handleResult(resultPartOne)
	} else {
		handleResult(resultPartTwo)
	}
}

func handleResult(result int) {
	fmt.Printf("\nresult: %d\n", result)
	if result == 0 {
		return
	}
	cmd := exec.Command("wl-copy", strconv.Itoa(result))
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("failed to copy to clipboard", err)
	}
}
