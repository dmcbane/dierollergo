package dierollerpkg

import "regexp"

type DieModifierType int

const (
	DieModifierTypeNull DieModifierType = iota
	DieModifierTypeAdd
	DieModifierTypeSubtract
	DieModifierTypeMultiply
)

func DmtFromString(value string) DieModifierType {
	if match, _ := regexp.MatchString("\\A\\d+\\z", value); match {
		return DieModifierTypeAdd
	} else if value == "" {
		return DieModifierTypeNull
	} else {
		switch value[0] {
		case '+', 'a':
			return DieModifierTypeAdd
		case '*', 'm':
			return DieModifierTypeMultiply
		case '-', 's':
			return DieModifierTypeSubtract
		default:
			return DieModifierTypeNull
		}
	}
}

func (dmt DieModifierType) String() string {
	switch dmt {
	case DieModifierTypeAdd:
		return "+"
	case DieModifierTypeMultiply:
		return "*"
	case DieModifierTypeSubtract:
		return "-"
	default:
		return ""
	}
}
