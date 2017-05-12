package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type DieRoll struct {
	Dice     uint32
	Sides    uint32
	Modifier DieModifier
	Keep     uint32
	History  DieRollHistory
}

func (dr DieRoll) Roll() DieRollResult {
	// generate the rolls
	rolls := DieRolls{}
	fmt.Printf("rolls = %v", rolls)
	rand.Seed(time.Now().UnixNano())
	for i := range rolls {
		rolls[i] = uint32(rand.Int63n(int64(dr.Sides)))
	}
	fmt.Printf("random rolls = %v", rolls)
	// sort them from highest to lowest
	// because Less for DieRolls is defined in reverse (>)
	sort.Sort(rolls)
	fmt.Printf("sorted rolls = %v", rolls)
	// sum the highest keep rolls
	var sum int64
	for i := range rolls[:dr.Keep] {
		sum += int64(i)
	}
	// apply the modifier
	switch dr.Modifier.ModType {
	case DieModifierTypeAdd:
		sum += int64(dr.Modifier.Amount)
	case DieModifierTypeMultiply:
		sum *= int64(dr.Modifier.Amount)
	case DieModifierTypeSubtract:
		sum -= int64(dr.Modifier.Amount)
	}
	result := DieRollResult{sum, rolls}
	dr.History.Append(result)
	return result
}

func (dr DieRoll) String() string {
	return fmt.Sprintf("dice: %d sides: %d mod: %v keep: %d history: (%v)", dr.Dice, dr.Sides, dr.Modifier, dr.Keep, dr.History)
}
