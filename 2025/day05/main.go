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

type stub struct {
	start int
	end   int
}

func fs1(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanningRanges := true
	ranges := make([]stub, 0)
	isFresh := 0
	for scanner.Scan() {
		line := scanner.Text()
		if scanningRanges {
			if line == "" {
				scanningRanges = false
				continue
			}
			split := strings.Split(line, "-")
			start, _ := strconv.Atoi(split[0])
			end, _ := strconv.Atoi(split[1])
			ranges = append(ranges, stub{start: start, end: end})
			continue
		}
		test, _ := strconv.Atoi(line)
		for _, r := range ranges {
			if test >= r.start && test <= r.end {
				isFresh++
				break
			}
		}
	}
	fmt.Printf("%d\n", isFresh)
	return isFresh
}

func fs2(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanningRanges := true
	ranges := make([]stub, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if scanningRanges {
			if line == "" {
				scanningRanges = false
				continue
			}
			split := strings.Split(line, "-")
			start, _ := strconv.Atoi(split[0])
			end, _ := strconv.Atoi(split[1])
			ranges = append(ranges, stub{start: start, end: end})
			continue
		}
	}
	sort.SliceStable(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})
	merged := make([]stub, 0)
	current := ranges[0]
	for i := 1; i < len(ranges); i++ {
		if ranges[i].start <= current.end+1 {
			if ranges[i].end > current.end {
				current.end = ranges[i].end
			}
		} else {
			merged = append(merged, current)
			current = ranges[i]
		}
	}
	merged = append(merged, current)
	ranges = merged
	fresh := 0
	for i := 0; i < len(ranges); i++ {
		fresh += ranges[i].end - ranges[i].start + 1
	}
	return fresh
}

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fs2(f)
}
