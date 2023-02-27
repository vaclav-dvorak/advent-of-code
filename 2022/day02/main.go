package main

import "fmt"

func main() {
	my, match, score := 0, 0, 0
	for i := 0; i < len(input); i++ {
		switch input[i][2] {
		case 'X':
			match = 0
		case 'Y':
			match = 3
		case 'Z':
			match = 6
		}
		switch input[i][0] {
		case 'A':
			switch match {
			case 0:
				my = 3
			case 3:
				my = 1
			case 6:
				my = 2
			}
		case 'B':
			switch match {
			case 0:
				my = 1
			case 3:
				my = 2
			case 6:
				my = 3
			}
		case 'C':
			switch match {
			case 0:
				my = 2
			case 3:
				my = 3
			case 6:
				my = 1
			}
		}
		score += my + match
		// fmt.Printf("%d + %d\n", my, match)

	}
	fmt.Printf("%d\n", score)
}
