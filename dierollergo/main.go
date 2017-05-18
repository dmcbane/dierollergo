package main

import (
	"fmt"
	"github.com/dmcbane/dierollergo/pkg"
)

func main() {
	var dice, sides, modifier, keep, iterations, verbose = GetParameters()

	dieroll := dieroller.NewDieRoll(*dice, *sides, *modifier, *keep)

	for i := uint32(0); i < *iterations; i++ {
		dieroll.Roll()
		if *verbose {
			fmt.Printf("%s => %d\n", dieroll.StandardStringVerbose(), dieroll.LastRoll().Result)
		} else {
			fmt.Printf("%d\n", dieroll.LastRoll().Result)
		}
	}
}
