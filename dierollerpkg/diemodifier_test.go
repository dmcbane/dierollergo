package dierollerpkg

import (
	"testing"
)

// type DieModifier struct {
// 	ModType DieModifierType
// 	Amount  uint32
// }
//
// func DieModifierParse(value string) *DieModifier {
// 	if value == "" {
// 		return new(DieModifier)
// 	} else {
// 		modtype := DmtFromString(value)
// 		re1, _ := regexp.Compile(DieModifierRegex)
// 		matches := re1.FindStringSubmatch(value)
// 		u, _ := strconv.ParseUint(matches[2], 10, 32)
// 		if u != 0 && modtype == DieModifierTypeNull {
// 			modtype = DieModifierTypeAdd
// 		}
// 		return &DieModifier{modtype, uint32(u)}
// 	}
// }
func TestDieModifierParse(t *testing.T) {
	dmp := DieModifierParse("3")
	st := DieModifier{DieModifierTypeAdd, uint32(3)}
	if *dmp != st {
		t.Errorf("Expected DieModifierParse('3') to be %s, got %s", st, *dmp)
	}
	dmp = DieModifierParse("+43")
	st = DieModifier{DieModifierTypeAdd, uint32(43)}
	if *dmp != st {
		t.Errorf("Expected DieModifierParse('+43') to be %s, got %s", st, *dmp)
	}
	dmp = DieModifierParse("a13")
	st = DieModifier{DieModifierTypeAdd, uint32(13)}
	if *dmp != st {
		t.Errorf("Expected DieModifierParse('a13') to be %s, got %s", st, *dmp)
	}
	dmp = DieModifierParse("*31")
	st = DieModifier{DieModifierTypeMultiply, uint32(31)}
	if *dmp != st {
		t.Errorf("Expected DieModifierParse('*31') to be %s, got %s", st, *dmp)
	}
	dmp = DieModifierParse("m34")
	st = DieModifier{DieModifierTypeMultiply, uint32(34)}
	if *dmp != st {
		t.Errorf("Expected DieModifierParse('m31') to be %s, got %s", st, *dmp)
	}
	dmp = DieModifierParse("-20")
	st = DieModifier{DieModifierTypeSubtract, uint32(20)}
	if *dmp != st {
		t.Errorf("Expected DieModifierParse('-20') to be %s, got %s", st, *dmp)
	}
	dmp = DieModifierParse("s2")
	st = DieModifier{DieModifierTypeSubtract, uint32(2)}
	if *dmp != st {
		t.Errorf("Expected DieModifierParse('s2') to be %s, got %s", st, *dmp)
	}
	dmp = DieModifierParse("/20")
	st = DieModifier{DieModifierTypeNull, uint32(0)}
	if *dmp != st {
		t.Errorf("Expected DieModifierParse('/20') to be %s, got %s", st, *dmp)
	}
	dmp = DieModifierParse("&20")
	st = DieModifier{DieModifierTypeNull, uint32(0)}
	if *dmp != st {
		t.Errorf("Expected DieModifierParse('&20') to be %s, got %s", st, *dmp)
	}
	dmp = DieModifierParse("")
	st = DieModifier{DieModifierTypeNull, uint32(0)}
	if *dmp != st {
		t.Errorf("Expected DieModifierParse('') to be %s, got %s", st, *dmp)
	}
}

//
// func (dm DieModifier) String() string {
// 	t := dm.ModType.String()
// 	if t == "" || dm.Amount == 0 {
// 		return ""
// 	} else {
// 		return fmt.Sprintf("%s%d", t, dm.Amount)
// 	}
// }

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
