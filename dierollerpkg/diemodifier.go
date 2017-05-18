package dierollerpkg

import (
	"fmt"
	"regexp"
	"strconv"
)

const DieModifierRegex = "\\A([amsAMS+*-]?)(\\d+)\\z|\\A\\z"

type DieModifier struct {
	ModType DieModifierType
	Amount  uint32
}

func DieModifierParse(value string) *DieModifier {
	if value == "" {
		return new(DieModifier)
	} else {
		modtype := DmtFromString(value)
		re1, _ := regexp.Compile(DieModifierRegex)
		matches := re1.FindStringSubmatch(value)
		var u uint64
		if len(matches) != 3 {
			return new(DieModifier)
		}
		u, _ = strconv.ParseUint(matches[2], 10, 32)
		if u != 0 && modtype == DieModifierTypeNull {
			modtype = DieModifierTypeAdd
		}
		return &DieModifier{modtype, uint32(u)}
	}
}

func (dm DieModifier) String() string {
	if dm.ModType == DieModifierTypeNull && dm.Amount == 0 {
		return ""
	} else {
		return fmt.Sprintf("%v%v", dm.ModType, dm.Amount)
	}
}
