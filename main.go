package main

import "fmt"

func main() {
	var dice, sides, modifier, keep, iterations, verbose = GetParameters()

	if *verbose {
		fmt.Printf("dice = %v\n", *dice)
		fmt.Printf("sides = %v\n", *sides)
		fmt.Printf("modifier = %v\n", *modifier)
		fmt.Printf("keep = %v\n", *keep)
		fmt.Printf("iterations = %v\n", *iterations)
		fmt.Printf("verbose = %v\n", *verbose)
	}

	dieroll := DieRoll{*dice, *sides, *DieModifierParse(*modifier), *keep, DieRollHistory{}}

	for i := uint32(0); i < *iterations; i++ {
		fmt.Printf("%v\n", dieroll.Roll())
	}
}
