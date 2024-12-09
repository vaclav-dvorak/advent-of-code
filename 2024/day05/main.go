package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type protocol struct {
	before []int
	after  []int
}

func checkManual(man []int, protos map[int]protocol) (valid bool) {
	valid = true
	for i, v := range man {
		if proto, ok := protos[v]; ok {
			befo := man[:i]
			afte := man[i+1:]
			for _, b := range befo {
				if slices.Contains(proto.after, b) {
					valid = false
					break
				}
			}
			for _, a := range afte {
				if slices.Contains(proto.before, a) {
					valid = false
					break
				}
			}
		}
		if !valid {
			break
		}
	}
	return valid
}

func correctManual(man []int, protos map[int]protocol) []int {
	for i, val := range man {
		if proto, ok := protos[val]; ok {
			befo := man[:i]
			afte := man[i+1:]
			for f, befVal := range befo {
				if slices.Contains(proto.after, befVal) {
					man[f] = val
					man[i] = befVal
					break
				}
			}
			for f, afteVal := range afte {
				if slices.Contains(proto.before, afteVal) {
					man[f+i+1] = val
					man[i] = afteVal
					break
				}
			}
		}
	}
	return man
}

func fs1(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	protos := make(map[int]protocol)
	redingProtos := true
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			redingProtos = false
			continue
		}
		if redingProtos {
			sp := strings.Split(line, "|")
			bef, _ := strconv.Atoi(sp[0])
			aft, _ := strconv.Atoi(sp[1])
			if _, ok := protos[bef]; !ok {
				protos[bef] = protocol{}
			}
			if _, ok := protos[aft]; !ok {
				protos[aft] = protocol{}
			}
			tmp := protos[bef]
			tmp.after = append(tmp.after, aft)
			protos[bef] = tmp
			tmp = protos[aft]
			tmp.before = append(tmp.before, bef)
			protos[aft] = tmp
		} else {
			manS := strings.Split(line, ",")
			manI := make([]int, 0)
			for _, item := range manS {
				v, _ := strconv.Atoi(item)
				manI = append(manI, v)
			}
			if checkManual(manI, protos) {
				mid := (len(manI) / 2)
				sum += manI[mid]
			}
		}
	}

	fmt.Printf("%d\n", sum)
	return sum
}

func fs2(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	protos := make(map[int]protocol)
	redingProtos := true
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			redingProtos = false
			continue
		}
		if redingProtos {
			sp := strings.Split(line, "|")
			bef, _ := strconv.Atoi(sp[0])
			aft, _ := strconv.Atoi(sp[1])
			if _, ok := protos[bef]; !ok {
				protos[bef] = protocol{}
			}
			if _, ok := protos[aft]; !ok {
				protos[aft] = protocol{}
			}
			tmp := protos[bef]
			tmp.after = append(tmp.after, aft)
			protos[bef] = tmp
			tmp = protos[aft]
			tmp.before = append(tmp.before, bef)
			protos[aft] = tmp
		} else {
			manS := strings.Split(line, ",")
			manI := make([]int, 0)
			for _, item := range manS {
				v, _ := strconv.Atoi(item)
				manI = append(manI, v)
			}
			if checkManual(manI, protos) { //if its fine we dont care about it
				continue
			}
			c := manI
			n := 0
			for !checkManual(c, protos) {
				n++
				if n > 100 { //brute force protection against infinite loop
					panic(0)
				}

				c = correctManual(c, protos)
			}
			mid := (len(c) / 2)
			sum += c[mid]
		}
	}

	fmt.Printf("%d\n", sum)
	return sum
}

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fs2(f)
}
