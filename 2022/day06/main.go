package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func solve(input string, lf int) int {
	for i := 0; i < len(input)-lf; i++ {
		found := map[byte]bool{}
		nok := false
		for look := 0; look < lf; look++ {
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
			fmt.Printf("%d -> %s\n", i+lf, input[i:i+lf])
			return i + lf
		}
	}
	return 0
}

func fs1(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		return solve(line, 4)
	}
	return 0
}

func fs2(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		return solve(line, 14)
	}
	return 0
}

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fs1(f)
}
