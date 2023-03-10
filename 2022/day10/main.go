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

type ticker struct {
	state int
}

var (
	err     error
	xState  int
	result  int
	monitor []rune
)

func (t *ticker) tick() {
	t.state++
	if (t.state-20)%40 == 0 {
		fmt.Printf("%d: x=%d\n", t.state, xState)
		result += t.state * xState
	}
}

func (t *ticker) tick2() {
	//before each tick we draw
	pos := t.state % 40
	if xState >= pos-1 && xState <= pos+1 {
		monitor[t.state] = '#'
	} else {
		monitor[t.state] = '.'
	}
	t.state++
}

func fs1(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	cycle := ticker{}
	xState = 1
	result = 0
	for scanner.Scan() {
		line := scanner.Text()
		params := strings.Split(line, " ")
		cmd := params[0]
		count := 0
		if cmd == "addx" {
			count, err = strconv.Atoi(params[1])
			if err != nil {
				log.Fatal(err)
			}
		}
		switch cmd {
		case "noop":
			cycle.tick()
		case "addx":
			cycle.tick()
			cycle.tick()
			xState += count
		}
	}
	return result
}

func fs2(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	monitor = make([]rune, 240)
	cycle := ticker{}
	xState = 1
	result = 0
	for scanner.Scan() {
		line := scanner.Text()
		params := strings.Split(line, " ")
		cmd := params[0]
		count := 0
		if cmd == "addx" {
			count, err = strconv.Atoi(params[1])
			if err != nil {
				log.Fatal(err)
			}
		}
		switch cmd {
		case "noop":
			cycle.tick2()
		case "addx":
			cycle.tick2()
			cycle.tick2()
			xState += count
		}
	}
	for y := 0; y < 6; y++ {
		for x := 0; x < 40; x++ {
			fmt.Printf("%c", monitor[y*40+x])
		}
		fmt.Print("\n")
	}
	return 42

}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fs2(f)
}
