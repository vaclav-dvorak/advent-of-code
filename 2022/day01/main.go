package main

import "fmt"

func main() {
	max1, max2, max3, cur := 0, 0, 0, 0
	for i := 0; i < len(input); i++ {
		if input[i] == -1 {
			if cur > max1 {
				max3, max2, max1 = max2, max1, cur
			} else if cur > max2 {
				max3, max2 = max2, cur
			} else if cur > max3 {
				max3 = cur
			}
			cur = 0
		} else {
			cur += input[i]
		}
	}
	fmt.Printf("%d\n", max1+max2+max3)
}
