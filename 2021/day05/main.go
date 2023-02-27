package main

import (
	"fmt"
	"math"
)

func main() {
	const size = 1000
	var (
		ocean = [size][size]int{}
	)
	initVectors()
	for _, v := range inputs {
		for m := 0; m <= v.len; m++ {
			ocean[v.start[1]+(m*v.stepy)][v.start[0]+(m*v.stepx)]++
		}
	}
	sum := 0
	for _, line := range ocean {
		for _, cell := range line {
			if cell >= 2 {
				sum++
			}
		}
	}
	fmt.Printf("%d\n", sum)
}

func initVectors() {
	for i, v := range inputs {
		inputs[i].stepx = 1
		inputs[i].stepy = 1
		inputs[i].len = int(math.Abs(float64(v.start[0] - v.end[0])))
		if v.start[0] == v.end[0] {
			inputs[i].stepx = 0
			inputs[i].len = int(math.Abs(float64(v.start[1] - v.end[1])))
		}
		if v.start[0] > v.end[0] {
			inputs[i].stepx = -1
		}
		if v.start[1] == v.end[1] {
			inputs[i].stepy = 0
		}
		if v.start[1] > v.end[1] {
			inputs[i].stepy = -1
		}
	}
}
