package main

import (
	"fmt"
	"strconv"
)

func main() {

	gamma, epsilon := "", ""
	for pos := 0; pos < len(inputs[0]); pos++ {
		zero := 0
		for _, bajt := range inputs {
			bit := bajt[pos : pos+1]
			if bit == "0" {
				zero++
			}
		}
		if zero > len(inputs)-zero { //? in this case 0 is most common bit
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}
	gd, _ := strconv.ParseInt(gamma, 2, 64)
	ed, _ := strconv.ParseInt(epsilon, 2, 64)
	od := findOx()
	cd := findCO2()
	fmt.Printf("gamma: %d\nepsilon: %d\nnum: %d\n", gd, ed, gd*ed)
	fmt.Printf("ox: %d\nco: %d\nnum: %d\n", od, cd, od*cd)
}

func findOx() int64 {
	ox := ""
	haystack := inputs
	for pos := 0; pos < len(inputs[0]); pos++ {
		zero := 0
		zeros := make([]string, 0)
		ones := make([]string, 0)
		for _, bajt := range haystack {
			bit := bajt[pos : pos+1]
			if bit == "0" {
				zero++
				zeros = append(zeros, bajt)
			} else {
				ones = append(ones, bajt)
			}
		}
		if zero > len(haystack)-zero { //? in this case 0 is most common bit
			haystack = zeros
		} else {
			haystack = ones
		}
		if len(haystack) == 1 {
			ox = haystack[0]
			break
		}
	}
	if ret, err := strconv.ParseInt(ox, 2, 64); err == nil {
		return ret
	}
	return 0
}

func findCO2() int64 {
	co := ""
	haystack := inputs
	for pos := 0; pos < len(inputs[0]); pos++ {
		zero := 0
		zeros := make([]string, 0)
		ones := make([]string, 0)
		for _, bajt := range haystack {
			bit := bajt[pos : pos+1]
			if bit == "0" {
				zero++
				zeros = append(zeros, bajt)
			} else {
				ones = append(ones, bajt)
			}
		}
		if zero <= len(haystack)-zero { //? in this case 0 is least common bit
			haystack = zeros
		} else {
			haystack = ones
		}
		if len(haystack) == 1 {
			co = haystack[0]
			break
		}
	}
	if ret, err := strconv.ParseInt(co, 2, 64); err == nil {
		return ret
	}
	return 0
}
