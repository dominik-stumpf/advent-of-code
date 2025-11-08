package main

import (
	"aoc/year_2021/day_05"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	resultPartOne, resultPartTwo := day_05.SolveBoth(strings.TrimSpace(day_05.Input))
	if resultPartTwo == 0 {
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
