package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

type dial struct {
	position   int
	struckZero int
}

func rotate(d *dial, steps int, side string) {
	if side == "R" {
		d.position = (d.position + steps%100) % 100
	} else if side == "L" {
		d.position = (d.position - steps%100 + 100) % 100
	}
	if d.position == 0 {
		d.struckZero++
	}
}

func rotate2(d *dial, steps int, side string) {
	if side == "R" {
		over := (d.position + steps - 1) / 100
		d.struckZero += over
		d.position = (d.position + steps%100) % 100
	} else if side == "L" {
		over := (d.position - steps + 1) / 100
		d.struckZero += -over
		if d.position != 0 && d.position-steps < 0 {
			d.struckZero++
		}
		d.position = (d.position - steps%100 + 100) % 100
	}
	if d.position == 0 {
		d.struckZero++
	}
}

type move struct {
	steps int
	side  string
}

func fs1(input io.Reader) int {
	var moves []move
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		s, _ := strconv.ParseInt(line[1:], 10, 0)
		moves = append(moves, move{
			steps: int(s),
			side:  string(line[0]),
		})
	}
	dialA := &dial{position: 50, struckZero: 0}
	for i := range moves {
		rotate(dialA, moves[i].steps, moves[i].side)
	}
	return dialA.struckZero
}

func fs2(input io.Reader) int {
	var moves []move
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		s, _ := strconv.ParseInt(line[1:], 10, 0)
		moves = append(moves, move{
			steps: int(s),
			side:  string(line[0]),
		})
	}
	dialA := &dial{position: 50, struckZero: 0}
	for i := range moves {
		rotate2(dialA, moves[i].steps, moves[i].side)
	}
	return dialA.struckZero
}

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fs2(f)
}
