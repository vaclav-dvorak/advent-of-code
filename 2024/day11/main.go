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

func solveStones(input io.Reader, gens int) int {
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
	blinks := 0
	pows := make([]int, 100)
	for blinks < gens {
		newStones := make([]int, 0)
		for i := 0; i < len(stones); i++ {
			if stones[i] == 0 {
				newStones = append(newStones, 1)
				continue
			}
			digits := countDigits(stones[i])
			if digits%2 != 0 {
				newStones = append(newStones, stones[i]*2024)
				continue
			}
			// we split stones
			half := digits / 2
			if ok := pows[half]; ok == 0 {
				pows[half] = int(math.Pow10(half))
			}
			left := stones[i] / pows[half]
			right := stones[i] - left*pows[half]
			newStones = append(newStones, left, right)
		}
		stones = newStones
		blinks++
		fmt.Printf("i: %d len: %d\n", blinks, len(stones))
	}
	fmt.Printf("len: %d\n", len(stones))
	return len(stones)
}

func fs1(input io.Reader) int {
	return solveStones(input, 25)
}

func fs2(input io.Reader) int {
	return 42
	// return solveStones(input, 75)
}

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fs2(f)
}
