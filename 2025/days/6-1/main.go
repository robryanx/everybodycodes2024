package main

import (
	"fmt"
	"slices"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	rows, err := util.ReadStrings("6-1", false, "\n")
	if err != nil {
		panic(err)
	}

	total := 0
	mentees := 0

	row := slices.Collect(rows)[0]
	for _, ch := range []byte(row) {
		if ch == 'A' {
			mentees++
		}

		if ch == 'a' {
			total += mentees
		}
	}

	fmt.Println(total)
}
