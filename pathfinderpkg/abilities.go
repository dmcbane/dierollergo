package pathfinderpkg

import (
	"sort"
)

// cost to buy ability score
// 1 through 6 and 19 through 45 are not
// legal values and are extrapolations
// for comparison only
func AbilityCost(ability int) (int, bool) {
	ability_cost := []int{
		-25, // 1
		-20, // 2
		-16, // 3
		-12, // 4
		-9,  // 5
		-6,  // 6
		-4,  // 7
		-2,  // 8
		-1,  // 9
		0,   // 10
		1,   // 11
		2,   // 12
		3,   // 13
		5,   // 14
		7,   // 15
		10,  // 16
		13,  // 17
		17,  // 18
		21,  // 19
		26,  // 20
		31,  // 21
		37,  // 22
		43,  // 23
		50,  // 24
		57,  // 25
		65,  // 26
		73,  // 27
		82,  // 28
		91,  // 29
		101, // 30
		111, // 31
		122, // 32
		133, // 33
		145, // 34
		157, // 35
		170, // 36
		183, // 37
		197, // 38
		211, // 39
		226, // 40
		241, // 41
		257, // 42
		273, // 43
		290, // 44
		307, // 45
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
		-5, // 1
		-4, // 2
		-4, // 3
		-3, // 4
		-3, // 5
		-2, // 6
		-2, // 7
		-1, // 8
		-1, // 9
		0,  // 10
		0,  // 11
		1,  // 12
		1,  // 13
		2,  // 14
		2,  // 15
		3,  // 16
		3,  // 17
		4,  // 18
		4,  // 19
		5,  // 20
		5,  // 21
		6,  // 22
		6,  // 23
		7,  // 24
		7,  // 25
		8,  // 26
		8,  // 27
		9,  // 28
		9,  // 29
		10, // 30
		10, // 31
		11, // 32
		11, // 33
		12, // 34
		12, // 35
		13, // 36
		13, // 37
		14, // 38
		14, // 39
		15, // 40
		15, // 41
		16, // 42
		16, // 43
		17, // 44
		17, // 45
	}
	if ability < 1 || ability > 45 {
		return 0, false
	} else {
		return ability_modifier[ability-1], true
	}
}

type Abilities []int

// implement sort.Interface
func (ab Abilities) Len() int           { return len(ab) }
func (ab Abilities) Less(i, j int) bool { return ab[i] < ab[j] }
func (ab Abilities) Swap(i, j int)      { ab[i], ab[j] = ab[j], ab[i] }

// for chaining...i.e. chained calls like in functional paradigm
func (ab Abilities) Sortf() Abilities { sort.Sort(ab); return ab }

func (ab Abilities) SumCosts() int {
	return SumCostOfAbilities(ab)
}

func (ab Abilities) SumModifiers() int {
	return SumModifiersFromAbilities(ab)
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

type CostModifierAbilities struct {
	Cost     int
	Modifier int
	Abils    Abilities
}

type CostModifierAbilitiesByCost []CostModifierAbilities

// implement sort.Interface
func (cmas CostModifierAbilitiesByCost) Len() int { return len(cmas) }

// decending sort on Cost
func (cmas CostModifierAbilitiesByCost) Less(i, j int) bool { return cmas[i].Cost > cmas[j].Cost }
func (cmas CostModifierAbilitiesByCost) Swap(i, j int)      { cmas[i], cmas[j] = cmas[j], cmas[i] }

type CostModifierAbilitiesByModifier []CostModifierAbilities

// implement sort.Interface
func (cmas CostModifierAbilitiesByModifier) Len() int { return len(cmas) }

// decending sort on Cost
func (cmas CostModifierAbilitiesByModifier) Less(i, j int) bool {
	return cmas[i].Modifier > cmas[j].Modifier
}
func (cmas CostModifierAbilitiesByModifier) Swap(i, j int) { cmas[i], cmas[j] = cmas[j], cmas[i] }
