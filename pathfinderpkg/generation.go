package pathfinderpkg

import (
	"github.com/dmcbane/dierollergo/dierollerpkg"
	"regexp"
	"sort"
	"strconv"
)

type GenerationType int

const (
	GenerationTypeClassic GenerationType = iota
	GenerationTypeStandard
	GenerationTypeHeroic
	GenerationTypePool
	GenerationTypePurchase
)

const GenerationTypePoolRegex = "\\A(\\d{1})/(\\d{1})/(\\d{1})/(\\d{1})/(\\d{1})/(\\d{1})\\z"
const GenerationTypePurchaseRegex = "(?i)\\A(low|standard|high|epic)\\z"

func IsValidPoolSpecification(value string) bool {
	if match, _ := regexp.MatchString(GenerationTypePoolRegex, value); !match {
		return false
	} else {
		re1, _ := regexp.Compile(GenerationTypePoolRegex)
		matches := re1.FindStringSubmatch(value)
		total := 0
		for i, v := range matches {
			if i == 0 {
				continue
			}
			iv, err := strconv.Atoi(v)
			if iv < 3 || err != nil {
				return false
			}
			total += iv
		}
		if total > 24 {
			return false
		} else {
			return true
		}
	}
}

func IsValidPurchaseSpecification(value string) bool {
	match, _ := regexp.MatchString(GenerationTypePurchaseRegex, value)
	return match
}

func ParsePoolSpecification(value string) *[]int {
	result := make([]int, 6)
	if IsValidPoolSpecification(value) {
		re1, _ := regexp.Compile(GenerationTypePoolRegex)
		matches := re1.FindStringSubmatch(value)
		for i, v := range matches {
			if i == 0 {
				continue
			}
			result[i-1], _ = strconv.Atoi(v)
		}
	}
	return &result
}

func PurchasePointsFromCampaignType(campaign string) int {
	switch campaign {
	case "low":
		return 10
	case "standard":
		return 15
	case "high":
		return 20
	case "epic":
		return 25
	default:
		return 0
	}
}

func (gt GenerationType) String() string {
	switch gt {
	case GenerationTypeClassic:
		return "Classic"
	case GenerationTypeStandard:
		return "Standard"
	case GenerationTypeHeroic:
		return "Heroic"
	case GenerationTypePool:
		return "Pool"
	case GenerationTypePurchase:
		return "Purchase"
	default:
		return ""
	}
}

func (gt GenerationType) GetGenerator(options string) func() CostModifierAbilities {
	switch gt {
	case GenerationTypeClassic:
		return getClassicGenerator()
	case GenerationTypeStandard:
		return getStandardGenerator()
	case GenerationTypeHeroic:
		return getHeroicGenerator()
	case GenerationTypePool:
		return getPoolGenerator(options)
	case GenerationTypePurchase:
		return getPurchaseGenerator(options)
	default:
		return getStandardGenerator()
	}
}

func getClassicGenerator() func() CostModifierAbilities {
	dr := dierollerpkg.NewDieRoll(3, 6, "", 3)
	return func() CostModifierAbilities {
		abils := Abilities{dr.Roll().Result, dr.Roll().Result, dr.Roll().Result, dr.Roll().Result, dr.Roll().Result, dr.Roll().Result}
		cost := abils.SumCosts()
		mod := abils.SumModifiers()
		return CostModifierAbilities{cost, mod, abils}
	}
}

func getStandardGenerator() func() CostModifierAbilities {
	dr := dierollerpkg.NewDieRoll(4, 6, "", 3)
	return func() CostModifierAbilities {
		abils := Abilities{dr.Roll().Result, dr.Roll().Result, dr.Roll().Result, dr.Roll().Result, dr.Roll().Result, dr.Roll().Result}
		cost := abils.SumCosts()
		mod := abils.SumModifiers()
		return CostModifierAbilities{cost, mod, abils}
	}
}

func getHeroicGenerator() func() CostModifierAbilities {
	dr := dierollerpkg.NewDieRoll(3, 6, "+6", 2)
	return func() CostModifierAbilities {
		abils := Abilities{dr.Roll().Result, dr.Roll().Result, dr.Roll().Result, dr.Roll().Result, dr.Roll().Result, dr.Roll().Result}
		cost := abils.SumCosts()
		mod := abils.SumModifiers()
		return CostModifierAbilities{cost, mod, abils}
	}
}

func getPoolGenerator(options string) func() CostModifierAbilities {
	a := ParsePoolSpecification(options)
	dr1 := dierollerpkg.NewDieRoll((*a)[0], 6, "", 3)
	dr2 := dierollerpkg.NewDieRoll((*a)[1], 6, "", 3)
	dr3 := dierollerpkg.NewDieRoll((*a)[2], 6, "", 3)
	dr4 := dierollerpkg.NewDieRoll((*a)[3], 6, "", 3)
	dr5 := dierollerpkg.NewDieRoll((*a)[4], 6, "", 3)
	dr6 := dierollerpkg.NewDieRoll((*a)[5], 6, "", 3)
	return func() CostModifierAbilities {
		abils := Abilities{dr1.Roll().Result, dr2.Roll().Result, dr3.Roll().Result, dr4.Roll().Result, dr5.Roll().Result, dr6.Roll().Result}
		cost := abils.SumCosts()
		mod := abils.SumModifiers()
		return CostModifierAbilities{cost, mod, abils}
	}
}

func getPurchaseGenerator(options string) func() CostModifierAbilities {
	points := PurchasePointsFromCampaignType(options)
	// make the slice maximum size
	legalAbilities := make(CostModifierAbilitiesByModifier, 12^6)
	count := 0
	for i1 := 18; i1 >= 7; i1-- {
		for i2 := 18; i2 >= 7; i2-- {
			for i3 := 18; i3 >= 7; i3-- {
				for i4 := 18; i4 >= 7; i4-- {
					for i5 := 18; i5 >= 7; i5-- {
						for i6 := 18; i6 >= 7; i6-- {
							abils := Abilities{i1, i2, i3, i4, i5, i6}
							cost := abils.SumCosts()
							if cost <= points {
								legalAbilities[count] = CostModifierAbilities{cost, abils.SumModifiers(), abils}
								count++
							}
						}
					}
				}
			}
		}
	}
	// delete abilities slots not used
	for i, v := range legalAbilities {
		if v.Abils[0] == 0 {
			legalAbilities = legalAbilities[:i]
			break
		}
	}
	// sort largest modifier total to smallest
	sort.Sort(legalAbilities)
	// context variable for which values have been used
	index := -1
	return func() CostModifierAbilities {
		index++
		return legalAbilities[index]
	}
}
