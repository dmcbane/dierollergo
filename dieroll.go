package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

type DieRoll struct {
	dice     uint32
	sides    uint32
	modifier DieModifier
	keep     uint32
	history  []DieRollResult
}

// constructor
func NewDieRoll(dice, sides uint32, modifier string, keep uint32) *DieRoll {
	return &DieRoll{
		dice:     dice,
		sides:    sides,
		modifier: *DieModifierParse(modifier),
		keep:     keep,
		history:  make([]DieRollResult, 0),
	}
}

// getters
func (dr *DieRoll) Dice() uint32 {
	return dr.dice
}
func (dr *DieRoll) Sides() uint32 {
	return dr.sides
}
func (dr *DieRoll) Modifier() DieModifier {
	return dr.modifier
}
func (dr *DieRoll) Keep() uint32 {
	return dr.keep
}
func (dr *DieRoll) History() []DieRollResult {
	return dr.history
}

// Last result returned by die roller
func (dr *DieRoll) LastRoll() DieRollResult {
	return dr.history[len(dr.history)-1]
}

// alias for LastRoll
var Result = (*DieRoll).LastRoll

// Appender
func (dr *DieRoll) AddHistory(drr DieRollResult) {
	old := dr.History()
	temp := append(old, drr)
	dr.history = temp
}

// Roll a die
func (dr *DieRoll) Roll() DieRollResult {
	// generate the rolls
	rolls := make(DieRolls, dr.Dice())
	rand.Seed(time.Now().UnixNano())
	for i := range rolls {
		rolls[i] = uint32(rand.Int63n(int64(dr.Sides()))) + uint32(1)
	}
	// sort them from highest to lowest
	// because Less for DieRolls is defined in reverse (>)
	sort.Sort(rolls)
	// sum the highest keep rolls
	sum := int64(0)
	kept := rolls[0:dr.Keep()]
	for _, v := range kept {
		sum += int64(v)
	}
	// apply the modifier
	switch dr.Modifier().ModType {
	case DieModifierTypeAdd:
		sum += int64(dr.Modifier().Amount)
	case DieModifierTypeMultiply:
		sum *= int64(dr.Modifier().Amount)
	case DieModifierTypeSubtract:
		sum -= int64(dr.Modifier().Amount)
	}
	result := DieRollResult{sum, kept}
	dr.AddHistory(result)
	return result
}

// override normal convert to string method
func (dr *DieRoll) String() string {
	return fmt.Sprintf("dice: %d sides: %d mod: %v keep: %d history: (%s)", dr.Dice(), dr.Sides(), dr.Modifier(), dr.Keep(), dr.HistoryAsString())
}

// output as standardized string
func (dr *DieRoll) standardstring(verbose bool) string {
	d := fmt.Sprintf("%dD%d", dr.Dice(), dr.Sides())
	k := fmt.Sprintf("K%d", dr.Keep())
	if dr.Keep() == dr.Dice() {
		k = ""
	}
	m := fmt.Sprintf("%v", dr.Modifier())
	r := ""
	if verbose {
		r = fmt.Sprintf(" (%s)", dr.LastRoll().Rolls)
	}
	return strings.Trim(fmt.Sprintf("%s%s%s%s", d, k, m, r), "")
}

func (dr *DieRoll) StandardString() string {
	return dr.standardstring(false)
}

func (dr *DieRoll) StandardStringVerbose() string {
	return dr.standardstring(true)
}

// convert history slice to a string
func (dr *DieRoll) HistoryAsString() string {
	s := make([]string, len(dr.history))
	for i, v := range dr.history {
		// convert to string and wrap with braces
		s[i] = fmt.Sprintf("{%s}", v)
	}
	return fmt.Sprintf("%d entries: %s", len(s), strings.Join(s, ","))
}
