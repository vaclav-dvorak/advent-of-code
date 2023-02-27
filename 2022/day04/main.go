package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	overlaps := 0
	for i := 0; i < len(input); i++ {
		split := strings.Split(input[i], ",")
		fi := strings.Split(split[0], "-")
		fl, _ := strconv.Atoi(fi[0])
		fh, _ := strconv.Atoi(fi[1])

		se := strings.Split(split[1], "-")
		sl, _ := strconv.Atoi(se[0])
		sh, _ := strconv.Atoi(se[1])

		if (fl >= sl && fl <= sh) || (fh >= sl && fh <= sh) || (sl >= fl && sl <= fh) || (sh >= fl && sh <= fh) {
			overlaps++
		}
	}
	fmt.Printf("%d\n", overlaps)
}
