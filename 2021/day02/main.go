package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	fwd, depth, aim := 0, 0, 0
	for _, cmd := range input {
		cmds := strings.Split(cmd, " ")
		l, err := strconv.Atoi(cmds[1])
		if err != nil {
			log.Fatal(err)
		}
		switch cmds[0] {
		case "forward":
			fwd += l
			depth += (aim * l)
		case "down":
			aim += l
		case "up":
			aim -= l
		}
	}
	fmt.Printf("fwd: %d\ndepth: %d\naim:%d\nnum: %d\n", fwd, depth, aim, fwd*depth)
}
