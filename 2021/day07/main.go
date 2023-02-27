package main

import (
	"fmt"
	"sort"
)

func main() {
	var fuel, minfuel, mincand int
	minfuel = -1
	sort.Ints(inputs)
	for candidate := inputs[0]; candidate <= inputs[len(inputs)-1]; candidate++ {
		fuel = 0
		for _, crab := range inputs {
			steps := candidate - crab
			if steps < 0 {
				steps = -1 * steps
			}
			fuel += (steps * (steps + 1)) / 2
		}
		if minfuel == -1 || minfuel > fuel {
			minfuel = fuel
			mincand = candidate
		}
	}
	// fmt.Printf("%v\n", median)
	fmt.Printf("candidate: %d\ngas: %d\n", mincand, minfuel)
}
