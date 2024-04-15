package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

type chunk struct {
	open  rune
	close rune
}

var score map[rune]int

func checkLine(in string) (ok bool, pt int) {
	ok = true
	score = map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	var curChunk chunk
	var openChunks []chunk
	for _, r := range in {
		cl := 'e'
		switch r {
		case '(':
			cl = ')'
		case '[':
			cl = ']'
		case '{':
			cl = '}'
		case '<':
			cl = '>'
		}
		if cl == 'e' { // we are closing chunk
			if r == curChunk.close { // chunk closed
				openChunks = openChunks[:len(openChunks)-1]
				if len(openChunks) > 0 {
					curChunk = openChunks[len(openChunks)-1]
				}

			} else { //corrupted line
				pt = score[r]
				ok = false
				return
			}
		} else {
			openChunks = append(openChunks, chunk{
				open:  r,
				close: cl,
			})
			curChunk = openChunks[len(openChunks)-1]
		}
	}
	return
}

func checkLine2(in string) (pt int) {
	score = map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
	var curChunk chunk
	var openChunks []chunk
	for _, r := range in {
		cl := 'e'
		switch r {
		case '(':
			cl = ')'
		case '[':
			cl = ']'
		case '{':
			cl = '}'
		case '<':
			cl = '>'
		}
		if cl == 'e' { // we are closing chunk
			if r == curChunk.close { // chunk closed
				openChunks = openChunks[:len(openChunks)-1]
				if len(openChunks) > 0 {
					curChunk = openChunks[len(openChunks)-1]
				}

			} else { //corrupted line
				return
			}
		} else {
			openChunks = append(openChunks, chunk{
				open:  r,
				close: cl,
			})
			curChunk = openChunks[len(openChunks)-1]
		}
	}
	for len(openChunks) > 0 {
		pop := openChunks[len(openChunks)-1]
		openChunks = openChunks[:len(openChunks)-1]
		pt = pt*5 + score[pop.close]
	}
	return
}

func fs1(input io.Reader) int {
	sum := 0
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		if ok, pt := checkLine(line); !ok {
			sum += pt
		}
	}

	return sum
}

func fs2(input io.Reader) int {
	var pts []int
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		pt := checkLine2(line)
		if pt != 0 { //we will add it to array and then find middle value
			pts = append(pts, pt)
			fmt.Printf("score : %d\n", pt)
		}
	}

	sort.Ints(pts)
	return pts[len(pts)/2 : (len(pts)/2)+1][0]
}

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fs2(f)
}
