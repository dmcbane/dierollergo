package dierollerpkg

import (
	"testing"
)

func TestLength(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping length test in short mode.")
	}
	var dr DieRolls
	var ln int
	dr = make(DieRolls, 0)
	ln = dr.Len()
	if ln != 0 {
		t.Error("Expected 0, got ", ln)
	}
	dr = make(DieRolls, 1)
	ln = dr.Len()
	if ln != 1 {
		t.Error("Expected 1, got ", ln)
	}
	dr = make(DieRolls, 10)
	ln = dr.Len()
	if ln != 10 {
		t.Error("Expected 10, got ", ln)
	}
}

func TestSwap(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping length test in short mode.")
	}
	dr := DieRolls{uint32(1), uint32(2), uint32(3)}
	dr.Swap(0, 1)
	if dr[0] != uint32(2) {
		t.Error("Expected dr[0] to be 2, got ", dr[0])
	}
	if dr[1] != uint32(1) {
		t.Error("Expected dr[1] to be 1, got ", dr[1])
	}
	dr.Swap(1, 2)
	if dr[1] != uint32(3) {
		t.Error("Expected dr[1] to be 3, got ", dr[1])
	}
	if dr[2] != uint32(1) {
		t.Error("Expected dr[2] to be 1, got ", dr[2])
	}
}

func TestLess(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping less test in short mode.")
	}
	dr := DieRolls{uint32(1), uint32(2), uint32(3)}
	// note that Less is actually implemented as Greater to provide a reverse sort
	result := dr.Less(0, 1)
	if result {
		t.Error("Expected dr.Less(0,1) to be false, got ", result)
	}
	result = dr.Less(1, 0)
	if !result {
		t.Error("Expected dr.Less(1,0) to be true, got ", result)
	}
}

func TestDieRollsString(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping string test in short mode.")
	}
	dr := DieRolls{uint32(1), uint32(2), uint32(3)}
	// note that Less is actually implemented as Greater to provide a reverse sort
	result := dr.String()
	if result != "1,2,3" {
		t.Error("Expected dr.String() to be 1,2,3, got ", result)
	}
}
