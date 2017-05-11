package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

func getParameters() (*uint32, *uint32, *string, *uint32, *uint32, *bool) {
	var (
		dieroller                     = kingpin.New("dieroller", "A roll playing game die roller.")
		diceFlag                      = dieroller.Flag("dice", "Number of dice to roll. Must be greater than 0.").Default("1").Short('d').Uint32()
		sidesFlag                     = dieroller.Flag("sides", "Number of sides per die. Must be greater than 0.").Default("20").Short('s').Uint32()
		modifierFlag                  = dieroller.Flag("modifier", "Modifier to the rolles.  The first character can optionally be one of +, -, or * followed by a number.  If the +, -, or * are missing, + is assumed.").Default("").Short('m').String()
		keepFlag                      = dieroller.Flag("keep", "Number of rolls to keep. Must be greater than 0 and less than or equal to number of dice.").Default("1").Short('k').Uint32()
		iterationsFlag                = dieroller.Flag("iterations", "Number of times to repeat the same rolls. Must be greater than 0.").Default("1").Short('i').Uint32()
		verbose                       = dieroller.Flag("verbose", "Display additional information.").Default("false").Short('v').Bool()
		diceArg                       = dieroller.Arg("dice", "Number of dice to roll. Must be greater than 0. (default to 1)").Uint32()
		sidesArg                      = dieroller.Arg("sides", "Number of sides per die. Must be greater than 0. (default to 20)").Uint32()
		modifierArg                   = dieroller.Arg("modifier", "Modifier to the rolles.  The first character can optionally be one of +, -, or * followed by a number.  If the +, -, or * are missing, + is assumed. (default to no modifier)").String()
		keepArg                       = dieroller.Arg("keep", "Number of rolls to keep. Must be greater than 0 and less than or equal to number of dice. (default to number of dice)").Uint32()
		iterationsArg                 = dieroller.Arg("iterations", "Number of times to repeat the same rolls. Must be greater than 0. (default to 1)").Uint32()
		dice, sides, keep, iterations *uint32
		modifier                      *string
	)
	kingpin.MustParse(dieroller.Parse(os.Args[1:]))
	if *diceArg == uint32(0) {
		dice = diceFlag
	} else {
		dice = diceArg
	}
	if *sidesArg == uint32(0) {
		sides = sidesFlag
	} else {
		sides = sidesArg
	}
	if *keepArg == uint32(0) {
		keep = keepFlag
	} else {
		keep = keepArg
	}
	if *modifierArg == "" {
		modifier = modifierFlag
	} else {
		modifier = modifierArg
	}
	if *iterationsArg == uint32(0) {
		iterations = iterationsFlag
	} else {
		iterations = iterationsArg
	}
	if *keep > *dice {
		dieroller.FatalUsage("\nError: keep (%v) must be <= dice (%v)\n", *keep, *dice)
	}

	return dice, sides, modifier, keep, iterations, verbose
}

func main() {
	var dice, sides, modifier, keep, iterations, verbose = getParameters()
	fmt.Printf("dice = %v\n", *dice)
	fmt.Printf("sides = %v\n", *sides)
	fmt.Printf("modifier = %v\n", *modifier)
	fmt.Printf("keep = %v\n", *keep)
	fmt.Printf("iterations = %v\n", *iterations)
	fmt.Printf("verbose = %v\n", *verbose)
}
