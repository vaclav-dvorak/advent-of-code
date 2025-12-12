package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func pow10(i int) int {
	if i == 0 {
		return 1
	}
	if i == 1 {
		return 10
	}
	if i == 2 {
		return 100
	}
	if i == 3 {
		return 1000
	}
	if i == 4 {
		return 10000
	}
	if i == 5 {
		return 100000
	}
	if i == 6 {
		return 1000000
	}
	if i == 7 {
		return 10000000
	}
	if i == 8 {
		return 100000000
	}
	if i == 9 {
		return 1000000000
	}
	if i == 10 {
		return 10000000000
	}
	if i == 11 {
		return 100000000000
	}
	panic("too big")
}

func fs1(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	bankSum := 0
	for scanner.Scan() {
		line := scanner.Text()
		bmax1 := 0
		bmax2 := 0
		bmax := 0
		for p, char := range line {
			bmax = bmax1*10 + bmax2
			i := int(char - '0')
			if bmax1 == 0 {
				bmax1 = int(i)
				bmax2 = int(line[p+1] - '0')
				continue
			}

			found := false
			if i > bmax1 {
				for j, tChar := range line {
					if j <= p {
						continue
					}
					j := int(tChar - '0')
					if i*10+j > bmax {
						bmax1 = i
						bmax2 = j
						found = true
					}
				}
			}
			if found {
				continue
			}
			if i > bmax2 {
				bmax2 = i
			}
		}
		bankSum += bmax1*10 + bmax2
	}

	return bankSum
}

func getMaxPos(line string, start int, end int) (int, int) {
	max := -1
	pos := -1
	for p := start; p < end; p++ {
		i := int(line[p] - '0')
		if i > max {
			max = i
			pos = p
		}
	}
	return max, pos
}

func fs2(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	bankSum := 0
	for scanner.Scan() {
		line := scanner.Text()
		l := len(line)
		linemax := 0
		p := 0
		maxLen := 12
		for maxLen > 0 {
			max, pos := getMaxPos(line, p, l-maxLen+1)
			p = pos + 1
			maxLen--
			linemax += max * pow10(maxLen)
		}
		bankSum += linemax
	}
	return bankSum
}

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fs2(f)
}
