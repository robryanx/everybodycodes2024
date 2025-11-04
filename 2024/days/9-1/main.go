package main

import (
	"fmt"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	rows, err := util.ReadInts("9-1", false, "\n")
	if err != nil {
		panic(err)
	}

	stamps := []int{10, 5, 3, 1}

	beetles := 0
	for sparkball := range rows {
		for sparkball > 0 {
			for i := 0; i < len(stamps); i++ {
				if sparkball >= stamps[i] {
					beetles++
					sparkball -= stamps[i]
					break
				}
			}
		}
	}

	fmt.Println(beetles)
}
