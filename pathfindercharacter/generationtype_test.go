package main

import (
	"testing"
)

func TestIsValidPoolSpecification(t *testing.T) {
	// bad string so invalid
	valid := IsValidPoolSpecification("1")
	if valid {
		t.Errorf("Expected IsValidPoolSpecification('1') to be false, got %s", valid)
	}
	// good string
	valid = IsValidPoolSpecification("3/3/3/3/3/9")
	if !valid {
		t.Errorf("Expected IsValidPoolSpecification('3/3/3/3/3/9') to be true, got %s", valid)
	}
	// one die under 3 so invalid
	valid = IsValidPoolSpecification("2/3/3/3/3/9")
	if valid {
		t.Errorf("Expected IsValidPoolSpecification('2/3/3/3/3/9') to be false, got %s", valid)
	}
	// total over 24 so invalid
	valid = IsValidPoolSpecification("4/3/3/3/3/9")
	if valid {
		t.Errorf("Expected IsValidPoolSpecification('4/3/3/3/3/9') to be false, got %s", valid)
	}
	// not enough values so invalid
	valid = IsValidPoolSpecification("3/3/3/3/3")
	if valid {
		t.Errorf("Expected IsValidPoolSpecification('3/3/3/3/3') to be false, got %s", valid)
	}
	// good string
	valid = IsValidPoolSpecification("3/3/3/3/3/3")
	if !valid {
		t.Errorf("Expected IsValidPoolSpecification('3/3/3/3/3/3') to be true, got %s", valid)
	}
	// non numeric value so invalid
	valid = IsValidPoolSpecification("3/3/a/3/3/3")
	if valid {
		t.Errorf("Expected IsValidPoolSpecification('3/3/a/3/3/3') to be false, got %s", valid)
	}
}

func TestIsValidPurchaseSpecification(t *testing.T) {
	valid := IsValidPurchaseSpecification("1")
	if valid {
		t.Errorf("Expected IsValidPurchaseSpecification('1') to be false, got %s", valid)
	}
	valid = IsValidPurchaseSpecification("low")
	if !valid {
		t.Errorf("Expected IsValidPurchaseSpecification('low') to be true, got %s", valid)
	}
	valid = IsValidPurchaseSpecification("LOW")
	if !valid {
		t.Errorf("Expected IsValidPurchaseSpecification('LOW') to be true, got %s", valid)
	}
	valid = IsValidPurchaseSpecification("standard")
	if !valid {
		t.Errorf("Expected IsValidPurchaseSpecification('standard') to be true, got %s", valid)
	}
	valid = IsValidPurchaseSpecification("STANDARD")
	if !valid {
		t.Errorf("Expected IsValidPurchaseSpecification('STANDARD') to be true, got %s", valid)
	}
	valid = IsValidPurchaseSpecification("high")
	if !valid {
		t.Errorf("Expected IsValidPurchaseSpecification('high') to be true, got %s", valid)
	}
	valid = IsValidPurchaseSpecification("HIGH")
	if !valid {
		t.Errorf("Expected IsValidPurchaseSpecification('HIGH') to be true, got %s", valid)
	}
	valid = IsValidPurchaseSpecification("epic")
	if !valid {
		t.Errorf("Expected IsValidPurchaseSpecification('epic') to be true, got %s", valid)
	}
	valid = IsValidPurchaseSpecification("EPIC")
	if !valid {
		t.Errorf("Expected IsValidPurchaseSpecification('EPIC') to be true, got %s", valid)
	}
}

func equal(x, y []int) bool {
	if len(x) != len(y) {
		return false
	} else {
		for i, v := range x {
			if y[i] != v {
				return false
			}
		}
		return true
	}
}

func TestParsePoolSpecification(t *testing.T) {
	// bad string so invalid
	got := ParsePoolSpecification("1")
	expected := []int{0, 0, 0, 0, 0, 0}
	if !equal(*got, expected) {
		t.Errorf("Expected ParsePoolSpecification('1') to be %s, got %s", expected, got)
	}
	// good string
	got = ParsePoolSpecification("3/3/3/3/3/9")
	expected = []int{3, 3, 3, 3, 3, 9}
	if !equal(*got, expected) {
		t.Errorf("Expected ParsePoolSpecification('3/3/3/3/3/9') to be %s, got %s", expected, got)
	}
	// one die under 3 so invalid
	got = ParsePoolSpecification("2/3/3/3/3/9")
	expected = []int{0, 0, 0, 0, 0, 0}
	if !equal(*got, expected) {
		t.Errorf("Expected ParsePoolSpecification('2/3/3/3/3/9') to be %s, got %s", expected, got)
	}
	// total over 24 so invalid
	got = ParsePoolSpecification("4/3/3/3/3/9")
	expected = []int{0, 0, 0, 0, 0, 0}
	if !equal(*got, expected) {
		t.Errorf("Expected ParsePoolSpecification('4/3/3/3/3/9') to be %s, got %s", expected, got)
	}
	// not enough values so invalid
	got = ParsePoolSpecification("3/3/3/3/3")
	expected = []int{0, 0, 0, 0, 0, 0}
	if !equal(*got, expected) {
		t.Errorf("Expected ParsePoolSpecification('3/3/3/3/3') to be %s, got %s", expected, got)
	}
	// good string
	got = ParsePoolSpecification("3/3/3/3/3/3")
	expected = []int{3, 3, 3, 3, 3, 3}
	if !equal(*got, expected) {
		t.Errorf("Expected ParsePoolSpecification('3/3/3/3/3/3') to be %s, got %s", expected, got)
	}
	// non numeric value so invalid
	got = ParsePoolSpecification("3/3/a/3/3/3")
	expected = []int{0, 0, 0, 0, 0, 0}
	if !equal(*got, expected) {
		t.Errorf("Expected ParsePoolSpecification('3/3/a/3/3/3') to be %s, got %s", expected, got)
	}
}
