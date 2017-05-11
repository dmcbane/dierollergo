package main

import "fmt"

func main() {
	var dice, sides, modifier, keep, iterations, verbose = getParameters()
	fmt.Printf("dice = %v\n", *dice)
	fmt.Printf("sides = %v\n", *sides)
	fmt.Printf("modifier = %v\n", *modifier)
	fmt.Printf("keep = %v\n", *keep)
	fmt.Printf("iterations = %v\n", *iterations)
	fmt.Printf("verbose = %v\n", *verbose)
}
