package dierollerpkg

import (
	"testing"
)

func TestDieModifierParse(t *testing.T) {
	dmp := DieModifierParse("3")
	st := DieModifier{DieModifierTypeAdd, 3}
	if *dmp != st {
		t.Errorf("Expected DieModifierParse('3') to be %s, got %s", st, *dmp)
	}
	dmp = DieModifierParse("+43")
	st = DieModifier{DieModifierTypeAdd, 43}
	if *dmp != st {
		t.Errorf("Expected DieModifierParse('+43') to be %s, got %s", st, *dmp)
	}
	dmp = DieModifierParse("a13")
	st = DieModifier{DieModifierTypeAdd, 13}
	if *dmp != st {
		t.Errorf("Expected DieModifierParse('a13') to be %s, got %s", st, *dmp)
	}
	dmp = DieModifierParse("*31")
	st = DieModifier{DieModifierTypeMultiply, 31}
	if *dmp != st {
		t.Errorf("Expected DieModifierParse('*31') to be %s, got %s", st, *dmp)
	}
	dmp = DieModifierParse("m34")
	st = DieModifier{DieModifierTypeMultiply, 34}
	if *dmp != st {
		t.Errorf("Expected DieModifierParse('m31') to be %s, got %s", st, *dmp)
	}
	dmp = DieModifierParse("-20")
	st = DieModifier{DieModifierTypeSubtract, 20}
	if *dmp != st {
		t.Errorf("Expected DieModifierParse('-20') to be %s, got %s", st, *dmp)
	}
	dmp = DieModifierParse("s2")
	st = DieModifier{DieModifierTypeSubtract, 2}
	if *dmp != st {
		t.Errorf("Expected DieModifierParse('s2') to be %s, got %s", st, *dmp)
	}
	dmp = DieModifierParse("/20")
	st = DieModifier{DieModifierTypeNull, 0}
	if *dmp != st {
		t.Errorf("Expected DieModifierParse('/20') to be %s, got %s", st, *dmp)
	}
	dmp = DieModifierParse("&20")
	st = DieModifier{DieModifierTypeNull, 0}
	if *dmp != st {
		t.Errorf("Expected DieModifierParse('&20') to be %s, got %s", st, *dmp)
	}
	dmp = DieModifierParse("")
	st = DieModifier{DieModifierTypeNull, 0}
	if *dmp != st {
		t.Errorf("Expected DieModifierParse('') to be %s, got %s", st, *dmp)
	}
}

func TestDieModifierString(t *testing.T) {
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
