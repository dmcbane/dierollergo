package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"regexp"
)

func GetParameters() (*uint32, *uint32, *string, *uint32, *uint32, *bool) {
	var (
		dieroller                     = kingpin.New("dieroller", "A roll playing game die roller.")
		diceFlag                      = dieroller.Flag("dice", "Number of dice to roll. Must be greater than 0.").Default("1").Short('d').Uint32()
		sidesFlag                     = dieroller.Flag("sides", "Number of sides per die. Must be greater than 0.").Default("20").Short('s').Uint32()
		modifierFlag                  = dieroller.Flag("modifier", "Modifier to the rolls.  The first character can optionally be one of 'a', 'm', or 's' followed by a number. These indicate addition, multiplication, and subtraction respectively. If the a, m, or s are missing, addition is assumed.").Default("").Short('m').String()
		keepFlag                      = dieroller.Flag("keep", "Number of rolls to keep. Must be greater than 0 and less than or equal to number of dice.").Default("1").Short('k').Uint32()
		iterationsFlag                = dieroller.Flag("iterations", "Number of times to repeat the same rolls. Must be greater than 0.").Default("1").Short('i').Uint32()
		verbose                       = dieroller.Flag("verbose", "Display additional information.").Default("false").Short('v').Bool()
		diceArg                       = dieroller.Arg("dice", "Number of dice to roll. Must be greater than 0. (default to 1)").Uint32()
		sidesArg                      = dieroller.Arg("sides", "Number of sides per die. Must be greater than 0. (default to 20)").Uint32()
		modifierArg                   = dieroller.Arg("modifier", "Modifier to the rolls.  The first character can optionally be one of 'a', 'm', or 's' followed by a number. These indicate addition, multiplication, and subtraction respectively. If the a, m, or s are missing, addition is assumed. (default to no modifier)").String()
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
		dieroller.FatalUsage("keep (%v) must be <= dice (%v).\n", *keep, *dice)
	}
	match, _ := regexp.MatchString("\\A([amsAMS]?)(\\d+)\\z|\\A\\z", *modifier)
	if !match {
		dieroller.FatalUsage("modifier (%v) is invalid.\n", *modifier)
	}

	return dice, sides, modifier, keep, iterations, verbose
}
