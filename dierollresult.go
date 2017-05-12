package main

import "fmt"

type DieRolls []uint32

func (dr DieRolls) Len() int           { return len(dr) }
func (dr DieRolls) Swap(i, j int)      { dr[i], dr[j] = dr[j], dr[i] }
func (dr DieRolls) Less(i, j int) bool { return dr[i] > dr[j] } // reverse sorted
func (dr DieRolls) String() string {
	// return fmt.Sprintf("(%v)", strings.Trim(strings.Join(strings.Fields(fmt.Sprintf("%v", dr)), ","), "[]"))
	return fmt.Sprintf("(%d,%d,%d,%d,%d,%d)", dr[0], dr[1], dr[2], dr[3], dr[4], dr[5])
}

type DieRollResult struct {
	Result int64
	Rolls  DieRolls
}

func (drr DieRollResult) String() string {
	return fmt.Sprintf("Result: %d (Rolls: %v)", drr.Result, drr.Rolls)
}
