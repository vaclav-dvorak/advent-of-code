package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	maxBlue  = 14
	maxRed   = 12
	maxGreen = 13
)

func isGamePossible(in string) (possible bool, id int) {
	possible = true
	split := strings.Split(in, ":")
	id, err := strconv.Atoi(split[0][5:])

	if err != nil {
		log.Fatal(err)
	}
	shows := strings.Split(split[1], ";")
	for _, show := range shows {
		colors := strings.Split(show, ",")
		for _, color := range colors {
			val := strings.Split(strings.TrimSpace(color), " ")
			count, err := strconv.Atoi(val[0])
			if err != nil {
				log.Fatal(err)
			}
			switch val[1] {
			case "red":
				if count > maxRed {
					possible = false
					return
				}
			case "green":
				if count > maxGreen {
					possible = false
					return
				}
			case "blue":
				if count > maxBlue {
					possible = false
					return
				}
			}
		}
	}

	return
}

func minValueGame(in string) (value int) {
	var mxRed, mxBlue, mxGreen int = 0, 0, 0
	split := strings.Split(in, ":")
	shows := strings.Split(split[1], ";")
	for _, show := range shows {
		colors := strings.Split(show, ",")
		for _, color := range colors {
			val := strings.Split(strings.TrimSpace(color), " ")
			count, err := strconv.Atoi(val[0])
			if err != nil {
				log.Fatal(err)
			}
			switch val[1] {
			case "red":
				if count > mxRed {
					mxRed = count
				}
			case "green":
				if count > mxGreen {
					mxGreen = count
				}
			case "blue":
				if count > mxBlue {
					mxBlue = count
				}
			}
		}
	}
	value = mxRed * mxGreen * mxBlue
	return
}

func fs1(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if possible, id := isGamePossible(line); possible {
			sum += id
		}
	}
	return sum
}

func fs2(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += minValueGame(line)
	}
	return sum
}

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fs2(f)
}
