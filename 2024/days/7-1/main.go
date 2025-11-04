package main

import (
	"cmp"
	"fmt"
	"slices"

	"github.com/robryanx/everybodycodes/util"
)

type device struct {
	id    byte
	power int
}

func main() {
	rows, err := util.ReadStrings("7-1", false, "\n")
	if err != nil {
		panic(err)
	}

	results := make([]device, 0, 100)
	for row := range rows {
		total := 0
		current := 10
		id := row[0]
		for i := 2; i < len(row); i += 2 {
			if row[i] == '+' {
				current++
			} else if row[i] == '-' {
				current--
			}

			total += current
		}

		results = append(results, device{
			id:    id,
			power: total,
		})
	}

	slices.SortFunc(results, func(a, b device) int {
		return cmp.Compare(b.power, a.power)
	})

	for _, r := range results {
		fmt.Print(string(r.id))
	}
	fmt.Println()
}
