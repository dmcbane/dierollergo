package main

import (
	"fmt"
	// 	"github.com/dmcbane/dierollergo/dierollerpkg"
)

func main() {
	gentype, options, number, verbose := GetParameters()
	for i := uint32(0); i < *number; i++ {
		fmt.Printf("%s - %s - %s\n", gentype, options, verbose)
	}

	// dieroll := dierollerpkg.NewDieRoll(*dice, *sides, *modifier, *keep)

	// for i := uint32(0); i < *iterations; i++ {
	// 	dieroll.Roll()
	// 	if *verbose {
	// 		fmt.Printf("%s => %d\n", dieroll.StandardStringVerbose(), dieroll.LastRoll().Result)
	// 	} else {
	// 		fmt.Printf("%d\n", dieroll.LastRoll().Result)
	// 	}
	// }
}
