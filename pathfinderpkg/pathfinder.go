package pathfinderpkg

import (
	"sort"
)

// cost to buy ability score
// 3 through 6 are not legal values
// and are extrapolations mirroring the
// top end of the scale for comparison
// only
func AbilityCost(ability int) (int, bool) {
	ability_cost := map[int]int{
		1:  -25,
		2:  -20,
		3:  -16,
		4:  -12,
		5:  -9,
		6:  -6,
		7:  -4,
		8:  -2,
		9:  -1,
		10: 0,
		11: 1,
		12: 2,
		13: 3,
		14: 5,
		15: 7,
		16: 10,
		17: 13,
		18: 17,
		19: 21,
		20: 26,
		21: 31,
		22: 37,
		23: 43,
		24: 50,
		25: 57,
		26: 65,
		27: 73,
		28: 82,
		29: 91,
		30: 101,
		31: 111,
		32: 122,
		33: 133,
		34: 145,
		35: 157,
		36: 170,
		37: 183,
		38: 197,
		39: 211,
		40: 226,
		41: 241,
		42: 257,
		43: 273,
		44: 290,
		45: 307,
	}
	value, ok := ability_cost[ability]
	return value, ok
}

// (define (cost-abilities ab)
//   (apply + (map ability->cost ab)))
func AbilityModifier(ability int) (int, bool) {
	ability_modifier := map[int]int{
		1:  -5,
		2:  -4,
		3:  -4,
		4:  -3,
		5:  -3,
		6:  -2,
		7:  -2,
		8:  -1,
		9:  -1,
		10: 0,
		11: 0,
		12: 1,
		13: 1,
		14: 2,
		15: 2,
		16: 3,
		17: 3,
		18: 4,
		19: 4,
		20: 5,
		21: 5,
		22: 6,
		23: 6,
		24: 7,
		25: 7,
		26: 8,
		27: 8,
		28: 9,
		29: 9,
		30: 10,
		31: 10,
		32: 11,
		33: 11,
		34: 12,
		35: 12,
		36: 13,
		37: 13,
		38: 14,
		39: 14,
		40: 15,
		41: 15,
		42: 16,
		43: 16,
		44: 17,
		45: 17,
	}
	value, ok := ability_modifier[ability]
	return value, ok
}

type Abilities []int
type MapFunc func(int) int
type ApplyFunc func(x, y int) int

// implement sort.Interface
func (ab Abilities) Len() int           { return len(ab) }
func (ab Abilities) Less(i, j int) bool { return ab[i] < ab[j] }
func (ab Abilities) Swap(i, j int)      { ab[i], ab[j] = ab[j], ab[i] }

// for chaining...i.e. chained calls like in functional paradigm
func (ab Abilities) Sortf() Abilities { sort.Sort(ab); return ab }

// apply a function to each ability returning the new Abilities
// with the result of the function call mapped to the position of the
// value passed to the function.  For example, when ab is {1, 2, 3}
// ab.Map(times2) returns a new Ability {2, 4, 6} and ab is unchanged
func (ab Abilities) Map(fn MapFunc) Abilities {
	result := make(Abilities, ab.Len())
	for i, v := range ab {
		result[i] = fn(v)
	}
	return result
}

// apply a function to each ability, changing Abilities in place
// with the result of the function call mapped to the position of the
// value passed to the function.  For exmple, when ab is {1, 2, 3}
// ab.Map!(times2) returns ab that has been changed to {2, 4, 6}
func (ab Abilities) MapInPlace(fn MapFunc) Abilities {
	for i, v := range ab {
		ab[i] = fn(v)
	}
	return ab
}

// apply a function to each ability and an accumulated value
// with the result being the value in the accumulator after function
// has been applied to all values. For example, when ab is {1, 2, 3}
//
// ab.Apply(0, sum) => 6
// ab.Apply(1, sum) => 7
// ab.Apply(10, sum) => 16
// ab.Apply(0, prod) => 0
// ab.Apply(1, prod) => 6
// ab.Apply(10, prod) => 60
func (ab Abilities) Apply(initialValue int, fn ApplyFunc) int {
	accumulator := initialValue
	for _, v := range ab {
		accumulator = fn(accumulator, v)
	}
	return accumulator
}

func SumCostOfAbilities(abilities Abilities) int {
	total := 0
	for _, v := range abilities {
		if value, ok := AbilityCost(v); ok {
			total += value
		}
	}
	return total
}

func SumModifiersFromAbilities(abilities Abilities) int {
	total := 0
	for _, v := range abilities {
		if value, ok := AbilityModifier(v); ok {
			total += value
		}
	}
	return total
}
