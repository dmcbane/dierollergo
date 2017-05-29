package main

import (
	"fmt"
	"sync"
)

func main() {
	gentype, options, number, verbose := GetParameters()
	fn := gentype.GetGenerator(options)
	if verbose {
		fmt.Printf("gentype: %v\noptions: %v\n", gentype, options)
	}
	var wg sync.WaitGroup
	for i := 0; i < number; i++ {
		wg.Add(1)
		if verbose {
			go func() {
				defer wg.Done()
				cma := fn()
				fmt.Printf("%v c: %v m: %v\n", cma.Abils, cma.Cost, cma.Modifier)
			}()
		} else {
			go func() {
				defer wg.Done()
				fmt.Println(fn().Abils)
			}()
		}
	}
	wg.Wait()
}
