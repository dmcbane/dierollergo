package dierollerpkg

import (
	"fmt"
	"regexp"
	"strconv"
)

const DieModifierRegex = "\\A([amsAMS+*-]?)(\\d+)\\z|\\A\\z"

type DieModifier struct {
	ModType DieModifierType
	Amount  int
}

func DieModifierIsValid(value string) bool {
	re1, _ := regexp.Compile(DieModifierRegex)
	matches := re1.FindStringSubmatch(value)
	return len(matches) == 3
}

func DieModifierParse(value string) *DieModifier {
	if value == "" {
		return new(DieModifier)
	}
	modtype := DmtFromString(value)
	re1, _ := regexp.Compile(DieModifierRegex)
	matches := re1.FindStringSubmatch(value)
	if len(matches) != 3 {
		return new(DieModifier)
	}
	u, _ := strconv.Atoi(matches[2])
	if u != 0 && modtype == DieModifierTypeNull {
		modtype = DieModifierTypeAdd
	}
	return &DieModifier{modtype, int(u)}
}

func (dm DieModifier) String() string {
	if dm.ModType == DieModifierTypeNull && dm.Amount == 0 {
		return ""
	}
	return fmt.Sprintf("%v%v", dm.ModType, dm.Amount)
}
