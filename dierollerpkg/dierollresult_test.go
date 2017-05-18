package dierollerpkg

import (
	"testing"
)

func TestDieRollResultString(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping string test in short mode.")
	}
	dr := DieRolls{1, 2, 3, 4, 5}
	drr := DieRollResult{45, dr}
	result := drr.String()
	s := "Result: 45 (Rolls: 1,2,3,4,5)"
	if result != s {
		t.Errorf("Expected drr.String() to be %s, got %s", s, result)
	}
}
