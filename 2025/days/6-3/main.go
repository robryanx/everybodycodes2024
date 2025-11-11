package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	rows, err := util.ReadStrings("6-3", false, "\n")
	if err != nil {
		panic(err)
	}

	total := 0
	mentees := [3]int{0, 0, 0}

	row := slices.Collect(rows)[0]
	segmentLength := len(row)
	segmentTotal := 0

	repeats := 1000
	window := 1000
	row = strings.Repeat(row, 3)

	// build the initial window
	for i := 0; i < window; i++ {
		if row[i] >= 'A' && row[i] <= 'C' {
			mentees[int(row[i]-'A')]++
		}
	}

	for i := 0; i < len(row); i++ {
		if i-window-1 >= 0 {
			trailing := row[i-window-1]
			if trailing >= 'A' && trailing <= 'C' {
				mentees[int(trailing-'A')]--
			}
		}

		if i+window < len(row) {
			preceding := row[i+window]
			if preceding >= 'A' && preceding <= 'C' {
				mentees[int(preceding-'A')]++
			}
		}

		if row[i] >= 'a' && row[i] <= 'c' {
			if i >= segmentLength && i < segmentLength*2 {
				segmentTotal += mentees[int(row[i]-'a')]
			}

			total += mentees[int(row[i]-'a')]
		}
	}

	total += segmentTotal * (repeats - 3)

	fmt.Println(total)
}
