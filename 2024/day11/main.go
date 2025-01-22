package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func countDigits(n int) int {
	if n == 0 {
		return 1 // Special case: 0 has 1 digit
	}
	count := 0
	for n > 0 {
		count++
		n /= 10
	}
	return count
}

func blinkStone(n, blinks int) int {
	if ok := stoneCache[stoneItem{n, blinks}]; ok != 0 {
		return stoneCache[stoneItem{n, blinks}]
	}
	if blinks == 0 {
		return 1
	}
	l := 0
	digits := countDigits(n)

	if n == 0 {
		l = blinkStone(1, blinks-1)
	} else if digits%2 != 0 {
		l = blinkStone(n*2024, blinks-1)
	} else {
		half := digits / 2
		pow := int(math.Pow10(half))
		left := n / pow
		l1 := blinkStone(left, blinks-1)
		l2 := blinkStone(n-left*pow, blinks-1)
		l = l1 + l2
	}
	stoneCache[stoneItem{n, blinks}] = l
	return l
}

type stoneItem struct {
	n      int
	blinks int
}

var stoneCache map[stoneItem]int

func solveStones(input io.Reader, gens int) int {
	stoneCache = make(map[stoneItem]int)
	stones := make([]int, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		for _, s := range split {
			n, _ := strconv.Atoi(s)
			stones = append(stones, n)
		}
	}
	sum := 0
	for i := 0; i < len(stones); i++ {
		sum += blinkStone(stones[i], gens)
	}
	fmt.Printf("sum: %d\n", sum)
	return sum
}

func fs1(input io.Reader) int {
	return solveStones(input, 25)
}

func fs2(input io.Reader) int {
	return solveStones(input, 75)
}

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fs2(f)
}
