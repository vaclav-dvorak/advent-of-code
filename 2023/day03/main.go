package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type part struct {
	num int
	x1  int
	y1  int
	x2  int
	y2  int
}

func (p part) contain(x, y int) (res bool) {
	res = (x >= p.x1) && (x <= p.x2) && (y >= p.y1) && (y <= p.y2)
	return
}

func solveEngine(in []string) (sum int) {
	number := 0
	for y, line := range in {
		for x := range line {
			num, err := strconv.Atoi(line[x : x+1])
			if err == nil {
				number = number*10 + num
			}
			if err != nil || x == len(line)-1 { //have to run this on end of lines too
				if number == 0 {
					continue
				}
				l := len(strconv.Itoa(number))
				xl := x - l - 1
				bigTest := false
				fmt.Printf("lookaround for (%d): x:%d, y:%d, l:%d\n", number, xl, y, l)

				//check sides
				if xl > 0 { //test left char
					test := in[y][xl:xl+1] != "."
					if test {
						bigTest = test
					}
				}
				if xl+l+1 != len(line) { //test right char
					test := in[y][xl+l+1:xl+l+2] != "."
					if test {
						bigTest = test
					}
				}

				if xl < 0 { // reset left bound so it doesnt hit boundary in line scan
					l = l - (0 - xl)
					xl = 0
				}

				if y > 0 { //test above line
					for tx := xl; tx < xl+l+2; tx++ {
						test := in[y-1][tx:tx+1] != "."
						if test {
							bigTest = test
						}
					}
				}

				if y < len(in)-1 { //test below line
					for tx := xl; tx < xl+l+2; tx++ {
						test := in[y+1][tx:tx+1] != "."
						if test {
							bigTest = test
						}
					}
				}

				if bigTest {
					sum += number
				}
				number = 0
			}
		}
	}
	return
}

func readEngine(in []string) (out []part) {
	number := 0
	for y, line := range in {
		for x := range line {
			num, err := strconv.Atoi(line[x : x+1])
			if err == nil {
				number = number*10 + num
			}
			if err != nil || x == len(line)-1 { //have to run this on end of lines too
				if number == 0 {
					continue
				}
				l := len(strconv.Itoa(number))
				x1 := x - l - 1
				if x1 < 0 {
					x1 = 0
				}
				y1 := y - 1
				if y1 < 0 {
					y1 = 0
				}
				x2 := x
				y2 := y + 1
				out = append(out, part{
					num: number,
					x1:  x1,
					y1:  y1,
					x2:  x2,
					y2:  y2,
				})
				number = 0
			}
		}
	}
	return
}

func solveGears(in []string, parts []part) (sum int) {
	sum = 0
	for y, line := range in {
		for x := range line {
			if line[x:x+1] != "*" { //!not a gear
				continue
			}
			around := make([]part, 0)
			for _, part := range parts {
				if part.contain(x, y) {
					around = append(around, part)
				}
			}
			if len(around) > 2 {
				fmt.Printf("gear x:%d, y:%d... more then two around\n", x, y)
			}
			if len(around) < 2 { //!less then two around
				continue
			}
			sum += around[0].num * around[1].num
		}
	}
	return
}

func fs1(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	engine := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		engine = append(engine, line)
	}
	sum := solveEngine(engine)
	fmt.Printf("solve %d\n", sum)
	return sum
}

func fs2(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	engine := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		engine = append(engine, line)
	}
	parts := readEngine(engine)
	sum := solveGears(engine, parts)
	fmt.Printf("solve %d\n", sum)
	return sum
}

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fs2(f)
}
