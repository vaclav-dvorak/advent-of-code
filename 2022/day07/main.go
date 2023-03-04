package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	curDir string
	files  map[string]int
	dirs   map[string]int
)

func parseInput(input io.Reader) {
	files = make(map[string]int)
	dirs = make(map[string]int)
	dirs["/"] = 0
	curDir = ""
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		if line[0] == '$' {
			line = line[2:]
			cmd := line[:2]
			switch cmd {
			case "cd":
				param := line[3:]
				if param == ".." {
					i := strings.LastIndex(curDir, "/")
					curDir = curDir[:i]
					continue
				}
				if param == "/" {
					curDir = ""
					continue
				}
				curDir += "/" + param
			case "ls":
			}
			continue
		}
		split := strings.Split(line, " ")
		size, name := split[0], split[1]
		if size == "dir" {
			dirs[curDir+"/"+name+"/"] = 0
		} else {
			sizeI, err := strconv.Atoi(size)
			if err != nil {
				log.Fatal(err)
			}
			files[curDir+"/"+name] = sizeI
		}
	}
}

func fs1(input io.Reader) int {
	parseInput(input)
	cut := 100000
	ret := 0
	for dir := range dirs {
		sum := 0
		for file, fSize := range files {
			if len(file) >= len(dir) && file[:len(dir)] == dir {
				sum += fSize
			}
		}
		if sum <= cut {
			ret += sum
		}
	}
	return ret
}

func fs2(input io.Reader) int {
	parseInput(input)
	diskSize := 70000000
	need := 30000000
	ret := 0
	rootSize := 0
	for _, fSize := range files {
		rootSize += fSize
	}
	need2free := need - diskSize + rootSize
	min := diskSize
	for dir := range dirs {
		sum := 0
		for file, fSize := range files {
			if len(file) >= len(dir) && file[:len(dir)] == dir {
				sum += fSize
			}
		}
		if sum > need2free && (sum-need2free) < min {
			min = sum - need2free
			// fmt.Printf("smallest %s\n", dir)
			ret = sum
		}
	}
	return ret
}

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fs2(f)
}
