package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

type octo struct {
	x       int
	y       int
	eng     int
	flashed bool
}

func (o *octo) flash() {
	if o.flashed {
		return
	}
	// inc all octos around me
	for x := o.x - 1; x <= o.x+1; x++ {
		for y := o.y - 1; y <= o.y+1; y++ {
			if x < 0 || y < 0 || x > 9 || y > 9 { //guard boundries
				continue
			}
			cave[x][y].eng++
		}
	}
	o.flashed = true
}

var cave [][]octo

func fs1(input io.Reader) int {
	flashes := 0
	scanner := bufio.NewScanner(input)
	cave = make([][]octo, 10)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		for x, val := range line {
			if cave[x] == nil {
				cave[x] = make([]octo, 10)
			}
			cave[x][y] = octo{
				x:       x,
				y:       y,
				eng:     int(val - '0'),
				flashed: false,
			}
		}
		y++
	}

	steps := 100
	for it := 0; it < steps; it++ {
		// add 1 eng to all octos in cave
		for y := 0; y < 10; y++ {
			for x := 0; x < 10; x++ {
				cave[x][y].eng++
			}
		}
		// flash all octos until done
		alldone := false
		for !alldone {
			alldone = true
			for y := 0; y < 10; y++ {
				for x := 0; x < 10; x++ {
					if cave[x][y].eng > 9 && !cave[x][y].flashed {
						cave[x][y].flash()
						flashes++
						alldone = false
					}
				}
			}
		}

		// reset all octos in cave
		for y := 0; y < 10; y++ {
			for x := 0; x < 10; x++ {
				if cave[x][y].flashed {
					cave[x][y].eng = 0
					cave[x][y].flashed = false
				}
			}
		}

	}
	// for y := 0; y < 10; y++ {
	// 	for x := 0; x < 10; x++ {
	// 		fmt.Print(cave[x][y].eng)
	// 	}
	// 	fmt.Println()
	// }

	return flashes
}

func fs2(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	cave = make([][]octo, 10)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		for x, val := range line {
			if cave[x] == nil {
				cave[x] = make([]octo, 10)
			}
			cave[x][y] = octo{
				x:       x,
				y:       y,
				eng:     int(val - '0'),
				flashed: false,
			}
		}
		y++
	}

	steps := 0
	totale := 100
	for totale > 0 {
		// add 1 eng to all octos in cave
		for y := 0; y < 10; y++ {
			for x := 0; x < 10; x++ {
				cave[x][y].eng++
			}
		}
		// flash all octos until done
		alldone := false
		for !alldone {
			alldone = true
			for y := 0; y < 10; y++ {
				for x := 0; x < 10; x++ {
					if cave[x][y].eng > 9 && !cave[x][y].flashed {
						cave[x][y].flash()
						alldone = false
					}
				}
			}
		}

		// reset all octos in cave and count total energy
		totale = 0
		for y := 0; y < 10; y++ {
			for x := 0; x < 10; x++ {
				if cave[x][y].flashed {
					cave[x][y].eng = 0
					cave[x][y].flashed = false
				}
				totale += cave[x][y].eng
			}
		}
		steps++

	}

	return steps
}

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fs2(f)
}
