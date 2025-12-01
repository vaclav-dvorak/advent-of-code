package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type cave struct {
	name         string
	isSmall      bool
	visitedTimes int
	adjacent     []*cave
}

var caves map[string]*cave
var pathCount int
var visited map[string]bool

func findPaths(caves map[string]*cave, current *cave, visited map[string]bool) {
	start := caves["start"]
	end := caves["end"]
	if current == end {
		pathCount++
		return
	}
	if current.isSmall {
		visited[current.name] = true
	}
	for _, next := range current.adjacent {
		if next == start || (next.isSmall && visited[next.name]) {
			continue
		}
		findPaths(caves, next, visited)
	}
	if current.isSmall {
		visited[current.name] = false
	}
}

func findPaths2(caves map[string]*cave, current *cave, visited map[string]bool, visitedTwice bool) {
	start := caves["start"]
	end := caves["end"]
	if pathCount > 40 {
		return
	}
	if current == end {
		pathCount++
		return
	}
	if current.isSmall {
		visited[current.name] = true
		current.visitedTimes++
		if current.visitedTimes >= 2 {
			current.visitedTimes--
			visitedTwice = true
		}
	}
	for _, next := range current.adjacent {
		if next == start || (next.isSmall && visited[next.name] && visitedTwice) {
			continue
		}
		fmt.Printf("pc: %d, From %s -> %s\n", pathCount, current.name, next.name)
		findPaths2(caves, next, visited, visitedTwice)

	}
	if current.isSmall {
		visited[current.name] = false
	}
}

func fs1(input io.Reader) int {
	caves = make(map[string]*cave)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		in := strings.Split(line, "-")
		if _, ok := caves[in[0]]; !ok {
			caves[in[0]] = &cave{
				name:    in[0],
				isSmall: in[0] == strings.ToLower(in[0]),
			}
		}
		if _, ok := caves[in[1]]; !ok {
			caves[in[1]] = &cave{
				name:    in[1],
				isSmall: in[1] == strings.ToLower(in[1]),
			}
		}
		caves[in[0]].adjacent = append(caves[in[0]].adjacent, caves[in[1]])
		caves[in[1]].adjacent = append(caves[in[1]].adjacent, caves[in[0]])
	}
	start := caves["start"]
	pathCount = 0
	visited = make(map[string]bool)
	findPaths(caves, start, visited)
	return pathCount
}

func fs2(input io.Reader) int {
	caves = make(map[string]*cave)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		in := strings.Split(line, "-")
		if _, ok := caves[in[0]]; !ok {
			caves[in[0]] = &cave{
				name:    in[0],
				isSmall: in[0] == strings.ToLower(in[0]),
			}
		}
		if _, ok := caves[in[1]]; !ok {
			caves[in[1]] = &cave{
				name:    in[1],
				isSmall: in[1] == strings.ToLower(in[1]),
			}
		}
		caves[in[0]].adjacent = append(caves[in[0]].adjacent, caves[in[1]])
		caves[in[1]].adjacent = append(caves[in[1]].adjacent, caves[in[0]])
	}
	start := caves["start"]
	pathCount = 0
	visited = make(map[string]bool)
	findPaths2(caves, start, visited, false)
	fmt.Printf("Total paths: %d\n", pathCount)
	// return pathCount
	return 42
}

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fs2(f)
}
