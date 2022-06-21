package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dmcbane/dierollergo/pathfinderpkg"
)

func GetParameters() (pathfinderpkg.GenerationType, string, int, bool) {
	const (
		classDesc = "The classic method: 3D6 per ability."
		standDesc = "The standard method: 4D6 keep high 3 per ability. (this is the default)"
		heroDesc  = "The heroic method: 2D6 plus 6 per ability."
		poolDesc  = "The pool method: 24D6 for all 6 abilities. The parameter specifies how many dice are assigned to each ability as follows: 3/3/3/3/3/9 with a minimum of 3 dice per ability."
		purcDesc  = "The purchase method: parameters are set according to cost. The parameter specifies the purchase type as follows: low, standard, high, and epic fantasy which provides 10, 15, 20, and 25 purchase points respectively."
		numbDesc  = "Number of characters to roll. Must be greater than 0."
		verbDesc  = "Display additional information."
	)
	var classic, standard, heroic, verbose bool
	var pool, purchase string
	var number int

	flag.BoolVar(&classic, "classic", false, classDesc)
	flag.BoolVar(&classic, "c", false, classDesc)
	flag.BoolVar(&standard, "standard", false, standDesc)
	flag.BoolVar(&standard, "s", false, standDesc)
	flag.BoolVar(&heroic, "heroic", false, heroDesc)
	flag.BoolVar(&heroic, "r", false, heroDesc)
	flag.StringVar(&pool, "pool", "", poolDesc)
	flag.StringVar(&pool, "l", "", poolDesc)
	flag.StringVar(&purchase, "purchase", "", purcDesc)
	flag.StringVar(&purchase, "p", "", purcDesc)
	flag.IntVar(&number, "number", 1, numbDesc)
	flag.IntVar(&number, "n", 1, numbDesc)
	flag.BoolVar(&verbose, "verbose", false, verbDesc)
	flag.BoolVar(&verbose, "v", false, verbDesc)

	flag.Parse()

	if (classic && standard) ||
		(classic && heroic) ||
		(classic && (pool != "")) ||
		(classic && (purchase != "")) ||
		(standard && heroic) ||
		(standard && (pool != "")) ||
		(standard && (purchase != "")) ||
		(heroic && (pool != "")) ||
		(heroic && (purchase != "")) ||
		((pool != "") && (purchase != "")) {
		fmt.Println(fmt.Errorf("defining multiple generation types is not allowed"))
		os.Exit(1)
	}
	if pool != "" && !pathfinderpkg.IsValidPoolSpecification(pool) {
		fmt.Println(fmt.Errorf("invalid pool specification"))
		os.Exit(1)
	}
	if purchase != "" && !pathfinderpkg.IsValidPurchaseSpecification(purchase) {
		fmt.Println(fmt.Errorf("invalid purchase type specification"))
		os.Exit(1)
	}

	var gentype pathfinderpkg.GenerationType
	var options string

	switch {
	case classic:
		gentype = pathfinderpkg.GenerationTypeClassic
		options = ""
	case standard:
		gentype = pathfinderpkg.GenerationTypeStandard
		options = ""
	case heroic:
		gentype = pathfinderpkg.GenerationTypeHeroic
		options = ""
	case pool != "":
		gentype = pathfinderpkg.GenerationTypePool
		options = pool
	case purchase != "":
		gentype = pathfinderpkg.GenerationTypePurchase
		options = purchase
	default:
		gentype = pathfinderpkg.GenerationTypeStandard
		options = ""
	}
	return gentype, options, number, verbose
}
