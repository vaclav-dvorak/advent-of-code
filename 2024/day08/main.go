package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type antenna struct {
	x int
	y int
	f rune
}

func findAntennas(city [][]rune) (antennas []antenna) {
	for y, row := range city {
		for x, cell := range row {
			if cell != '.' {
				antennas = append(antennas, antenna{x, y, cell})
			}
		}
	}
	return
}

func hashAntennas(antennas []antenna) (antennasHash map[string]bool) {
	antennasHash = make(map[string]bool)
	for _, a := range antennas {
		antennasHash[fmt.Sprintf("%c,%d,%d", a.f, a.x, a.y)] = true
	}
	return
}

func findAntinodes(city [][]rune, antennas []antenna, antennasHash map[string]bool) (count int) {
	for y, row := range city {
		for x := range row {
			for _, a := range antennas {
				distX := a.x - x
				distY := a.y - y
				if distX == 0 && distY == 0 {
					continue
				}
				hash := fmt.Sprintf("%c,%d,%d", a.f, x+2*distX, y+2*distY)
				if ok := antennasHash[hash]; ok {
					count++
					break
				}
			}
		}
	}
	return
}

// greatest common divisor (GCD) via Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	if a < 0 {
		a = -a
	}
	return a
}

func findAntinodeLines(city [][]rune, antennas []antenna) (count int) {
	for y, row := range city {
		for x := range row {
			for _, a := range antennas {
				distX := a.x - x
				distY := a.y - y
				if distX == 0 && distY == 0 {
					// we are on antenna
					found := false
					for _, at := range antennas {
						if at.f == a.f && at.x != a.x && at.y != a.y {
							found = true
							break
						}
					}
					if found { // we found another antenna with same frequency
						count++
						break
					}
				}
				// we will compare vectors
				di := gcd(distX, distY)
				vector := fmt.Sprintf("%d,%d", distX/di, distY/di)
				found := false
				for _, at := range antennas {
					if (at.x == a.x && at.y == a.y) || (at.x == x && at.y == y) || (at.f != a.f) { //do not same antenna twice
						continue
					}
					distTX := at.x - x
					distTY := at.y - y
					diT := gcd(distTX, distTY)
					vectorT := fmt.Sprintf("%d,%d", distTX/diT, distTY/diT)
					if vector == vectorT {
						// we are on same vector
						found = true
						break
					}
				}
				if found { // we found another antenna with same frequency
					count++
					break
				}
			}
		}
	}
	return
}

func fs1(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	city := make([][]rune, 0)
	for scanner.Scan() {
		line := scanner.Text()
		city = append(city, []rune(line))
	}

	antennas := findAntennas(city)
	antennasHash := hashAntennas(antennas)
	c := findAntinodes(city, antennas, antennasHash)
	fmt.Printf("antinodes: %d\n", c)
	return c
}

func fs2(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	city := make([][]rune, 0)
	for scanner.Scan() {
		line := scanner.Text()
		city = append(city, []rune(line))
	}

	antennas := findAntennas(city)
	c := findAntinodeLines(city, antennas)
	fmt.Printf("antinodes: %d\n", c)
	return c
}

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fs2(f)
}
