package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

func processLine(in string) (out int) {
	var uniq map[int]bool = map[int]bool{2: true, 4: true, 3: true, 7: true}
	split := strings.Split(in, " | ")
	digits := strings.Split(split[1], " ")

	for _, digit := range digits {
		l := len(digit)
		if found := uniq[l]; found {
			out++
		}
	}
	return
}

func containsSignals(hay string, needle string) (out bool) {
	out = false
	haySl := strings.Split(hay, "")
	needleSl := strings.Split(needle, "")
	for _, search := range needleSl {
		found := false
		for _, hays := range haySl {
			if search == hays {
				found = true
				continue
			}
		}
		if !found {
			return
		}
	}
	out = true
	return
}

func sortSignals(in []string) (out []string) {
	for _, s := range in {
		sp := strings.Split(s, "")
		sort.Strings(sp)
		out = append(out, strings.Join(sp, ""))
	}
	return
}

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func solveNumber(signals []string, l int, c string) (out []string, s string) {
	for i, s := range signals {
		if c == "" && len(s) == l {
			out := remove(signals, i)
			return out, s
		}
		if c != "" && len(s) == l && containsSignals(s, c) {
			out := remove(signals, i)
			return out, s
		}
	}
	fmt.Print("ERR - number not found!\n")
	return signals, ""
}

func solveLine(in string) (out int) {
	signalMap := make(map[int]string)
	split := strings.Split(in, " | ")
	signals := strings.Split(split[0], " ")
	signals = sortSignals(signals)
	signals, signalMap[1] = solveNumber(signals, 2, "")           // get string for 1
	signals, signalMap[4] = solveNumber(signals, 4, "")           // get string for 4
	signals, signalMap[8] = solveNumber(signals, 7, "")           // get string for 8
	signals, signalMap[7] = solveNumber(signals, 3, "")           // get string for 7
	signals, signalMap[9] = solveNumber(signals, 6, signalMap[4]) // get string for 9
	signals, signalMap[0] = solveNumber(signals, 6, signalMap[1]) // get string for 0
	signals, signalMap[6] = solveNumber(signals, 6, "")           // get string for 6
	signals, signalMap[3] = solveNumber(signals, 5, signalMap[1]) // get string for 3
	// only 5 and 2 is left... 5 is contained in 6
	if containsSignals(signalMap[6], signals[0]) {
		// if this is true 0 index is five
		signalMap[5] = signals[0]
		signalMap[2] = signals[1]
	} else {
		signalMap[2] = signals[0]
		signalMap[5] = signals[1]
	}

	transpileMap := make(map[string]int)
	for i, s := range signalMap {
		transpileMap[s] = i
	}
	number := strings.Split(split[1], " ")
	number = sortSignals(number)
	out = transpileMap[number[0]]*1000 + transpileMap[number[1]]*100 + transpileMap[number[2]]*10 + transpileMap[number[3]]

	return
}

func fs1(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += processLine(line)
	}
	return sum
}

func fs2(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += solveLine(line)
	}
	fmt.Printf("solved: %d", sum)
	return sum
}

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fs2(f)
}
