package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// function, which takes a string as
// argument and return the reverse of string.
func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func main() {
	res := start
	r, _ := regexp.Compile("([0-9]+)")
	for _, s := range input {
		in := r.FindAllString(s, -1)
		count, _ := strconv.Atoi(in[0])
		fromI, _ := strconv.Atoi(in[1])
		from := res[fromI-1]
		toI, _ := strconv.Atoi(in[2])
		to := res[toI-1]
		// fmt.Printf("%s , c=%d , l=%d\n", from, count, len(from))
		moved := from[len(from)-count:]
		// moved = reverse(moved)
		newFrom := from[:len(from)-count]
		newTo := strings.Join([]string{to, moved}, "")
		// fmt.Printf("%d : %s(%s) -> %s -> %s(%s)\n", count, newFrom,from,moved, newTo,to)
		res[fromI-1] = newFrom
		res[toI-1] = newTo
		// fmt.Printf("%+v\n", res)

	}
	output := ""
	for _, s := range res {
		output += s[len(s)-1:]
	}
	fmt.Printf("%s\n", output)
}
