package pathfinderpkg

type MapFunc func(int) int
type ApplyFunc func(x, y int) int

// apply a function to each ability returning the new Abilities
// with the result of the function call mapped to the position of the
// value passed to the function.  For example, when ab is {1, 2, 3}
// ab.Map(times2) returns a new Ability {2, 4, 6} and ab is unchanged
func (ab Abilities) Map(fn MapFunc) Abilities {
	result := make(Abilities, ab.Len())
	for i, v := range ab {
		result[i] = fn(v)
	}
	return result
}

// apply a function to each ability, changing Abilities in place
// with the result of the function call mapped to the position of the
// value passed to the function.  For exmple, when ab is {1, 2, 3}
// ab.Map!(times2) returns ab that has been changed to {2, 4, 6}
func (ab Abilities) MapInPlace(fn MapFunc) Abilities {
	for i, v := range ab {
		ab[i] = fn(v)
	}
	return ab
}

// apply a function to each ability and an accumulated value
// with the result being the value in the accumulator after function
// has been applied to all values. For example, when ab is {1, 2, 3}
//
// ab.Apply(0, sum) => 6
// ab.Apply(1, sum) => 7
// ab.Apply(10, sum) => 16
// ab.Apply(0, prod) => 0
// ab.Apply(1, prod) => 6
// ab.Apply(10, prod) => 60
func (ab Abilities) Apply(initialValue int, fn ApplyFunc) int {
	accumulator := initialValue
	for _, v := range ab {
		accumulator = fn(accumulator, v)
	}
	return accumulator
}
