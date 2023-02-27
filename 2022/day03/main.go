package main

import "fmt"

func findItem(l, p, m string) (r rune, e error) {
	for _, x := range l {
		for _, y := range p {
			if x == y {
				for _, z := range m {
					if x == z {
						r = x
						return
					}
				}
			}
		}
	}
	e = fmt.Errorf("no match found")
	return
}

func calcPriority(r rune) (p int) {
	test := int(r - 'A')
	if test > 25 {
		p = int(r-'a') + 1
	} else {
		p = test + 27
	}
	return
}

func main() {
	score := 0
	for i := 0; i < len(input); i += 3 {
		item, err := findItem(input[i], input[i+1], input[i+2])
		if err != nil {
			fmt.Printf("%d: %s %s %s - no match found\n", i, input[i], input[i+1], input[i+2])
			return
		}
		score += calcPriority(item)
	}
	fmt.Printf("%d\n", score)
}
