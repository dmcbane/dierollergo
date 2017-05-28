package main

import (
	"encoding/csv"
	"fmt"
	"github.com/dmcbane/dierollergo/pathfinderpkg"
	"log"
	"os"
)

func main() {
	max := 46
	top := max - 1
	arr := make([]string, 8)
	idx := 0
	f, err := os.Create("C:\\Users\\dmcbane\\GoogleDrive\\dev\\go\\src\\github.com\\dmcbane\\dierollergo\\all_ability_scores\\out.csv")
	if err != nil {
		log.Fatalln("error writing record to csv:", err)
		panic(err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	for i := top; i >= 0; i-- {
		for j := top; j >= 0; j-- {
			for k := top; k >= 0; k-- {
				for m := top; m >= 0; m-- {
					for n := top; n >= 0; n-- {
						for p := top; p >= 0; p-- {
							si, sj, sk, sm, sn, sp := fmt.Sprint(i), fmt.Sprint(j), fmt.Sprint(k), fmt.Sprint(m), fmt.Sprint(n), fmt.Sprint(p)
							aci, _ := pathfinderpkg.AbilityCost(i)
							acj, _ := pathfinderpkg.AbilityCost(j)
							ack, _ := pathfinderpkg.AbilityCost(k)
							acm, _ := pathfinderpkg.AbilityCost(m)
							acn, _ := pathfinderpkg.AbilityCost(n)
							acp, _ := pathfinderpkg.AbilityCost(p)
							sq := fmt.Sprint(aci + acj + ack + acm + acn + acp)
							ami, _ := pathfinderpkg.AbilityModifier(i)
							amj, _ := pathfinderpkg.AbilityModifier(j)
							amk, _ := pathfinderpkg.AbilityModifier(k)
							amm, _ := pathfinderpkg.AbilityModifier(m)
							amn, _ := pathfinderpkg.AbilityModifier(n)
							amp, _ := pathfinderpkg.AbilityModifier(p)
							sr := fmt.Sprint(ami + amj + amk + amm + amn + amp)
							arr = []string{si, sj, sk, sm, sn, sp, sq, sr}
							idx++

							if err := w.Write(arr); err != nil {
								log.Fatalln("error writing record to csv:", err)
							}
						}
					}
				}
			}
		}
	}
	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatalln(err)
	}
}
