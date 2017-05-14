package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	// cost to buy ability score
	// 3 through 6 are not legal values
	// and are extrapolations mirroring the
	// top end of the scale for comparison
	// only
	ability_cost := map[int]int{
		1:  -25,
		2:  -20,
		3:  -16,
		4:  -12,
		5:  -9,
		6:  -6,
		7:  -4,
		10: 0,
		11: 1,
		12: 2,
		13: 3,
		14: 5,
		15: 7,
		16: 10,
		17: 13,
		18: 17,
		19: 21,
		20: 26,
		21: 31,
		22: 37,
		23: 43,
		24: 50,
		25: 57,
		26: 65,
		27: 73,
		28: 82,
		29: 91,
		30: 101,
		31: 111,
		32: 122,
		33: 133,
		34: 145,
		35: 157,
		36: 170,
		37: 183,
		38: 197,
		39: 211,
		40: 226,
		41: 241,
		42: 257,
		43: 273,
		44: 290,
		45: 307,
	}

	// (define (cost-abilities ab)
	//   (apply + (map ability->cost ab)))

	ability_modifier := map[int]int{
		1:  -5,
		2:  -4,
		3:  -4,
		4:  -3,
		5:  -3,
		6:  -2,
		7:  -2,
		8:  -1,
		9:  -1,
		10: 0,
		11: 0,
		12: 1,
		13: 1,
		14: 2,
		15: 2,
		16: 3,
		17: 3,
		18: 4,
		19: 4,
		20: 5,
		21: 5,
		22: 6,
		23: 6,
		24: 7,
		25: 7,
		26: 8,
		27: 8,
		28: 9,
		29: 9,
		30: 10,
		31: 10,
		32: 11,
		33: 11,
		34: 12,
		35: 12,
		36: 13,
		37: 13,
		38: 14,
		39: 14,
		40: 15,
		41: 15,
		42: 16,
		43: 16,
		44: 17,
		45: 17,
	}

	max := 45
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
							sq := fmt.Sprint(ability_cost[i] + ability_cost[j] + ability_cost[k] + ability_cost[m] + ability_cost[n] + ability_cost[p])
							sr := fmt.Sprint(ability_modifier[i] + ability_modifier[j] + ability_modifier[k] + ability_modifier[m] + ability_modifier[n] + ability_modifier[p])
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
