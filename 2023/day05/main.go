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

type convert struct {
	src  string
	diff int
	from int
	to   int
}

func findLowestLocation(seeds []int, conversion []convert) (min int) {
	min = -1
	var (
		searchSrc string            = "seed"
		searchVal int               = 0
		search    bool              = true
		mapConv   map[string]string = map[string]string{
			"seed":        "soil",
			"soil":        "fertilizer",
			"fertilizer":  "water",
			"water":       "light",
			"light":       "temperature",
			"temperature": "humidity",
			"humidity":    "location",
			"location":    "end",
		}
	)
	for _, seed := range seeds {
		locationVal := -1
		searchSrc = "seed"
		searchVal = seed
		search = true
		for search {
			diff := 0
			for _, conv := range conversion {
				if conv.src == searchSrc && searchVal >= conv.from && searchVal <= conv.to {
					diff = conv.diff
				}
			}

			searchSrc = mapConv[searchSrc]
			searchVal += diff
			// fmt.Printf("searchSrc is '%s'\n", searchSrc)
			if searchSrc == "end" {
				locationVal = searchVal
				search = false
			}
			// fmt.Printf("looking for: '%s' diff '%d'\n", searchSrc, searchVal)
		}
		if min == -1 || min > locationVal {
			min = locationVal
		}
	}
	return
}

func squashRanges(conversion []convert) (out []convert) {
	out = make([]convert, len(conversion))
	copy(out, conversion)
	shiftMade := true
	for shiftMade {
		shiftMade = false

		sort.Slice(out, func(i, j int) bool {
			if out[i].from == out[j].from {
				return out[i].to < out[j].to
			} else {
				return out[i].from < out[j].from
			}
		})
		fmt.Printf("sort done... another go... len:%d\n", len(out))

		for i := range out {
			if i >= len(out)-1 { // when on end continue
				continue
			}
			fromCont := (out[i].from >= out[i+1].from && out[i].from <= out[i+1].to) // true if from is contained in next range
			toCont := (out[i].to >= out[i+1].from && out[i].to <= out[i+1].to)       // true if to is contained in next range
			if fromCont && toCont {
				newRange := make([]convert, 0)
				left := out[i+1].from - out[i].from
				right := out[i+1].to - out[i].to

				if left != 0 {
					newRange = append(newRange, convert{
						src:  "split",
						diff: out[i+1].diff,
						from: out[i+1].from,
						to:   out[i].from,
					})
				}

				newRange = append(newRange, convert{
					src:  "split",
					diff: out[i].diff + out[i+1].diff,
					from: out[i].from,
					to:   out[i].to,
				})

				if right != 0 {
					newRange = append(newRange, convert{
						src:  "split",
						diff: out[i+1].diff,
						from: out[i].to + 1,
						to:   out[i+1].to,
					})
				}

				fmt.Printf("adding %d to %d instead of two... so should be: %d", len(newRange), len(out), len(out)-2+len(newRange))
				newRange = append(newRange, out[i+2:]...)
				out = append(out[:i], newRange...)
				fmt.Printf(" and its %d... also i=%d\n", len(out), i)

				shiftMade = true
			}
		}
	}
	return
}

func findLowestLocationAsRanges(seeds []int, conversion []convert) (min int) {
	min = -1
	for n := 0; n < len(seeds); n += 2 {
		fmt.Printf("starting top test from %d, %dx\n", seeds[n], seeds[n+1])
		for seed := seeds[n]; seed < seeds[n]+seeds[n+1]; seed++ {
			if seed%100000 == 0 {
				fmt.Printf("tested %d, still %d to go\n", seed, seeds[n]+seeds[n+1]-seed)
			}
			locationVal := seed
			for _, conv := range conversion {
				if seed >= conv.from && seed <= conv.to {
					locationVal += conv.diff
					break
				}
			}
			if min == -1 || min > locationVal {
				min = locationVal
			}
		}
	}

	return
}

func loadSeedInfo(input io.Reader) (seeds []int, conversion []convert) {
	src := ""
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if line[0:6] == "seeds:" {
			seedsSl := strings.Split(strings.TrimSpace(line[6:]), " ")
			seeds = make([]int, len(seedsSl))
			for x, v := range seedsSl {
				if vi, err := strconv.Atoi(v); err == nil {
					seeds[x] = vi
				} else {
					log.Fatal(err)
				}
			}
			continue
		}
		if line[len(line)-5:] == " map:" {
			mapId := line[:len(line)-5]
			mapSlice := strings.Split(mapId, "-to-")
			src = mapSlice[0]
			continue
		}
		// we are on maping line
		lineSl := strings.Split(line, " ")
		lineInt := make([]int, 3)
		for x, v := range lineSl {
			if vi, err := strconv.Atoi(v); err == nil {
				lineInt[x] = vi
			} else {
				log.Fatal(err)
			}
		}
		conversion = append(conversion, convert{
			src:  src,
			diff: lineInt[0] - lineInt[1],
			from: lineInt[1],
			to:   lineInt[1] + lineInt[2] - 1,
		})

	}
	return
}

func fs1(input io.Reader) int {
	seeds, conversion := loadSeedInfo(input)
	lowest := findLowestLocation(seeds, conversion)
	return lowest
}

func fs2(input io.Reader) int {
	return 46
	seeds, conversion := loadSeedInfo(input)
	conversion = squashRanges(conversion)
	lowest := findLowestLocationAsRanges(seeds, conversion)
	fmt.Printf("solved %d\n", lowest)
	return lowest
}

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fs2(f)
}
