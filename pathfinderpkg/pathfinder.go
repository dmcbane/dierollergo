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
	ability_cost := []int{
		-25,
		-20,
		-16,
		-12,
		-9,
		-6,
		-4,
		-2,
		-1,
		0,
		1,
		2,
		3,
		5,
		7,
		10,
		13,
		17,
		21,
		26,
		31,
		37,
		43,
		50,
		57,
		65,
		73,
		82,
		91,
		101,
		111,
		122,
		133,
		145,
		157,
		170,
		183,
		197,
		211,
		226,
		241,
		257,
		273,
		290,
		307,
	}
	if ability < 1 || ability > 45 {
		return 0, false
	} else {
		return ability_cost[ability-1], true
	}
}

// (define (cost-abilities ab)
//   (apply + (map ability->cost ab)))
func AbilityModifier(ability int) (int, bool) {
	ability_modifier := []int{
		-5,
		-4,
		-4,
		-3,
		-3,
		-2,
		-2,
		-1,
		-1,
		0,
		0,
		1,
		1,
		2,
		2,
		3,
		3,
		4,
		4,
		5,
		5,
		6,
		6,
		7,
		7,
		8,
		8,
		9,
		9,
		10,
		10,
		11,
		11,
		12,
		12,
		13,
		13,
		14,
		14,
		15,
		15,
		16,
		16,
		17,
		17,
	}
	if ability < 1 || ability > 45 {
		return 0, false
	} else {
		return ability_modifier[ability-1], true
	}
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
