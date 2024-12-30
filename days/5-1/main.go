package main

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/robryanx/everybodycodes2024/util"
)

func main() {
	strs, err := util.ReadStrings("5-1", false, "\n")
	if err != nil {
		panic(err)
	}

	var gridRaw [][]int
	for str := range strs {
		row := make([]int, 0, len(str)/2)
		for i := 0; i < len(str); i += 2 {
			value, err := strconv.Atoi(string(str[i]))
			if err != nil {
				panic(err)
			}
			row = append(row, value)
		}

		gridRaw = append(gridRaw, row)
	}

	var grid [][]int
	for x := 0; x < len(gridRaw[0]); x++ {
		col := make([]int, 0, len(gridRaw))
		for y := 0; y < len(gridRaw); y++ {
			col = append(col, gridRaw[y][x])
		}

		grid = append(grid, col)
	}

	rounds := 10
	for x := 0; x < rounds; x++ {
		current := x % len(grid)
		next := (x + 1) % len(grid)

		clapper := grid[current][0]
		grid[current] = grid[current][1:]

		if clapper <= len(grid[next]) {
			grid[next] = slices.Insert(grid[next], clapper-1, clapper)
		} else {
			pos := clapper - len(grid[next])
			pos = len(grid[next]) - pos

			grid[next] = slices.Insert(grid[next], pos+1, clapper)
		}
	}

	var shouted string
	for y := 0; y < len(grid); y++ {
		shouted += strconv.Itoa(grid[y][0])
	}

	fmt.Println(shouted)
}
