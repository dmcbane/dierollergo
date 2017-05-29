package main

import (
	"fmt"
	"github.com/dmcbane/dierollergo/dierollerpkg"
)

func main() {
	var dice, sides, modifier, keep, iterations, verbose = GetParameters()

	dieroll := dierollerpkg.NewDieRoll(dice, sides, modifier, keep)

	for i := 0; i < iterations; i++ {
		dieroll.Roll()
		if verbose {
			fmt.Printf("%s => %d\n", dieroll.StandardStringVerbose(), dieroll.LastRoll().Result)
		} else {
			fmt.Printf("%d\n", dieroll.LastRoll().Result)
		}
	}
}
