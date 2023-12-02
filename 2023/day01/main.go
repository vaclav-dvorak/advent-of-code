package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

func calcLine(in string) (ret int) {
	var first, last int = -1, -1
	for i := range in {
		if n, err := strconv.Atoi(in[i : i+1]); err == nil {
			if first < 0 {
				first = n
			}
			last = n
		}
	}
	ret = first*10 + last
	return
}

func calcLineWord(in string) (ret int) {
	var first, last int = -1, -1
	for i := range in {
		num := -1
		if i < len(in)-2 {
			switch test := in[i : i+3]; test {
			case "one":
				num = 1
			case "two":
				num = 2
			case "six":
				num = 6
			}
		}

		if i < len(in)-3 {
			switch test := in[i : i+4]; test {
			case "four":
				num = 4
			case "five":
				num = 5
			case "nine":
				num = 9
			}
		}

		if i < len(in)-4 {
			switch test := in[i : i+5]; test {
			case "three":
				num = 3
			case "seven":
				num = 7
			case "eight":
				num = 8
			}
		}

		if n, err := strconv.Atoi(in[i : i+1]); err == nil {
			num = n
		}

		if num >= 0 {
			if first < 0 {
				first = num
			}
			last = num
		}

	}
	ret = first*10 + last
	return
}

func fs1(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += calcLine(line)
	}

	return sum
}

func fs2(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += calcLineWord(line)
	}

	return sum
}

func main() {
	f, err := os.Open("test2.txt")
	if err != nil {
		log.Fatal(err)
	}
	fs2(f)
}
