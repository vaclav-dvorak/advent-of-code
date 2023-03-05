package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

var (
	forest [][]int
)

func testTree(x, y, tree int) bool {
	xlen := len(forest[0])
	ylen := len(forest)
	// test left
	testvis := true
	for iy := 0; iy < y; iy++ {
		if forest[x][iy] >= tree {
			testvis = false
			break
		}
	}
	if testvis {
		return true
	}
	// test right
	testvis = true
	for iy := y + 1; iy < ylen; iy++ {
		if forest[x][iy] >= tree {
			testvis = false
			break
		}
	}
	if testvis {
		return true
	}
	// test top
	testvis = true
	for ix := 0; ix < x; ix++ {
		if forest[ix][y] >= tree {
			testvis = false
			break
		}
	}
	if testvis {
		return true
	}
	// test bottom
	testvis = true
	for ix := x + 1; ix < xlen; ix++ {
		if forest[ix][y] >= tree {
			testvis = false
			break
		}
	}
	if testvis {
		return true
	} else {
		return false
	}
}

func getScore(x, y, tree int) int {
	xlen := len(forest[0])
	ylen := len(forest)
	scene := 1
	// test left
	score := 0
	for iy := y - 1; iy >= 0; iy-- {
		score++
		if forest[x][iy] >= tree {
			break
		}
	}
	scene *= score
	// test right
	score = 0
	for iy := y + 1; iy < ylen; iy++ {
		score++
		if forest[x][iy] >= tree {
			break
		}
	}
	scene *= score
	// test top
	score = 0
	for ix := x - 1; ix >= 0; ix-- {
		score++
		if forest[ix][y] >= tree {
			break
		}
	}
	scene *= score
	// test bottom
	score = 0
	for ix := x + 1; ix < xlen; ix++ {
		score++
		if forest[ix][y] >= tree {
			break
		}
	}
	scene *= score
	return scene
}

func parseInput(input io.Reader) {
	scanner := bufio.NewScanner(input)
	forest = make([][]int, 0)
	l := 0
	for scanner.Scan() {
		line := scanner.Text()
		forest = append(forest, make([]int, len(line)))
		for i, ch := range line {
			forest[l][i] = int(ch - '0')
		}
		l++
	}
}

func fs1(input io.Reader) int {
	parseInput(input)
	visible := 0
	for x, row := range forest {
		for y, tree := range row {
			if testTree(x, y, tree) {
				visible++
			}
		}
	}
	return visible
}

func fs2(input io.Reader) int {
	parseInput(input)

	maxScene := 0
	for x, row := range forest {
		for y, tree := range row {
			sceneScore := getScore(x, y, tree)
			if sceneScore > maxScene {
				maxScene = sceneScore
			}
		}
	}
	return maxScene
}

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fs2(f)
}
