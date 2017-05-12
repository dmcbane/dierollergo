package main

import (
	"fmt"
	"strings"
)

type DieRollHistory []DieRollResult

func (drh DieRollHistory) Append(value DieRollResult) {
	drh = append(drh, value)
}

func (drh DieRollHistory) String() string {
	var value []string
	for _, a := range drh {
		value = append(value, fmt.Sprintf("%v", a))
	}
	return strings.Join(value, ", ")
}
