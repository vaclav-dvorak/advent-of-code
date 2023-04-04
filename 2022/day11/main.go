package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	items        []int64
	opType       string
	opArg        int
	test         int64
	ifTrueThrow  int
	ifFalseThrow int
}

var (
	herd     []monkey
	herd_mod int64
)

func parseHerd(input io.Reader) {
	scanner := bufio.NewScanner(input)
	herd = make([]monkey, 10)
	herd_mod = 1
	uk := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		params := strings.Split(line, ":")
		instr := strings.Trim(params[0], " ")
		if len(instr) > 7 && instr[0:7] == "Monkey " {
			v, err := strconv.Atoi(instr[7:8])
			if err != nil {
				log.Fatal(err)
			}
			uk = v
			continue
		}
		switch instr {
		case "Starting items":
			items := strings.Split(strings.Trim(params[1], " "), ", ")
			for _, item := range items {
				v, err := strconv.ParseInt(item, 10, 64)
				if err != nil {
					log.Fatal(err)
				}
				herd[uk].items = append(herd[uk].items, v)
			}
		case "Operation":
			t := params[1][11:]
			herd[uk].opType = "m"
			if t == "* old" {
				herd[uk].opType = "s"
			} else {
				if t[:1] == "+" {
					herd[uk].opType = "a"
				}
				v, err := strconv.Atoi(t[2:])
				if err != nil {
					log.Fatal(err)
				}
				herd[uk].opArg = v
			}
		case "Test":
			v, err := strconv.ParseInt(params[1][14:], 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			herd[uk].test = v
			herd_mod *= v
		case "If true":
			v, err := strconv.Atoi(params[1][17:])
			if err != nil {
				log.Fatal(err)
			}
			herd[uk].ifTrueThrow = v
		case "If false":
			v, err := strconv.Atoi(params[1][17:])
			if err != nil {
				log.Fatal(err)
			}
			herd[uk].ifFalseThrow = v
		}
	}
}

func monkeyBusiness(rounds int, relief bool) int {
	inspections := make([]int, 10)
	for round := 0; round < rounds; round++ {
		for i, monk := range herd {
			for _, item := range monk.items {
				newWorry := int64(0)
				switch monk.opType {
				case "s":
					newWorry = item * item
				case "a":
					newWorry = item + int64(monk.opArg)
				case "m":
					newWorry = item * int64(monk.opArg)
				}
				if relief {
					newWorry = newWorry / 3 //? relief
				} else {
					newWorry = newWorry % herd_mod
				}
				if newWorry%monk.test == 0 {
					herd[monk.ifTrueThrow].items = append(herd[monk.ifTrueThrow].items, newWorry)
				} else {
					herd[monk.ifFalseThrow].items = append(herd[monk.ifFalseThrow].items, newWorry)
				}
				inspections[i]++
			}
			herd[i].items = make([]int64, 0)
		}
	}
	// fmt.Printf("%+v", herd)
	// fmt.Printf("%+v", inspections)
	sort.Slice(inspections, func(i, j int) bool { return inspections[i] > inspections[j] })
	return inspections[0] * inspections[1]
}

func fs1(input io.Reader) int {
	parseHerd(input)
	return monkeyBusiness(20, true)
}

func fs2(input io.Reader) int {
	parseHerd(input)
	return monkeyBusiness(10000, false)
}

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fs2(f)
}
