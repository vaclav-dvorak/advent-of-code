package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type point struct {
	y int
	x int
}

func findLowPointsSum(in [][]int) (out int) {
	for y, row := range in {
		for x, val := range row {
			l, u, r, d := 10, 10, 10, 10
			if x > 0 {
				l = in[y][x-1]
			}
			if x < len(row)-1 {
				r = in[y][x+1]
			}
			if y > 0 {
				u = in[y-1][x]
			}
			if y < len(in)-1 {
				d = in[y+1][x]
			}
			if val < u && val < r && val < d && val < l {
				out += val + 1
			}
		}
	}
	return
}

func findBasins(in [][]int) (out int) {
	// find low points first
	basins := make([]map[string]bool, 0)
	maxX, maxY := 0, len(in)-1
	for y, row := range in {
		maxX = len(row) - 1
		for x, val := range row {
			l, u, r, d := 10, 10, 10, 10
			if x > 0 {
				l = in[y][x-1]
			}
			if x < maxX {
				r = in[y][x+1]
			}
			if y > 0 {
				u = in[y-1][x]
			}
			if y < maxY {
				d = in[y+1][x]
			}
			if val < u && val < r && val < d && val < l {
				basins = append(basins, map[string]bool{fmt.Sprintf("%d|%d", y, x): true})
			}
		}
	}
	for _, basin := range basins {
		candidates := make([]point, 0)
		for coor := range basin {
			start := strings.Split(coor, "|")
			startY, _ := strconv.Atoi(start[0])
			startX, _ := strconv.Atoi(start[1])
			candidates = append(candidates, point{x: startX, y: startY})
			break // we only care about first point
		}
		for len(candidates) > 0 {
			for _, cand := range candidates {
				if _, ok := basin[fmt.Sprintf("%d|%d", cand.y-1, cand.x)]; !ok && cand.y != 0 && in[cand.y-1][cand.x] != 9 {
					candidates = append(candidates, point{y: cand.y - 1, x: cand.x})
					basin[fmt.Sprintf("%d|%d", cand.y-1, cand.x)] = true
				}
				if _, ok := basin[fmt.Sprintf("%d|%d", cand.y, cand.x+1)]; !ok && cand.x != maxX && in[cand.y][cand.x+1] != 9 {
					candidates = append(candidates, point{y: cand.y, x: cand.x + 1})
					basin[fmt.Sprintf("%d|%d", cand.y, cand.x+1)] = true
				}
				if _, ok := basin[fmt.Sprintf("%d|%d", cand.y+1, cand.x)]; !ok && cand.y != maxY && in[cand.y+1][cand.x] != 9 {
					candidates = append(candidates, point{y: cand.y + 1, x: cand.x})
					basin[fmt.Sprintf("%d|%d", cand.y+1, cand.x)] = true
				}
				if _, ok := basin[fmt.Sprintf("%d|%d", cand.y, cand.x-1)]; !ok && cand.x != 0 && in[cand.y][cand.x-1] != 9 {
					candidates = append(candidates, point{y: cand.y, x: cand.x - 1})
					basin[fmt.Sprintf("%d|%d", cand.y, cand.x-1)] = true
				}
				candidates = candidates[1:] // remove candidate from slice
			}
		}
	}

	sort.Slice(basins, func(i, j int) bool { //sort basins to find biggest ones
		return len(basins[i]) > len(basins[j])
	})

	return len(basins[0]) * len(basins[1]) * len(basins[2])
}

func fs1(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	heightMap := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		for x, val := range line {
			row[x] = int(val - '0')
		}
		heightMap = append(heightMap, row)
	}
	return findLowPointsSum(heightMap)
}

func fs2(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	heightMap := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		for x, val := range line {
			row[x] = int(val - '0')
		}
		heightMap = append(heightMap, row)
	}
	return findBasins(heightMap)
}

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fs2(f)
}
