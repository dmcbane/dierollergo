package dierollerpkg

import (
	"testing"
)

func TestConstructorAndGetters(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping constructor test in short mode.")
	}
	var dr DieRoll
	dr = *NewDieRoll(10, 20, "", 5)
	if dr.Dice() != uint32(10) {
		t.Error("Expected dr.Dice() to be 10, got ", dr.Dice())
	}
	if dr.Sides() != uint32(20) {
		t.Error("Expected dr.Sides() to be 20, got ", dr.Sides())
	}
	if dr.Modifier().ModType != DieModifierTypeNull {
		t.Error("Expected dr.Modifier().ModType to be DieModifierTypeNull, got ", dr.Modifier().ModType)
	}
	if dr.Modifier().Amount != uint32(0) {
		t.Error("Expected dr.Modifier().Amount to be 0, got ", dr.Modifier().Amount)
	}
	if dr.Keep() != uint32(5) {
		t.Error("Expected dr.Keep() to be 5, got ", dr.Keep())
	}
	if len(dr.History()) != 0 {
		t.Error("Expected dr.History() to be empty, got ", dr.History())
	}
}

func TestRoll(t *testing.T) {
	dr := *NewDieRoll(10, 20, "", 5)
	r := dr.Roll()
	if r.Result < int64(5) || r.Result > int64(100) {
		t.Error("Expected r.Result to be between 5 and 100 inclusive, got ", r.Result)
	}
	if uint32(len(r.Rolls)) != dr.Keep() {
		t.Errorf("Expected r.Rolls to have %d entries, got ", dr.Keep(), len(r.Rolls))
	}
	if int64(r.Rolls[0]+r.Rolls[1]+r.Rolls[2]+r.Rolls[3]+r.Rolls[4]) != r.Result {
		t.Errorf("Expected r.Result to be the sum of %d,%d,%d,%d, & %d, got ", r.Rolls[0], r.Rolls[1], r.Rolls[2], r.Rolls[3], r.Rolls[4], r.Result)
	}
}

func dierollresultequals(a, b DieRollResult) bool {
	rollsequal := true
	for i, v := range a.Rolls {
		if v != b.Rolls[i] {
			rollsequal = false
			break
		}
	}
	return a.Result == b.Result && rollsequal
}

func TestLastRoll(t *testing.T) {
	dr := *NewDieRoll(10, 20, "", 5)
	if dr.LastRoll() != nil {
		t.Error("Expected the dr.LastRoll() to be empty before the first roll.")
	}
	r := dr.Roll()
	s := dr.LastRoll()
	if !dierollresultequals(r, *s) {
		t.Error("Expected the result of dr.Roll() and dr.LastRoll() to be the same, got ", r, s)
	}
}

func TestAddHistory(t *testing.T) {
	dr := *NewDieRoll(10, 20, "", 5)
	r := dr.Roll()
	dr.Roll()
	dr.Roll()
	dr.AddHistory(r)
	s := dr.LastRoll()
	if len(dr.History()) != 4 {
		t.Error("Expected dr.History() to contain 4 entries, got ", len(dr.History()))
	}
	if !dierollresultequals(r, *s) {
		t.Error("Expected the last entry in dr.History() to be the same as the first, got ", r, s)
	}
}

func TestDieRollString(t *testing.T) {
	dr := *NewDieRoll(10, 20, "", 5)
	s := "dice: 10 sides: 20 mod:  keep: 5 history: (0 entries: )"
	if dr.String() != s {
		t.Error("Expected dr.String() to be", s, ", got", dr.String())
	}
	dr.AddHistory(DieRollResult{15, DieRolls{1, 2, 3, 4, 5}})
	s = "dice: 10 sides: 20 mod:  keep: 5 history: (1 entries: {Result: 15 (Rolls: 1,2,3,4,5)})"
	if dr.String() != s {
		t.Error("Expected dr.String() to be", s, ", got", dr.String())
	}
}

func TestDieRollStandardString(t *testing.T) {
	dr := *NewDieRoll(10, 20, "", 5)
	s := "10D20K5"
	if dr.StandardString() != s {
		t.Error("Expected dr.StandardString() to be", s, ", got", dr.StandardString())
	}
	dr = *NewDieRoll(10, 20, "3", 5)
	s = "10D20K5+3"
	if dr.StandardString() != s {
		t.Error("Expected dr.StandardString() to be", s, ", got", dr.StandardString())
	}
	dr = *NewDieRoll(10, 20, "+3", 5)
	s = "10D20K5+3"
	if dr.StandardString() != s {
		t.Error("Expected dr.StandardString() to be", s, ", got", dr.StandardString())
	}
	dr = *NewDieRoll(10, 20, "a3", 5)
	s = "10D20K5+3"
	if dr.StandardString() != s {
		t.Error("Expected dr.StandardString() to be", s, ", got", dr.StandardString())
	}
	dr = *NewDieRoll(10, 20, "-3", 5)
	s = "10D20K5-3"
	if dr.StandardString() != s {
		t.Error("Expected dr.StandardString() to be", s, ", got", dr.StandardString())
	}
	dr = *NewDieRoll(10, 20, "s3", 5)
	s = "10D20K5-3"
	if dr.StandardString() != s {
		t.Error("Expected dr.StandardString() to be", s, ", got", dr.StandardString())
	}
	dr = *NewDieRoll(10, 20, "*3", 5)
	s = "10D20K5*3"
	if dr.StandardString() != s {
		t.Error("Expected dr.StandardString() to be", s, ", got", dr.StandardString())
	}
	dr = *NewDieRoll(10, 20, "m3", 5)
	s = "10D20K5*3"
	if dr.StandardString() != s {
		t.Error("Expected dr.StandardString() to be", s, ", got", dr.StandardString())
	}
}

func TestDieRollStandardStringVerboseVerbose(t *testing.T) {
	dr := *NewDieRoll(10, 20, "", 5)
	s := "10D20K5"
	if dr.StandardStringVerbose() != s {
		t.Error("Expected dr.StandardStringVerbose() to be", s, ", got", dr.StandardStringVerbose())
	}
	dr = *NewDieRoll(10, 20, "3", 5)
	s = "10D20K5+3"
	if dr.StandardStringVerbose() != s {
		t.Error("Expected dr.StandardStringVerbose() to be", s, ", got", dr.StandardStringVerbose())
	}
	dr = *NewDieRoll(10, 20, "+3", 5)
	s = "10D20K5+3"
	if dr.StandardStringVerbose() != s {
		t.Error("Expected dr.StandardStringVerbose() to be", s, ", got", dr.StandardStringVerbose())
	}
	dr = *NewDieRoll(10, 20, "a3", 5)
	s = "10D20K5+3"
	if dr.StandardStringVerbose() != s {
		t.Error("Expected dr.StandardStringVerbose() to be", s, ", got", dr.StandardStringVerbose())
	}
	dr = *NewDieRoll(10, 20, "-3", 5)
	s = "10D20K5-3"
	if dr.StandardStringVerbose() != s {
		t.Error("Expected dr.StandardStringVerbose() to be", s, ", got", dr.StandardStringVerbose())
	}
	dr = *NewDieRoll(10, 20, "s3", 5)
	s = "10D20K5-3"
	if dr.StandardStringVerbose() != s {
		t.Error("Expected dr.StandardStringVerbose() to be", s, ", got", dr.StandardStringVerbose())
	}
	dr = *NewDieRoll(10, 20, "*3", 5)
	s = "10D20K5*3"
	if dr.StandardStringVerbose() != s {
		t.Error("Expected dr.StandardStringVerbose() to be", s, ", got", dr.StandardStringVerbose())
	}
	dr = *NewDieRoll(10, 20, "m3", 5)
	s = "10D20K5*3"
	if dr.StandardStringVerbose() != s {
		t.Error("Expected dr.StandardStringVerbose() to be", s, ", got", dr.StandardStringVerbose())
	}

	dr = *NewDieRoll(10, 20, "", 5)
	dr.AddHistory(DieRollResult{15, DieRolls{1, 2, 3, 4, 5}})
	s = "10D20K5 (1,2,3,4,5)"
	if dr.StandardStringVerbose() != s {
		t.Error("Expected dr.StandardStringVerbose() to be", s, ", got", dr.StandardStringVerbose())
	}
	dr = *NewDieRoll(10, 20, "3", 5)
	dr.AddHistory(DieRollResult{15, DieRolls{1, 2, 3, 4, 5}})
	s = "10D20K5+3 (1,2,3,4,5)"
	if dr.StandardStringVerbose() != s {
		t.Error("Expected dr.StandardStringVerbose() to be", s, ", got", dr.StandardStringVerbose())
	}
	dr = *NewDieRoll(10, 20, "+3", 5)
	dr.AddHistory(DieRollResult{15, DieRolls{1, 2, 3, 4, 5}})
	s = "10D20K5+3 (1,2,3,4,5)"
	if dr.StandardStringVerbose() != s {
		t.Error("Expected dr.StandardStringVerbose() to be", s, ", got", dr.StandardStringVerbose())
	}
	dr = *NewDieRoll(10, 20, "a3", 5)
	dr.AddHistory(DieRollResult{15, DieRolls{1, 2, 3, 4, 5}})
	s = "10D20K5+3 (1,2,3,4,5)"
	if dr.StandardStringVerbose() != s {
		t.Error("Expected dr.StandardStringVerbose() to be", s, ", got", dr.StandardStringVerbose())
	}
	dr = *NewDieRoll(10, 20, "-3", 5)
	dr.AddHistory(DieRollResult{15, DieRolls{1, 2, 3, 4, 5}})
	s = "10D20K5-3 (1,2,3,4,5)"
	if dr.StandardStringVerbose() != s {
		t.Error("Expected dr.StandardStringVerbose() to be", s, ", got", dr.StandardStringVerbose())
	}
	dr = *NewDieRoll(10, 20, "s3", 5)
	dr.AddHistory(DieRollResult{15, DieRolls{1, 2, 3, 4, 5}})
	s = "10D20K5-3 (1,2,3,4,5)"
	if dr.StandardStringVerbose() != s {
		t.Error("Expected dr.StandardStringVerbose() to be", s, ", got", dr.StandardStringVerbose())
	}
	dr = *NewDieRoll(10, 20, "*3", 5)
	dr.AddHistory(DieRollResult{15, DieRolls{1, 2, 3, 4, 5}})
	s = "10D20K5*3 (1,2,3,4,5)"
	if dr.StandardStringVerbose() != s {
		t.Error("Expected dr.StandardStringVerbose() to be", s, ", got", dr.StandardStringVerbose())
	}
	dr = *NewDieRoll(10, 20, "m3", 5)
	dr.AddHistory(DieRollResult{15, DieRolls{1, 2, 3, 4, 5}})
	s = "10D20K5*3 (1,2,3,4,5)"
	if dr.StandardStringVerbose() != s {
		t.Error("Expected dr.StandardStringVerbose() to be", s, ", got", dr.StandardStringVerbose())
	}
}

func TestDieRollHistoryAsString(t *testing.T) {
	dr := *NewDieRoll(10, 20, "", 5)
	s := "0 entries: "
	if dr.HistoryAsString() != s {
		t.Error("Expected dr.HistoryAsString() to be", s, ", got", dr.HistoryAsString())
	}
	dr.AddHistory(DieRollResult{15, DieRolls{1, 2, 3, 4, 5}})
	s = "1 entries: {Result: 15 (Rolls: 1,2,3,4,5)}"
	if dr.HistoryAsString() != s {
		t.Error("Expected dr.HistoryAsString() to be", s, ", got", dr.HistoryAsString())
	}
}
