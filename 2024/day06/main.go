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

type guard struct {
	x int
	y int
	o byte
}

func doGuarding(g guard, house [][]rune) (stuck bool) {
	done := false
	house[g.y][g.x] = 'X' //mark start as visited
	visited := make(map[string]bool)
	stuck = false
	for !done {
		visited[fmt.Sprintf("%d,%d,%c", g.x, g.y, g.o)] = true
		ahead := ' '
		switch g.o {
		case 'n':
			ahead = house[g.y-1][g.x]
		case 'e':
			ahead = house[g.y][g.x+1]
		case 's':
			ahead = house[g.y+1][g.x]
		case 'w':
			ahead = house[g.y][g.x-1]
		}
		if ahead == 'E' {
			// we are done on edge
			done = true
			continue
		}
		if ahead == '#' {
			// turn right
			switch g.o {
			case 'n':
				g.o = 'e'
			case 'e':
				g.o = 's'
			case 's':
				g.o = 'w'
			case 'w':
				g.o = 'n'
			}
		}
		if ahead == '.' || ahead == 'X' {
			// walk forward
			switch g.o {
			case 'n':
				g.y -= 1
			case 'e':
				g.x += 1
			case 's':
				g.y += 1
			case 'w':
				g.x -= 1
			}
			house[g.y][g.x] = 'X' //mark as visited
		}
		if ok := visited[fmt.Sprintf("%d,%d,%c", g.x, g.y, g.o)]; ok { // if we are at same square facing same direction as before we are stuck
			stuck = true
			break
		}
	}
	return
}

func evalHouse(house [][]rune) int {
	count := 0
	for _, row := range house {
		for _, cell := range row {
			if cell == 'X' {
				count++
			}
		}
	}
	return count
}

func fs1(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	g := guard{}
	house := make([][]rune, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(house) == 0 {
			house = append(house, []rune(strings.Repeat("E", len(line)+2)))
		}
		house = append(house, []rune("E"+line+"E"))
		if strings.Contains(line, "^") {
			x := strings.Index(line, "^") + 1
			y := len(house) - 1
			g = guard{x, y, 'n'}
		}
	}
	house = append(house, []rune(strings.Repeat("E", len(house[0]))))
	doGuarding(g, house)
	return evalHouse(house)
}

func fs2(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	g := guard{}
	house := make([][]rune, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(house) == 0 {
			house = append(house, []rune(strings.Repeat("E", len(line)+2)))
		}
		house = append(house, []rune("E"+line+"E"))
		if strings.Contains(line, "^") {
			x := strings.Index(line, "^") + 1
			y := len(house) - 1
			g = guard{x, y, 'n'}
		}
	}
	house = append(house, []rune(strings.Repeat("E", len(house[0]))))

	oHouse := make([][]rune, len(house))
	for i := range house {
		oHouse[i] = make([]rune, len(house[i]))
		copy(oHouse[i], house[i])
	}
	oGuard := g

	for _, row := range oHouse {
		fmt.Printf("%s\n", string(row))
	}

	doGuarding(g, house) // do initial patrol se we find obstacle candidates
	obstacles := make([]string, 0)
	for y, row := range house {
		for x, cell := range row {
			if cell == 'X' { // all visited cells are potential obstacles
				obstacles = append(obstacles, fmt.Sprintf("%d,%d", x, y))
			}
		}
	}

	stuck := 0
	for _, obs := range obstacles {
		xy := strings.Split(obs, ",")
		x, _ := strconv.Atoi(xy[0])
		y, _ := strconv.Atoi(xy[1])
		testHouse := make([][]rune, len(oHouse))
		for i := range oHouse {
			testHouse[i] = make([]rune, len(oHouse[i]))
			copy(testHouse[i], oHouse[i])
		}
		testGuard := oGuard
		testHouse[y][x] = '#' // add tested obstacle
		if doGuarding(testGuard, testHouse) {
			stuck++
		}
	}
	fmt.Printf("stuck: %d\n", stuck)
	return stuck
}

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fs2(f)
}
