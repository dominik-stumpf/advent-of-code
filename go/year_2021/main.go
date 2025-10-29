package main

import (
	"aoc/year_2021/day_03"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	result := day_03.Solve(strings.TrimSpace(day_03.Input))
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
