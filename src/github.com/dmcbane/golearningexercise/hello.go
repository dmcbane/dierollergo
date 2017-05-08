package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var list [100]int64
	var sides = int64(20)
	rand.Seed(time.Now().UnixNano())
	for i := range list {
		list[i] = rand.Int63n(sides)
	}
	fmt.Printf("%+v\n", list)
}
