package main

import (
	"fmt"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	rows, err := util.ReadInts("16-1", false, ",")
	if err != nil {
		panic(err)
	}

	total := 0
	cols := 90
	for num := range rows {
		total += cols / num
	}

	fmt.Println(total)
}
