package dierollerpkg

import "fmt"

type DieRollResult struct {
	Result int64
	Rolls  DieRolls
}

// override convert to string function
func (drr DieRollResult) String() string {
	return fmt.Sprintf("Result: %d (Rolls: %v)", drr.Result, drr.Rolls)
}
