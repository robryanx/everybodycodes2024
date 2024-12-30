package main

import (
	"fmt"
	"math"
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
		col := make([]int, 0, len(gridRaw)+1)
		for y := 0; y < len(gridRaw); y++ {
			col = append(col, gridRaw[y][x])
		}

		grid = append(grid, col)
	}

	shoutedCount := make(map[int]int, 0)
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

		var shouted int
		for y := len(grid) - 1; y >= 0; y-- {
			shouted += int(math.Pow(10, float64(digits(shouted)))) * grid[y][0]
		}

		shoutedCount[shouted] += 1
		if shoutedCount[shouted] == 2024 {
			fmt.Println((x + 1) * shouted)
			break
		}
	}
}

func digits(num int) int {
	count := 0
	for num > 0 {
		num = num / 10
		count++
	}

	return count
}
