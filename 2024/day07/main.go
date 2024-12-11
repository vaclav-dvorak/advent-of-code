package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func testLine(nums []int, pow int, sum int) (valid bool) {
	valid = false
	opersMax := int(math.Pow(float64(pow), float64(len(nums)-1))) - 1
	stringRep := strconv.FormatInt(int64(opersMax), pow)
	if len(stringRep) > 19 {
		fmt.Printf("too many nums")
		panic(0)
	}
	for x := 0; x <= opersMax; x++ {
		mask := fmt.Sprintf("%020s", strconv.FormatInt(int64(x), pow))
		outcome := nums[0]
		for i, num := range nums {
			if i == 0 {
				continue
			}
			if mask[20-i] == '0' {
				outcome += num
			} else if mask[20-i] == '1' {
				outcome *= num
			} else {
				outcome, _ = strconv.Atoi(fmt.Sprintf("%d%d", outcome, num))
			}
		}
		if outcome == sum {
			valid = true
			break
		}
	}
	return
}

func fs1(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	validSum := 0
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ": ")
		sum, _ := strconv.Atoi(split[0])
		numsS := strings.Split(split[1], " ")
		nums := make([]int, len(numsS))
		for i, n := range numsS {
			nums[i], _ = strconv.Atoi(n)
		}
		if testLine(nums, 2, sum) {
			validSum += sum
		}
	}
	fmt.Printf("validSum: %d\n", validSum)
	return validSum
}

func fs2(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	validSum := 0
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ": ")
		sum, _ := strconv.Atoi(split[0])
		numsS := strings.Split(split[1], " ")
		nums := make([]int, len(numsS))
		for i, n := range numsS {
			nums[i], _ = strconv.Atoi(n)
		}
		if testLine(nums, 3, sum) {
			validSum += sum
		}
	}
	fmt.Printf("validSum: %d\n", validSum)
	return validSum
}

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fs2(f)
}
