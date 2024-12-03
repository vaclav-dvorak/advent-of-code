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

func fs1(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	list1 := []int{}
	list2 := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "   ")
		if wi, e := strconv.Atoi(split[0]); e != nil {
			panic(e)
		} else {
			list1 = append(list1, wi)
		}
		if wi, e := strconv.Atoi(split[1]); e != nil {
			panic(e)
		} else {
			list2 = append(list2, wi)
		}
	}
	sort.Ints(list1)
	sort.Ints(list2)

	diff := 0
	for i, v1 := range list1 {
		v2 := list2[i]
		p := v1 - v2
		if p < 0 {
			diff -= p
		} else {
			diff += p
		}
	}
	return diff
}

func fs2(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	list1 := []int{}
	list2 := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "   ")
		if wi, e := strconv.Atoi(split[0]); e != nil {
			panic(e)
		} else {
			list1 = append(list1, wi)
		}
		if wi, e := strconv.Atoi(split[1]); e != nil {
			panic(e)
		} else {
			list2 = append(list2, wi)
		}
	}
	sum := 0
	for _, v1 := range list1 {
		count := 0
		for _, v2 := range list2 {
			if v1 == v2 {
				count++
			}
		}
		sum += v1 * count
	}
	return sum
}

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fs2(f)
}
