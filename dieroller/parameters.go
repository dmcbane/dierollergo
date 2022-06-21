package main

import (
	"flag"
	"fmt"
	"os"
)

func GetParameters() (int, int, string, int, int, bool) {
	const (
		diceDesc    = "Number of dice to roll. Must be greater than 0."
		sidesDesc   = "Number of sides per die. Must be greater than 0."
		modDesc     = "Modifier to the rolls.  The first character can optionally be one of 'a', 'm', or 's' followed by a number. These indicate addition, multiplication, and subtraction respectively. If the a, m, or s are missing, addition is assumed."
		keepDesc    = "Number of rolls to keep. Must be greater than 0 and less than or equal to number of dice."
		iterDesc    = "Number of times to repeat the same rolls. Must be greater than 0."
		verboseDesc = "Display additional information."
	)
	var dice, sides, keep, iterations int
	var modifier string
	var verbose bool
	flag.IntVar(&dice, "dice", 1, diceDesc)
	flag.IntVar(&dice, "d", 1, diceDesc)
	flag.IntVar(&sides, "sides", 20, sidesDesc)
	flag.IntVar(&sides, "s", 20, sidesDesc)
	flag.StringVar(&modifier, "modifier", "", modDesc)
	flag.StringVar(&modifier, "m", "", modDesc)
	flag.IntVar(&keep, "keep", dice, keepDesc)
	flag.IntVar(&keep, "k", dice, keepDesc)
	flag.IntVar(&iterations, "iterations", 1, iterDesc)
	flag.IntVar(&iterations, "i", 1, iterDesc)
	flag.BoolVar(&verbose, "verbose", false, verboseDesc)
	flag.BoolVar(&verbose, "v", false, verboseDesc)
	flag.Parse()
	if keep > dice {
		// panic("The number of dice to keep cannot exceed the number of dice rolled.")
		fmt.Println(fmt.Errorf("the number of dice to keep (%v) cannot exceed the number of dice rolled (%v)", keep, dice))
		os.Exit(1)
	}
	return dice, sides, modifier, keep, iterations, verbose
}
