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

type knot struct {
	x int
	y int
}

var (
	rope    []knot
	visited map[string]bool
)

func headMoves(head *knot, d string) (err error) {
	switch d {
	case "U":
		head.y++
	case "D":
		head.y--
	case "R":
		head.x++
	case "L":
		head.x--
	default:
		err = fmt.Errorf("unknown dir: %s", d)
	}
	return
}

func knotFollows(tail *knot, head knot, isTail bool) {
	if tail.y == head.y {
		if (tail.x - 2) == head.x { // move direct left
			tail.x--
		}
		if (tail.x + 2) == head.x { // move direct right
			tail.x++
		}
	} else if tail.x == head.x {
		if (tail.y - 2) == head.y { // move direct down
			tail.y--
		}
		if (tail.y + 2) == head.y { // move direct right
			tail.y++
		}
	} else if (tail.y-1 > head.y) || (tail.y+1 < head.y) || (tail.x-1 > head.x) || (tail.x+1 < head.x) {
		if head.x > tail.x && head.y > tail.y { // move right/up
			tail.x++
			tail.y++
		}
		if head.x > tail.x && head.y < tail.y { // move right/down
			tail.x++
			tail.y--
		}
		if head.x < tail.x && head.y > tail.y { // move left/up
			tail.x--
			tail.y++
		}
		if head.x < tail.x && head.y < tail.y { // move left/down
			tail.x--
			tail.y--
		}
	}
	if isTail {
		tailPos := fmt.Sprintf("%d|%d", tail.x, tail.y)
		if _, ok := visited[tailPos]; !ok {
			visited[tailPos] = true
		}
	}
}

func solveRope(ln int, input io.Reader) int {
	visited = make(map[string]bool, 0)
	visited["0|0"] = true //seed start as visited position
	rope = make([]knot, ln)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		params := strings.Split(line, " ")
		dir := params[0]
		count, err := strconv.Atoi(params[1])
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < count; i++ {
			for n := range rope {
				if n == 0 {
					if err := headMoves(&rope[n], dir); err != nil {
						log.Fatal(err)
					}
					continue
				}
				knotFollows(&rope[n], rope[n-1], n == (ln-1))
			}
		}
	}
	return len(visited)
}

func fs1(input io.Reader) int {
	return solveRope(2, input)
}

func fs2(input io.Reader) int {
	return solveRope(10, input)
}

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fs2(f)
}
