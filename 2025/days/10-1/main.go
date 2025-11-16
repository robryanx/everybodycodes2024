package main

import (
	"fmt"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	rows, err := util.ReadStrings("10-1", false, "\n")
	if err != nil {
		panic(err)
	}

	var grid [][]byte
	for row := range rows {
		grid = append(grid, []byte(row))
	}

	var start [2]int
	for y := range grid {
		for x := range grid {
			if grid[y][x] == 'D' {
				start = [2]int{y, x}
			}
		}
	}

	total := move(start, 0, grid, 0)

	fmt.Println(total)
}

var moves = [][2]int{
	{-2, -1},
	{-2, 1},
	{-1, -2},
	{-1, 2},
	{1, -2},
	{1, 2},
	{2, -1},
	{2, 1},
}

func move(curr [2]int, count int, grid [][]byte, total int) int {
	if curr[0] < 0 ||
		curr[0] > len(grid)-1 ||
		curr[1] < 0 ||
		curr[1] > len(grid[0])-1 {
		return total
	}

	if grid[curr[0]][curr[1]] == 'S' {
		total++
	}
	grid[curr[0]][curr[1]] = 'X'

	if count < 4 {
		for _, m := range moves {
			total = move([2]int{curr[0] + m[0], curr[1] + m[1]}, count+1, grid, total)
		}
	}

	return total
}
