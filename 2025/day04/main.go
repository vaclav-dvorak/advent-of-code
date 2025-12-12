package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func fs1(input io.Reader) int {
	were := make([][]int, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := make([]int, 0)
		for i := 0; i < len(scanner.Text()); i++ {
			if scanner.Text()[i] == '.' {
				line = append(line, 0)
			} else {
				line = append(line, 1)
			}
		}
		were = append(were, line)
		_ = line
	}
	canLift := 0
	for y := 0; y < len(were); y++ {
		for x := 0; x < len(were[y]); x++ {
			if were[y][x] == 0 {
				continue
			}
			sum := 0
			if x-1 >= 0 { // left
				sum += were[y][x-1]
				if y-1 >= 0 { // top left
					sum += were[y-1][x-1]
				}
				if y+1 < len(were) { // bottom left
					sum += were[y+1][x-1]
				}
			}
			if x+1 < len(were[y]) { // right
				sum += were[y][x+1]
				if y-1 >= 0 { // top right
					sum += were[y-1][x+1]
				}
				if y+1 < len(were) { // bottom right
					sum += were[y+1][x+1]
				}
			}
			if y-1 >= 0 { // top
				sum += were[y-1][x]
			}
			if y+1 < len(were) { // bottom
				sum += were[y+1][x]
			}
			if sum < 4 {
				canLift++
			}
		}
	}
	return canLift
}

func fs2(input io.Reader) int {
	were := make([][]int, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := make([]int, 0)
		for i := 0; i < len(scanner.Text()); i++ {
			if scanner.Text()[i] == '.' {
				line = append(line, 0)
			} else {
				line = append(line, 1)
			}
		}
		were = append(were, line)
		_ = line
	}
	canLift := true
	lifted := 0
	for canLift {
		canLift = false
		newWere := make([][]int, len(were))
		for y := 0; y < len(were); y++ {
			newWere[y] = make([]int, len(were[y]))
			copy(newWere[y], were[y])
		}
		for y := 0; y < len(were); y++ {
			for x := 0; x < len(were[y]); x++ {
				if were[y][x] == 0 {
					continue
				}
				sum := 0
				if x-1 >= 0 { // left
					sum += were[y][x-1]
					if y-1 >= 0 { // top left
						sum += were[y-1][x-1]
					}
					if y+1 < len(were) { // bottom left
						sum += were[y+1][x-1]
					}
				}
				if x+1 < len(were[y]) { // right
					sum += were[y][x+1]
					if y-1 >= 0 { // top right
						sum += were[y-1][x+1]
					}
					if y+1 < len(were) { // bottom right
						sum += were[y+1][x+1]
					}
				}
				if y-1 >= 0 { // top
					sum += were[y-1][x]
				}
				if y+1 < len(were) { // bottom
					sum += were[y+1][x]
				}
				if sum < 4 {
					canLift = true
					lifted++
					newWere[y][x] = 0
				}
			}
		}
		were = newWere
	}
	return lifted
}

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fs2(f)
}
