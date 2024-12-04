package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func searchWord(ws []string) (count int) {
	for y, line := range ws {
		for x, w := range line {
			if w != 'X' {
				continue
			}
			if x > 2 { //can check left
				t := fmt.Sprintf("X%c%c%c", ws[y][x-1], ws[y][x-2], ws[y][x-3])
				if t == "XMAS" {
					count++
				}
				if y > 2 { //can check left/top
					t := fmt.Sprintf("X%c%c%c", ws[y-1][x-1], ws[y-2][x-2], ws[y-3][x-3])
					if t == "XMAS" {
						count++
					}
				}
				if y < len(ws)-3 { //can check left/bottom
					t := fmt.Sprintf("X%c%c%c", ws[y+1][x-1], ws[y+2][x-2], ws[y+3][x-3])
					if t == "XMAS" {
						count++
					}
				}
			}

			if x < len(line)-3 { //can check right
				t := fmt.Sprintf("X%c%c%c", ws[y][x+1], ws[y][x+2], ws[y][x+3])
				if t == "XMAS" {
					count++
				}
				if y > 2 { //can check right/top
					t := fmt.Sprintf("X%c%c%c", ws[y-1][x+1], ws[y-2][x+2], ws[y-3][x+3])
					if t == "XMAS" {
						count++
					}
				}
				if y < len(ws)-3 { //can check right/bottom
					t := fmt.Sprintf("X%c%c%c", ws[y+1][x+1], ws[y+2][x+2], ws[y+3][x+3])
					if t == "XMAS" {
						count++
					}
				}
			}

			if y > 2 { //can check top
				t := fmt.Sprintf("X%c%c%c", ws[y-1][x], ws[y-2][x], ws[y-3][x])
				if t == "XMAS" {
					count++
				}
			}

			if y < len(ws)-3 { //can check bottom
				t := fmt.Sprintf("X%c%c%c", ws[y+1][x], ws[y+2][x], ws[y+3][x])
				if t == "XMAS" {
					count++
				}
			}
		}
	}
	return count
}

func searchWord2(ws []string) (count int) {
	for y, line := range ws {
		for x, w := range line {
			if w != 'A' {
				continue
			}
			if x == 0 || x == len(line)-1 || y == 0 || y == len(ws)-1 { //we ignore the edges
				continue
			}
			t1 := fmt.Sprintf("%cA%c", ws[y-1][x-1], ws[y+1][x+1])
			t2 := fmt.Sprintf("%cA%c", ws[y+1][x-1], ws[y-1][x+1])
			if (t1 == "MAS" || t1 == "SAM") && (t2 == "MAS" || t2 == "SAM") {
				count++
			}
		}
	}
	return count
}

func fs1(input io.Reader) int {
	ws := []string{}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		ws = append(ws, line)
	}
	c := searchWord(ws)
	fmt.Printf("count: %d\n", c)
	return c
}

func fs2(input io.Reader) int {
	ws := []string{}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		ws = append(ws, line)
	}
	c := searchWord2(ws)
	fmt.Printf("count: %d\n", c)
	return c
}

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fs2(f)
}
