package main

import (
	"fmt"
	"slices"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	rows, err := util.ReadStrings("6-2", false, "\n")
	if err != nil {
		panic(err)
	}

	total := 0
	menteesA := 0
	menteesB := 0
	menteesC := 0

	row := slices.Collect(rows)[0]
	for _, ch := range []byte(row) {
		if ch == 'A' {
			menteesA++
		}
		if ch == 'a' {
			total += menteesA
		}
		if ch == 'B' {
			menteesB++
		}
		if ch == 'b' {
			total += menteesB
		}
		if ch == 'C' {
			menteesC++
		}
		if ch == 'c' {
			total += menteesC
		}
	}

	fmt.Println(total)
}
