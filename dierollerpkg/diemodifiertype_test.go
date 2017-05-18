package dierollerpkg

import (
	"testing"
)

func TestDieModifierTypeString(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping string test in short mode.")
	}
	var dmt DieModifierType
	dmt = DieModifierTypeAdd
	result := dmt.String()
	s := "+"
	if result != s {
		t.Errorf("Expected dmt.String() to be +, got", result)
	}
	dmt = DieModifierTypeNull
	result = dmt.String()
	s = ""
	if result != s {
		t.Errorf("Expected dmt.String() to be '', got", result)
	}
	dmt = DieModifierTypeMultiply
	result = dmt.String()
	s = "*"
	if result != s {
		t.Errorf("Expected dmt.String() to be *, got", result)
	}
	dmt = DieModifierTypeSubtract
	result = dmt.String()
	s = "-"
	if result != s {
		t.Errorf("Expected dmt.String() to be -, got", result)
	}
}

func TestDieModifierTypeFromString(t *testing.T) {
	dmt := DmtFromString("3")
	if dmt != DieModifierTypeAdd {
		t.Error("Expected DmtFromString('3') to be DieModifierTypeAdd, got", dmt)
	}
	dmt = DmtFromString("+43")
	if dmt != DieModifierTypeAdd {
		t.Error("Expected DmtFromString('+43') to be DieModifierTypeAdd, got", dmt)
	}
	dmt = DmtFromString("a13")
	if dmt != DieModifierTypeAdd {
		t.Error("Expected DmtFromString('a13') to be DieModifierTypeAdd, got", dmt)
	}
	dmt = DmtFromString("*31")
	if dmt != DieModifierTypeMultiply {
		t.Error("Expected DmtFromString('*31') to be DieModifierTypeMultiply, got", dmt)
	}
	dmt = DmtFromString("m34")
	if dmt != DieModifierTypeMultiply {
		t.Error("Expected DmtFromString('m34') to be DieModifierTypeMultiply, got", dmt)
	}
	dmt = DmtFromString("-20")
	if dmt != DieModifierTypeSubtract {
		t.Error("Expected DmtFromString('-20') to be DieModifierTypeSubtract, got", dmt)
	}
	dmt = DmtFromString("s2")
	if dmt != DieModifierTypeSubtract {
		t.Error("Expected DmtFromString('s2') to be DieModifierTypeSubtract, got", dmt)
	}
	dmt = DmtFromString("/20")
	if dmt != DieModifierTypeNull {
		t.Error("Expected DmtFromString('/20') to be DieModifierTypeNull, got", dmt)
	}
	dmt = DmtFromString("&20")
	if dmt != DieModifierTypeNull {
		t.Error("Expected DmtFromString('&20') to be DieModifierTypeNull, got", dmt)
	}
	dmt = DmtFromString("")
	if dmt != DieModifierTypeNull {
		t.Error("Expected DmtFromString('') to be DieModifierTypeNull, got", dmt)
	}
}
