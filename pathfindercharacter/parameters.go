package main

import (
	"github.com/dmcbane/dierollergo/pathfinderpkg"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

func GetParameters() (pathfinderpkg.GenerationType, string, int, bool) {
	var (
		pathfinder = kingpin.New("pathfinder", "Generate Pathfinder character abilities.")
		classic    = pathfinder.Flag("classic", "The classic method: 3D6 per ability.").Default("false").Short('c').Bool()
		standard   = pathfinder.Flag("standard", "The standard method: 4D6 keep high 3 per ability. (this is the default)").Default("false").Short('s').Bool()
		heroic     = pathfinder.Flag("heroic", "The heroic method: 2D6 plus 6 per ability.").Default("false").Short('r').Bool()
		pool       = pathfinder.Flag("pool", "The pool method: 24D6 for all 6 abilities. The parameter specifies how many dice are assigned to each ability as follows: 3/3/3/3/3/9 with a minimum of 3 dice per ability.").Default("").Short('l').String()
		purchase   = pathfinder.Flag("purchase", "The purchase method: parameters are set according to cost. The parameter specifies the purchase type as follows: low, standard, high, and epic fantasy which provides 10, 15, 20, and 25 purchase points respectively.").Default("").Short('p').String()
		number     = pathfinder.Flag("number", "Number of characters to roll. Must be greater than 0.").Default("1").Short('n').Int()
		verbose    = pathfinder.Flag("verbose", "Display additional information.").Default("false").Short('v').Bool()
	)
	kingpin.MustParse(pathfinder.Parse(os.Args[1:]))
	if (*classic && *standard) ||
		(*classic && *heroic) ||
		(*classic && (*pool != "")) ||
		(*classic && (*purchase != "")) ||
		(*standard && *heroic) ||
		(*standard && (*pool != "")) ||
		(*standard && (*purchase != "")) ||
		(*heroic && (*pool != "")) ||
		(*heroic && (*purchase != "")) ||
		((*pool != "") && (*purchase != "")) {
		pathfinder.FatalUsage("multiple generation types cannot be defined.\n")
	}
	if *pool != "" && !pathfinderpkg.IsValidPoolSpecification(*pool) {
		pathfinder.FatalUsage("invalid pool specification.\n")
	}
	if *purchase != "" && !pathfinderpkg.IsValidPurchaseSpecification(*purchase) {
		pathfinder.FatalUsage("invalid purchase type specification.\n")
	}

	var gentype pathfinderpkg.GenerationType
	var options string

	switch {
	case *classic:
		gentype = pathfinderpkg.GenerationTypeClassic
		options = ""
	case *standard:
		gentype = pathfinderpkg.GenerationTypeStandard
		options = ""
	case *heroic:
		gentype = pathfinderpkg.GenerationTypeHeroic
		options = ""
	case *pool != "":
		gentype = pathfinderpkg.GenerationTypePool
		options = *pool
	case *purchase != "":
		gentype = pathfinderpkg.GenerationTypePurchase
		options = *purchase
	default:
		gentype = pathfinderpkg.GenerationTypeStandard
		options = ""
	}
	return gentype, options, *number, *verbose
}
