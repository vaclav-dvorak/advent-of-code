package main

import "fmt"

func main() {
	inc := 0
	for i := 3; i < len(input); i++ {
		if (input[i] + input[i-1] + input[i-2]) > (input[i-1] + input[i-2] + input[i-3]) {
			inc++
		}
	}
	fmt.Printf("%d", inc)
}
