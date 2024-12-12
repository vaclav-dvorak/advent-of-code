package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type antenna struct {
	f rune
	x int
	y int
}

// Function to find antennas and create the hash map
func findAntennasAndHash(city [][]rune) (antennas []antenna, antennasHash map[antenna]bool) {
	antennasHash = make(map[antenna]bool)
	for y, row := range city {
		for x, cell := range row {
			if cell != '.' {
				a := antenna{cell, x, y}
				antennas = append(antennas, a)
				// Populate the antennasHash map with the transformed coordinates
				antennasHash[a] = true
			}
		}
	}
	return
}

func findAntinodes(city [][]rune, antennas []antenna, antennasHash map[antenna]bool) (count int) {
	for y, row := range city {
		for x := range row {
			for _, a := range antennas {
				distX := a.x - x
				distY := a.y - y
				if distX == 0 && distY == 0 {
					continue
				}
				if ok := antennasHash[antenna{a.f, x + 2*distX, y + 2*distY}]; ok {
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
				distX, distY := a.x-x, a.y-y
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
				vector := [2]int{distX / di, distY / di}
				found := false
				for _, at := range antennas {
					if (at.x == a.x && at.y == a.y) || (at.x == x && at.y == y) || (at.f != a.f) { //do not same antenna twice
						continue
					}
					distTX, distTY := at.x-x, at.y-y
					diT := gcd(distTX, distTY)
					vectorT := [2]int{distTX / diT, distTY / diT}
					if vector == vectorT {
						// we are on same vector
						found = true
						break
					}
				}
				if found { // we found another antenna on same vectorwith same frequency
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

	antennas, antennasHash := findAntennasAndHash(city)
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

	antennas, _ := findAntennasAndHash(city)
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
