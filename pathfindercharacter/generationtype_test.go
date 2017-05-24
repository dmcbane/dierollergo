package main

import (
	"testing"
)

func TestIsValidPoolSpecification(t *testing.T) {
	valid := IsValidPoolSpecification("1")
	if valid {
		t.Errorf("Expected IsValidPoolSpecification('1') to be false, got %s", valid)
	}
	valid = IsValidPoolSpecification("3/3/3/3/3/9")
	if !valid {
		t.Errorf("Expected IsValidPoolSpecification('3/3/3/3/3/9') to be true, got %s", valid)
	}
	valid = IsValidPoolSpecification("2/3/3/3/3/9")
	if valid {
		t.Errorf("Expected IsValidPoolSpecification('2/3/3/3/3/9') to be false, got %s", valid)
	}
	valid = IsValidPoolSpecification("4/3/3/3/3/9")
	if valid {
		t.Errorf("Expected IsValidPoolSpecification('4/3/3/3/3/9') to be false, got %s", valid)
	}
	valid = IsValidPoolSpecification("3/3/3/3/3")
	if valid {
		t.Errorf("Expected IsValidPoolSpecification('3/3/3/3/3') to be false, got %s", valid)
	}
	valid = IsValidPoolSpecification("3/3/3/3/3/3")
	if !valid {
		t.Errorf("Expected IsValidPoolSpecification('3/3/3/3/3/3') to be true, got %s", valid)
	}
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
