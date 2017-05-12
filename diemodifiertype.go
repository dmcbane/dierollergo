package main

type DieModifierType int

const (
	DieModifierTypeNull DieModifierType = iota
	DieModifierTypeAdd
	DieModifierTypeSubtract
	DieModifierTypeMultiply
)

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
