package main

import (
	"fmt"
)

func main() {
	readings := initReadings()
	getFirstAnswer(readings)
}

func getFirstAnswer(in []reading) {
	ident := 0
	valid := map[int]struct{}{2: {}, 3: {}, 4: {}, 7: {}}
	for _, val := range in {
		for _, sig := range val.signals {
			if _, ok := valid[len(sig)]; ok {
				ident++
				// fmt.Printf("%s\n", sig)
			}
		}
	}
	fmt.Printf("identified: %d\n", ident)
}
