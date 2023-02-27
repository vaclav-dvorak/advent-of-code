package main

import (
	"fmt"
)

func main() {
	for i := 0; i < len(input)-14; i++ {
		found := map[byte]bool{}
		nok := false
		for look := 0; look < 14; look++ {
			ch := input[i+look]
			_, ok := found[ch]
			if ok {
				nok = true
				break
			} else {
				found[ch] = true
			}
		}
		if !nok {
			fmt.Printf("%d -> %s\n", i+14, input[i:i+14])

			break
		}
	}
}
