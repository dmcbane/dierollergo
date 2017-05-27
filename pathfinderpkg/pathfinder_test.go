package pathfinderpkg

import (
	"math/rand"
	"reflect"
	"runtime"
	"testing"
)

func TestAbilityCost(t *testing.T) {
	// Costs outside of range
	if value, ok := AbilityCost(0); value != 0 || ok {
		t.Errorf("Expected AbilityCost(0) to yield 0/false, got %v/%v", value, ok)
	}
	if value, ok := AbilityCost(46); value != 0 || ok {
		t.Errorf("Expected AbilityCost(46) to yield 0/false, got %v/%v", value, ok)
	}
	for i := 1; i < 46; i++ {
		if value, ok := AbilityCost(i); !ok {
			t.Errorf("Expected AbilityCost(%v) to yield a valid value and true, got %v/%v", i, value, ok)
		}
	}
}

func TestAbilityModifier(t *testing.T) {
	// Modifier outside of range
	if value, ok := AbilityModifier(0); value != 0 || ok {
		t.Errorf("Expected AbilityModifier(0) to yield 0/false, got %v/%v", value, ok)
	}
	if value, ok := AbilityModifier(46); value != 0 || ok {
		t.Errorf("Expected AbilityModifier(46) to yield 0/false, got %v/%v", value, ok)
	}
	for i := 1; i < 46; i++ {
		if value, ok := AbilityModifier(i); !ok {
			t.Errorf("Expected AbilityModifier(%v) to yield a valid value and true, got %v/%v", i, value, ok)
		}
	}
}

func testAbilityFunctionOverRange(t *testing.T, hi, lo int, fn1 func(int) (int, bool), fn2 func(Abilities) int) {
	for i1 := lo; i1 < hi; i1++ {
		for i2 := lo; i2 < hi; i2++ {
			for i3 := lo; i3 < hi; i3++ {
				for i4 := lo; i4 < hi; i4++ {
					for i5 := lo; i5 < hi; i5++ {
						for i6 := 1; i6 < 46; i6++ {
							abilities := Abilities{i1, i2, i3, i4, i5, i6}
							expected := abilities.Map(func(x int) int { v, _ := fn1(x); return v }).Apply(0, func(x, y int) int { return x + y })
							got := fn2(abilities)
							if expected != got {
								t.Errorf("Expected %v(%v) to yield %v, got %v", runtime.FuncForPC(reflect.ValueOf(fn2).Pointer()).Name(), abilities, expected, got)
							}
						}
					}
				}
			}
		}
	}
}

func TestSumCostOfAbilitiesOutOfRange(t *testing.T) {
	// ability scores outside of range
	abilities := Abilities{0, 0, 0, 0, 0, 0}
	got := SumCostOfAbilities(abilities)
	if got != 0 {
		t.Errorf("Expected SumCostOfAbilities(%v) to be 0, got %v", abilities, got)
	}
	abilities = Abilities{46, 46, 46, 46, 46, 46}
	got = SumCostOfAbilities(abilities)
	if got != 0 {
		t.Errorf("Expected SumCostOfAbilities(%v) to be 0, got %v", abilities, got)
	}
}

func TestSumCostOfAbilities(t *testing.T) {
	max := 46
	rng := 10
	var hi, lo int
	if testing.Short() {
		lo = rand.Intn(max - rng)
		hi = lo + rng
		testAbilityFunctionOverRange(t, hi, lo, AbilityCost, SumCostOfAbilities)
	} else {
		t.Run("Full", func(t *testing.T) {
			t.Run("1", func(t *testing.T) { testAbilityFunctionOverRange(t, 24, 1, AbilityCost, SumCostOfAbilities) })
			t.Run("2", func(t *testing.T) { testAbilityFunctionOverRange(t, 46, 24, AbilityCost, SumCostOfAbilities) })
		})
	}
}

func TestSumModifiersFromAbilitiesOutOfRange(t *testing.T) {
	// ability scores outside of range
	abilities := Abilities{0, 0, 0, 0, 0, 0}
	got := SumModifiersFromAbilities(abilities)
	if got != 0 {
		t.Errorf("Expected SumModifiersFromAbilities(%v) to be 0, got %v", abilities, got)
	}
	abilities = Abilities{46, 46, 46, 46, 46, 46}
	got = SumModifiersFromAbilities(abilities)
	if got != 0 {
		t.Errorf("Expected SumModifiersFromAbilities(%v) to be 0, got %v", abilities, got)
	}
}

func TestSumModifiersFromAbilities(t *testing.T) {
	max := 46
	rng := 10
	var hi, lo int
	if testing.Short() {
		lo = rand.Intn(max - rng)
		hi = lo + rng
		testAbilityFunctionOverRange(t, hi, lo, AbilityModifier, SumModifiersFromAbilities)
	} else {
		t.Run("Full", func(t *testing.T) {
			t.Run("1", func(t *testing.T) { testAbilityFunctionOverRange(t, 24, 1, AbilityModifier, SumModifiersFromAbilities) })
			t.Run("2", func(t *testing.T) {
				testAbilityFunctionOverRange(t, 46, 24, AbilityModifier, SumModifiersFromAbilities)
			})
		})
	}
}

func BenchmarkSumCostOfAbilities(b *testing.B) {
	abilities := Abilities{15, 14, 14, 14, 13, 12}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SumCostOfAbilities(abilities)
	}
}

func BenchmarkSumModifierFromAbilities(b *testing.B) {
	abilities := Abilities{15, 14, 14, 14, 13, 12}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SumModifiersFromAbilities(abilities)
	}
}

func BenchmarkSumCostFunctional(b *testing.B) {
	abilities := Abilities{15, 14, 14, 14, 13, 12}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		abilities.Map(func(x int) int { v, _ := AbilityCost(x); return v }).Apply(0, func(x, y int) int { return x + y })
	}
}

func BenchmarkSumModifierFunctional(b *testing.B) {
	abilities := Abilities{15, 14, 14, 14, 13, 12}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		abilities.Map(func(x int) int { v, _ := AbilityModifier(x); return v }).Apply(0, func(x, y int) int { return x + y })
	}
}
