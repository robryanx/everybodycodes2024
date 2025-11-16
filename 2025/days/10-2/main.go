package main

import (
	"fmt"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	rows, err := util.ReadStrings("10-2", false, "\n")
	if err != nil {
		panic(err)
	}

	var grid [][]byte
	for row := range rows {
		grid = append(grid, []byte(row))
	}

	maxY := len(grid) - 1
	maxX := len(grid[0]) - 1

	lookup := map[int]int{}

	var start [2]int
	sheep := map[int]struct{}{}
	hideout := map[int]struct{}{}
	for y := range grid {
		for x := range grid {
			switch grid[y][x] {
			case 'D':
				start = [2]int{y, x}
			case 'S':
				sheep[y*100+x] = struct{}{}
			case '#':
				hideout[y*100+x] = struct{}{}
			}
		}
	}

	total := move(start, maxX, maxY, 0, sheep, hideout, lookup, 0)

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

func move(curr [2]int, maxY, maxX, count int, sheep, hideout map[int]struct{}, lookup map[int]int, total int) int {
	if curr[0] < 0 ||
		curr[0] > maxY ||
		curr[1] < 0 ||
		curr[1] > maxX {
		return total
	}

	if checkCount, ok := lookup[curr[0]*100+curr[1]]; ok {
		if checkCount <= count {
			return total
		}
	}
	lookup[curr[0]*100+curr[1]] = count

	if count > 0 {
		if _, okHideout := hideout[curr[0]*100+curr[1]]; !okHideout {
			if _, ok := sheep[(curr[0]-count+1)*100+curr[1]]; ok {
				total++
				delete(sheep, (curr[0]-count+1)*100+curr[1])
			}
			if _, ok := sheep[(curr[0]-count)*100+curr[1]]; ok {
				total++
				delete(sheep, (curr[0]-count)*100+curr[1])
			}
		}
	}

	if count < 20 {
		for _, m := range moves {
			total = move([2]int{curr[0] + m[0], curr[1] + m[1]}, maxY, maxX, count+1, sheep, hideout, lookup, total)
		}
	}

	return total
}
