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

type report struct {
	levels []int
}

func safeReport(ls []int) (safe bool, problem int) {
	safe = true
	asc := false
	if ls[0] < ls[1] {
		asc = true
	}
	for i := 0; i < len(ls)-1; i++ {
		if asc && ls[i] > ls[i+1] { // should be ascending but is not
			safe = false
			problem = i + 1
			break
		}
		if !asc && ls[i] < ls[i+1] { // should be descending but is not
			safe = false
			problem = i + 1
			break
		}
		diff := ls[i] - ls[i+1]
		if diff < 0 {
			diff = -diff
		}
		if diff == 0 || diff > 3 {
			safe = false
			problem = i + 1
			break
		}
	}
	return
}

func findSafe(rs []report, modul bool) int {
	safeC := 0
	for _, r := range rs {
		safe, problem := safeReport(r.levels)
		if safe {
			safeC++
		} else {
			if modul {
				t1 := make([]int, len(r.levels))
				t2 := make([]int, len(r.levels))
				t3 := make([]int, len(r.levels))
				copy(t1, r.levels)
				copy(t2, r.levels)
				copy(t3, r.levels)
				tests := [][]int{
					t3[1:],                                  // remove first
					append(t1[:problem-1], t1[problem:]...), // remove left of problem
					append(t2[:problem], t2[problem+1:]...), // remove right of problem
				}
				for _, test := range tests {
					safe, _ = safeReport(test)
					if safe {
						safeC++
						break
					}
				}
			}
		}
	}
	fmt.Printf("Safe count: %d\n", safeC)
	return safeC
}

func fs1(input io.Reader) int {
	reports := make([]report, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		levS := strings.Split(line, " ")
		levels := make([]int, 0)
		for _, l := range levS {
			levI, _ := strconv.Atoi(l)
			levels = append(levels, levI)
		}
		reports = append(reports, report{
			levels: levels,
		})
	}
	return findSafe(reports, false)
}

func fs2(input io.Reader) int {
	reports := make([]report, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		levS := strings.Split(line, " ")
		levels := make([]int, 0)
		for _, l := range levS {
			levI, _ := strconv.Atoi(l)
			levels = append(levels, levI)
		}
		reports = append(reports, report{
			levels: levels,
		})
	}
	return findSafe(reports, true)
}

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fs2(f)
}
