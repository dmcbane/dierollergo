package main

import (
	"dieroller/dierollerpkg"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func getParameters() (int, int, string, int, int, bool) {
	const (
		diceDesc    = "The `quantity` of dice to roll. Must be greater than 0."
		sidesDesc   = "The `number` of sides per die. Must be greater than 0."
		modDesc     = "A `modifier` to be applied to the rolls.  The first character can optionally be one of 'a', 'm', or 's' followed by a number. These indicate addition, multiplication, and subtraction respectively. If the a, m, or s are missing, addition is assumed."
		keepDesc    = "The `quantity` of rolls to keep. Must be greater than 0 and less than or equal to number of dice."
		iterDesc    = "The `number` of times to repeat the same rolls. Must be greater than 0."
		verboseDesc = "Display additional information."
		helpDesc    = "Display help on command line options."
	)

	var (
		dice, sides, keep, iterations int
		modifier                      string
		verbose, help                 bool
	)

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Dieroller — digital dice for role playing gamers.\n\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s <options>\n\n  where the options can be defined by position\n\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "  [dice [sides [modifier [keep]]]]\n\n  or flags\n\n")
		flag.PrintDefaults()
		fmt.Fprintf(flag.CommandLine.Output(), "\nExamples:\n  %s 5\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "  %s 1 10\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "  %s 3 6 +3\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "  %s --dice 5 --sides 100 --modifier +4 --keep 3\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "  %s --dice 4 --sides 6 --keep 3\n", os.Args[0])
	}

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
	flag.BoolVar(&help, "help", false, helpDesc)
	flag.BoolVar(&help, "h", false, helpDesc)

	flag.Parse()

	if flag.Arg(0) != "" {
		if i, err := strconv.Atoi(flag.Arg(0)); err == nil {
			dice = i
			keepNotDefault := false
			flag.Visit(func(f *flag.Flag) {
				if f.Name == "keep" || f.Name == "k" {
					keepNotDefault = true
				}
			})
			if !keepNotDefault {
				keep = i
			}
		}
	}
	if flag.Arg(1) != "" {
		if i, err := strconv.Atoi(flag.Arg(1)); err == nil {
			sides = i
		}
	}
	if flag.Arg(2) != "" {
		modifier = flag.Arg(2)
	}
	if flag.Arg(3) != "" {
		if i, err := strconv.Atoi(flag.Arg(3)); err == nil {
			keep = i
		}
	}

	if help {
		flag.Usage()
		os.Exit(0)
	}
	if keep > dice {
		fmt.Println(fmt.Errorf("the number of dice to keep (%v) cannot exceed the number of dice rolled (%v)", keep, dice))
		os.Exit(1)
	}
	if modifier != "" && !dierollerpkg.DieModifierIsValid(modifier) {
		fmt.Println(fmt.Errorf("the modifier is invalid (%s)", modifier))
		os.Exit(1)
	}
	return dice, sides, modifier, keep, iterations, verbose
}
