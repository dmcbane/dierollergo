package main

import (
	"fmt"
	"strings"
)

type DieRolls []uint32

func (dr DieRolls) Len() int           { return len(dr) }
func (dr DieRolls) Swap(i, j int)      { dr[i], dr[j] = dr[j], dr[i] }
func (dr DieRolls) Less(i, j int) bool { return dr[i] > dr[j] } // reverse sorted

func (dr DieRolls) String() string {
	str := make([]string, len(dr))
	for i, v := range dr {
		str[i] = fmt.Sprint(v)
	}
	return strings.Join(str, ",")
}
