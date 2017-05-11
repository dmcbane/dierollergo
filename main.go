package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	dice          = kingpin.Flag("dice", "Number of dice to roll. Must be greater than 0.").Default("1").Short('d').Uint32()
	sides         = kingpin.Flag("sides", "Number of sides per die. Must be greater than 0.").Default("20").Short('s').Uint32()
	modifier      = kingpin.Flag("modifier", "Modifier to the rolles.  The first character can optionally be one of +, -, or * followed by a number.  If the +, -, or * are missing, + is assumed.").Default("").Short('m').String()
	keep          = kingpin.Flag("keep", "Number of rolls to keep. Must be greater than 0 and less than or equal to number of dice.").Default("1").Short('k').Uint32()
	iterations    = kingpin.Flag("iterations", "Number of times to repeat the same rolls. Must be greater than 0.").Default("1").Short('i').Uint32()
	verbose       = kingpin.Flag("verbose", "Display additional information.").Default("false").Short('v').Bool()
	diceArg       = kingpin.Arg("dice", "Number of dice to roll. Must be greater than 0. (default to 1)").Uint32()
	sidesArg      = kingpin.Arg("sides", "Number of sides per die. Must be greater than 0. (default to 20)").Uint32()
	modifierArg   = kingpin.Arg("modifier", "Modifier to the rolles.  The first character can optionally be one of +, -, or * followed by a number.  If the +, -, or * are missing, + is assumed. (default to no modifier)").String()
	keepArg       = kingpin.Arg("keep", "Number of rolls to keep. Must be greater than 0 and less than or equal to number of dice. (default to number of dice)").Uint32()
	iterationsArg = kingpin.Arg("iterations", "Number of times to repeat the same rolls. Must be greater than 0. (default to 1)").Uint32()
)

func main() {
	kingpin.Parse()
	fmt.Printf("dice = %v\n", *dice)
	fmt.Printf("diceArg = %v\n", *diceArg)
	fmt.Printf("sides = %v\n", *sides)
	fmt.Printf("sidesArg = %v\n", *sidesArg)
	fmt.Printf("modifier = %v\n", *modifier)
	fmt.Printf("modifierArg = %v\n", *modifierArg)
	fmt.Printf("keep = %v\n", *keep)
	fmt.Printf("keepArg = %v\n", *keepArg)
	fmt.Printf("iterations = %v\n", *iterations)
	fmt.Printf("iterationsArg = %v\n", *iterationsArg)
	fmt.Printf("verbose = %v\n", *verbose)
}
