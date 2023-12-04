package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func processLine(in string) (out int) {
	var uniq map[int]bool = map[int]bool{2: true, 4: true, 3: true, 7: true}
	split := strings.Split(in, " | ")
	digits := strings.Split(split[1], " ")

	for _, digit := range digits {
		l := len(digit)
		if found := uniq[l]; found {
			out++
		}
	}
	return
}

func fs1(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += processLine(line)
	}
	return sum
}

func fs2(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		_ = line
	}

	return 42
}

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fs1(f)
}
