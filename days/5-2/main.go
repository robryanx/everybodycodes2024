package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/robryanx/everybodycodes2024/util"
)

func main() {
	rows, err := util.ReadStrings("5-2", false, "\n")
	if err != nil {
		panic(err)
	}

	var gridRaw [][]int
	for row := range rows {
		nums := strings.Split(row, " ")
		rowBuild := make([]int, 0, len(nums))
		for _, num := range nums {
			value, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			rowBuild = append(rowBuild, value)
		}

		gridRaw = append(gridRaw, rowBuild)
	}

	var grid [][]int
	for x := 0; x < len(gridRaw[0]); x++ {
		col := make([]int, 0, len(gridRaw))
		for y := 0; y < len(gridRaw); y++ {
			col = append(col, gridRaw[y][x])
		}

		grid = append(grid, col)
	}

	shoutedCount := make(map[string]int, 0)
	for x := 0; ; x++ {
		current := x % len(grid)
		next := (x + 1) % len(grid)

		clapper := grid[current][0]
		wrappedPos := clapper % (len(grid[next]) * 2)
		if wrappedPos == 0 {
			wrappedPos = (len(grid[next]) * 2)
		}

		grid[current] = grid[current][1:]

		if wrappedPos <= len(grid[next]) {
			grid[next] = slices.Insert(grid[next], wrappedPos-1, clapper)
		} else {
			pos := wrappedPos - len(grid[next])
			pos = len(grid[next]) - pos

			grid[next] = slices.Insert(grid[next], pos+1, clapper)
		}

		var shouted string
		for y := 0; y < len(grid); y++ {
			shouted += strconv.Itoa(grid[y][0])
		}

		if _, ok := shoutedCount[shouted]; !ok {
			shoutedCount[shouted] = 1
		} else {
			shoutedCount[shouted]++
		}

		if shoutedCount[shouted] == 2024 {
			value, err := strconv.Atoi(shouted)
			if err != nil {
				panic(err)
			}

			fmt.Println((x + 1) * value)
			break
		}
	}
}
