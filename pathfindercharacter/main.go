package main

import (
	"fmt"
)

func main() {
	gentype, options, number, verbose := GetParameters()
	fn := gentype.GetGenerator(options)
	if verbose {
		fmt.Printf("gentype: %v\noptions: %v\n", gentype, options)
	}

	for i := uint32(0); i < number; i++ {
		if verbose {
			cma := fn()
			fmt.Printf("%v c: %v m: %v\n", cma.Abils, cma.Cost, cma.Modifier)
		} else {
			fmt.Println(fn().Abils)
		}
	}
}
