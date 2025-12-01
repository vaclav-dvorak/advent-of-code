package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

type cave struct {
	name     string
	isSmall  bool
	adjacent []*cave
}

var pathCount int

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

func findPaths2(caves map[string]*cave, current *cave, visited map[string]int) {
	start := caves["start"]
	end := caves["end"]
	if current == end {
		pathCount++
		return
	}
	if current.isSmall {
		visited[current.name]++
	}

	// Check if any small cave has been visited twice in this path
	hasTwiceVisited := false
	for name, count := range visited {
		if caves[name].isSmall && count == 2 {
			hasTwiceVisited = true
			break
		}
	}

	for _, next := range current.adjacent {
		if next == start {
			continue
		}
		// If we've already used our "double visit" on one small cave,
		// we can't visit any other small cave that's already been visited
		if next.isSmall {
			if hasTwiceVisited && visited[next.name] >= 1 {
				continue
			}
			if !hasTwiceVisited && visited[next.name] >= 2 {
				continue
			}
		}
		findPaths2(caves, next, visited)
	}
	if current.isSmall {
		visited[current.name]--
	}
}

func readInput(input io.Reader) (caves map[string]*cave) {
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
	return
}

func fs1(input io.Reader) int {
	caves := readInput(input)
	pathCount = 0
	findPaths(caves, caves["start"], make(map[string]bool))
	return pathCount
}

func fs2(input io.Reader) int {
	caves := readInput(input)
	pathCount = 0
	findPaths2(caves, caves["start"], make(map[string]int))
	return pathCount
}

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fs2(f)
}
