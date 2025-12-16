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

type prob struct {
	nums []int
	op   string
}

func fs1(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	problems := []prob{}
	for scanner.Scan() {
		line := scanner.Text()
		inp := strings.Split(line, " ")
		po := 0
		for i := range inp {
			val := strings.TrimSpace(inp[i])
			if val == "" {
				continue
			}
			if val == "+" || val == "*" {
				problems[po].op = val
			} else {
				val, _ := strconv.Atoi(val)
				if len(problems) <= po {
					problems = append(problems, prob{})
				}
				problems[po].nums = append(problems[po].nums, val)
			}
			po++
		}
	}
	
	sum := 0
	for _, p := range problems {
		problemSum := 0
		for _, n := range p.nums {
			if p.op == "+" {
				problemSum += n
			} else if p.op == "*" {
				if problemSum == 0 {
					problemSum = 1
				}
				problemSum *= n
			}
		}
		sum += problemSum
	}
	fmt.Println(sum)
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
