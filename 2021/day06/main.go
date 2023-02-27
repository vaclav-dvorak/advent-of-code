package main

import (
	"fmt"
)

func main() {
	const days = 256
	pool := make([]int, 9)
	for _, fish := range inputs {
		pool[fish]++
	}
	fmt.Printf("%v\n", pool)
	for day := 0; day < days; day++ {
		pop := pool[0]
		newpool := append(pool[1:], 0)
		newpool[6] += pop
		newpool[8] += pop
		pool = newpool
		// fmt.Printf("%v\n", newpool)
	}
	sum := 0
	for _, i := range pool {
		sum += i
	}
	fmt.Printf("%d\n", sum)
}
