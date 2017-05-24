package main

import (
	// "github.com/dmcbane/dierollergo/dierollerpkg"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	// "regexp"
)

func GetParameters() (*GenerationType, *string, *uint32, *bool) {
	var (
		pathfinder   = kingpin.New("pathfinder", "Generate Pathfinder character abilities.")
		classic      = pathfinder.Flag("classic", "The classic method: 3D6 per ability.").Default("false").Short('c').Bool()
		standard     = pathfinder.Flag("standard", "The standard method: 4D6 keep high 3 per ability. (this is the default)").Default("false").Short('s').Bool()
		heroic       = pathfinder.Flag("heroic", "The heroic method: 2D6 plus 6 per ability.").Default("false").Short('r').Bool()
		pool         = pathfinder.Flag("pool", "The pool method: 24D6 for all 6 abilities. The parameter specifies how many dice are assigned to each ability as follows: 3/3/3/3/3/9 with a minimum of 3 dice per ability.").Default("").Short('l').String()
		purchaseType = pathfinder.Flag("purchase", "The purchase method: parameters are set according to cost. The parameter specifies the purchase type as follows: low, standard, high, and epic fantasy which provides 10, 15, 20, and 25 purchase points respectively.").Default("").Short('p').String()
		number       = pathfinder.Flag("number", "Number of characters to roll. Must be greater than 0.").Default("1").Short('n').Uint32()
		verbose      = pathfinder.Flag("verbose", "Display additional information.").Default("false").Short('v').Bool()
	)
	kingpin.MustParse(pathfinder.Parse(os.Args[1:]))
	if (*classic && *standard) ||
		(*classic && *heroic) ||
		(*classic && (*pool != "")) ||
		(*classic && (*purchaseType != "")) ||
		(*standard && *heroic) ||
		(*standard && (*pool != "")) ||
		(*standard && (*purchaseType != "")) ||
		(*heroic && (*pool != "")) ||
		(*heroic && (*purchaseType != "")) ||
		((*pool != "") && (*purchaseType != "")) {
		pathfinder.FatalUsage("multiple generation types cannot be defined.\n")
	}
	// if *keep > *dice {
	// 	pathfinder.FatalUsage("keep (%v) must be <= dice (%v).\n", *keep, *dice)
	// }
	// match, _ := regexp.MatchString(pathfinderpkg.DieModifierRegex, *modifier)
	// if !match {
	// 	pathfinder.FatalUsage("modifier (%v) is invalid.\n", *modifier)
	// }

	test := GenerationTypeClassic
	return &test, nil, number, verbose
}
