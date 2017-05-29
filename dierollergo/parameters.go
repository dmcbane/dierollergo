package main

import (
	"github.com/dmcbane/dierollergo/dierollerpkg"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"regexp"
)

func GetParameters() (int, int, string, int, int, bool) {
	var (
		dieroller                     = kingpin.New("dieroller", "A roll playing game die roller.")
		diceFlag                      = dieroller.Flag("dice", "Number of dice to roll. Must be greater than 0.").Default("1").Short('d').Int()
		sidesFlag                     = dieroller.Flag("sides", "Number of sides per die. Must be greater than 0.").Default("20").Short('s').Int()
		modifierFlag                  = dieroller.Flag("modifier", "Modifier to the rolls.  The first character can optionally be one of 'a', 'm', or 's' followed by a number. These indicate addition, multiplication, and subtraction respectively. If the a, m, or s are missing, addition is assumed.").Default("").Short('m').String()
		keepFlag                      = dieroller.Flag("keep", "Number of rolls to keep. Must be greater than 0 and less than or equal to number of dice.").Short('k').Int()
		iterationsFlag                = dieroller.Flag("iterations", "Number of times to repeat the same rolls. Must be greater than 0.").Default("1").Short('i').Int()
		verbose                       = dieroller.Flag("verbose", "Display additional information.").Default("false").Short('v').Bool()
		diceArg                       = dieroller.Arg("dice", "Number of dice to roll. Must be greater than 0. (default to 1)").Int()
		sidesArg                      = dieroller.Arg("sides", "Number of sides per die. Must be greater than 0. (default to 20)").Int()
		modifierArg                   = dieroller.Arg("modifier", "Modifier to the rolls.  The first character can optionally be one of 'a', 'm', or 's' followed by a number. These indicate addition, multiplication, and subtraction respectively. If the a, m, or s are missing, addition is assumed. (default to no modifier)").String()
		keepArg                       = dieroller.Arg("keep", "Number of rolls to keep. Must be greater than 0 and less than or equal to number of dice. (default to number of dice)").Int()
		iterationsArg                 = dieroller.Arg("iterations", "Number of times to repeat the same rolls. Must be greater than 0. (default to 1)").Int()
		dice, sides, keep, iterations int
		modifier                      string
	)
	kingpin.MustParse(dieroller.Parse(os.Args[1:]))
	if *diceArg == 0 {
		dice = *diceFlag
	} else {
		dice = *diceArg
	}
	if *sidesArg == 0 {
		sides = *sidesFlag
	} else {
		sides = *sidesArg
	}
	if *keepArg == 0 {
		if *keepFlag == 0 {
			keep = dice
		} else {
			keep = *keepFlag
		}
	} else {
		keep = *keepArg
	}
	if *modifierArg == "" {
		modifier = *modifierFlag
	} else {
		modifier = *modifierArg
	}
	if *iterationsArg == 0 {
		iterations = *iterationsFlag
	} else {
		iterations = *iterationsArg
	}
	if keep > dice {
		dieroller.FatalUsage("keep (%v) must be <= dice (%v).\n", keep, dice)
	}
	match, _ := regexp.MatchString(dierollerpkg.DieModifierRegex, modifier)
	if !match {
		dieroller.FatalUsage("modifier (%v) is invalid.\n", modifier)
	}

	return dice, sides, modifier, keep, iterations, *verbose
}
