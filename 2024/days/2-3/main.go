package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	rowsIter, err := util.ReadStrings("2-3", false, "\n")
	if err != nil {
		panic(err)
	}

	allRows := slices.Collect(rowsIter)

	grid := gridFromRows(allRows[2:])

	wordsStr := allRows[0][strings.Index(allRows[0], ":")+1:]

	allIndexes := make(map[string]struct{})
	for _, word := range strings.Split(wordsStr, ",") {
		for y := 0; y < len(grid); y++ {
			row := slices.Clone(grid[y])
			row = append(row, row[:len(word)-1]...)
			indexes := strIndexes(row, word)
			for k := range indexes {
				allIndexes[fmt.Sprintf("%d-%d", y, k%len(grid[y]))] = struct{}{}
			}

			rowReverse := slices.Clone(grid[y])
			rowReverse = reverse(rowReverse)
			rowReverse = append(rowReverse, rowReverse[:len(word)-1]...)
			indexes = strIndexes(rowReverse, word)
			for k := range indexes {
				allIndexes[fmt.Sprintf("%d-%d", y, (len(grid[y])-(k%len(grid[y]))-1))] = struct{}{}
			}
		}

		for x := 0; x < len(grid[0]); x++ {
			col := column(grid, x)
			column := slices.Clone(col)
			indexes := strIndexes(column, word)
			for k := range indexes {
				allIndexes[fmt.Sprintf("%d-%d", k, x)] = struct{}{}
			}

			columnReverse := slices.Clone(col)
			indexes = strIndexes(reverse(columnReverse), word)
			for k := range indexes {
				allIndexes[fmt.Sprintf("%d-%d", len(col)-k-1, x)] = struct{}{}
			}
		}
	}

	fmt.Println(len(allIndexes))
}

func column(grid [][]byte, index int) []byte {
	col := make([]byte, 0, len(grid))
	for y := 0; y < len(grid); y++ {
		col = append(col, grid[y][index])
	}
	return col
}

func gridFromRows(rows []string) [][]byte {
	var grid [][]byte
	for _, row := range rows {
		grid = append(grid, []byte(row))
	}

	return grid
}

func reverse(s []byte) []byte {
	out := make([]byte, 0, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		out = append(out, s[i])
	}
	return out
}

func strIndexes(str []byte, word string) map[int]struct{} {
	indexes := make(map[int]struct{}, len(str))

	for i := 0; i <= len(str)-len(word); i++ {
		if str[i] == word[0] {
			match := true
			for j := i + 1; j < i+len(word); j++ {
				if str[j] != word[j-i] {
					match = false
					break
				}
			}

			if match {
				for j := i; j < i+len(word); j++ {
					indexes[j] = struct{}{}
				}
			}
		}
	}

	return indexes
}
