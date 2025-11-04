package main

import (
	"fmt"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	strs, err := util.ReadBytes("1-1", false)
	if err != nil {
		panic(err)
	}

	sum := 0

	for _, char := range strs {
		switch char {
		case 'B':
			sum++
		case 'C':
			sum += 3
		}
	}

	fmt.Println(sum)
}
