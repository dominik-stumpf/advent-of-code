package day_14

import (
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	if len(Input) != 821 {
		t.Skipf("invalid input")
		return
	}
	result := SolvePartOne(Input)
	answer := 2435
	if result != answer {
		t.Errorf("expected %d, received %d", answer, result)
	}
}
