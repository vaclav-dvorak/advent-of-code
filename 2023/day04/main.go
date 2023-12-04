package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type card struct {
	line   string
	copies int
	win    bool
}

func solveCard(in string) (value int) {
	split := strings.Split(in, ":")
	nums := strings.Split(split[1], "|")
	winMap := make(map[int]bool)
	for x := 1; x < len(nums[0]); x += 3 {
		n, err := strconv.Atoi(strings.TrimSpace(nums[0][x : x+2]))
		if err != nil {
			log.Fatal(err)
		}
		winMap[n] = true
	}
	for x := 1; x < len(nums[1]); x += 3 {
		n, err := strconv.Atoi(strings.TrimSpace(nums[1][x : x+2]))
		if err != nil {
			log.Fatal(err)
		}
		if found := winMap[n]; found {
			if value == 0 {
				value = 1
			} else {
				value *= 2
			}
		}
	}

	return
}

func processCards(in []card) (sum int) {
	var scoreboard map[int]int = map[int]int{
		0:    0,
		1:    1,
		2:    2,
		4:    3,
		8:    4,
		16:   5,
		32:   6,
		64:   7,
		128:  8,
		256:  9,
		512:  10,
		1024: 11,
	}
	input := in
	for i, card := range input {
		val := scoreboard[solveCard(card.line)]
		if val != 0 {
			input[i].win = true
			for x := 1; x <= val; x++ {
				input[i+x].copies += input[i].copies
			}
		}
	}

	for _, card := range input {
		sum += card.copies
	}
	return
}

func fs1(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += solveCard(line)
	}

	return sum
}

func fs2(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	cards := make([]card, 0)
	for scanner.Scan() {
		line := scanner.Text()
		cards = append(cards, card{
			line:   line,
			copies: 1,
			win:    false,
		})
	}
	sum := processCards(cards)
	fmt.Printf("solve: %d", sum)
	return sum
}

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fs2(f)
}
