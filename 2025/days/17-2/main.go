package main

import (
	"fmt"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	rows, err := util.ReadStrings("17-2", false, "\n")
	if err != nil {
		panic(err)
	}

	var grid [][]byte
	for row := range rows {
		grid = append(grid, []byte(row))
	}

	radius := 1
	midY := len(grid) / 2
	midX := len(grid[0]) / 2

	best := 0
	bestRadius := 0

	for {
		total := 0
		for y := range len(grid) {
			for x := range len(grid[0]) {
				if grid[y][x] == '.' {
					continue
				}

				dist := ((midX - x) * (midX - x)) + ((midY - y) * (midY - y))
				if dist == 0 {
					continue
				}
				if dist <= radius*radius {
					total += int(grid[y][x] - '0')
					grid[y][x] = '.'
				}
			}
		}

		if total == 0 {
			break
		}
		if total > best {
			best = total
			bestRadius = radius
		}

		radius++
	}

	fmt.Println(best * bestRadius)
}
