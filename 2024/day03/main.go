package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

func fs1(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	r, _ := regexp.Compile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	opreg, _ := regexp.Compile("[0-9]{1,3}")
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		muls := r.FindAllString(line, -1)
		for _, mul := range muls {
			ops := opreg.FindAllString(mul, -1)
			a, _ := strconv.Atoi(ops[0])
			b, _ := strconv.Atoi(ops[1])
			sum += a * b
		}
	}
	fmt.Printf("Sum: %d\n", sum)
	return sum
}

func fs2(input io.Reader) int {
	r, _ := regexp.Compile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	opreg, _ := regexp.Compile("[0-9]{1,3}")
	sum := 0
	scanner := bufio.NewScanner(input)
	mulsOn := true
	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i < len(line); i++ {
			if line[i] == 'd' {
				ahead := ""
				if i < len(line)-4 {
					ahead = line[i : i+4]
				}
				if ahead == "do()" {
					mulsOn = true
				}
				if i < len(line)-7 {
					ahead = line[i : i+7]
				}
				if ahead == "don't()" {
					mulsOn = false
				}
			}
			if line[i] == 'm' && mulsOn {
				ahead := ""
				if i < len(line)-12 {
					ahead = line[i : i+12]
				} else {
					ahead = line[i:]
				}
				muls := r.FindAllString(ahead, -1)
				for _, mul := range muls {
					ops := opreg.FindAllString(mul, -1)
					a, _ := strconv.Atoi(ops[0])
					b, _ := strconv.Atoi(ops[1])
					sum += a * b
				}
			}
		}
	}
	fmt.Printf("Sum: %d\n", sum)
	return sum
}

func main() {
	f, err := os.Open("test2.txt")
	if err != nil {
		log.Fatal(err)
	}
	fs2(f)
}
