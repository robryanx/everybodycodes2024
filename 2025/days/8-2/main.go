package main

import (
	"fmt"
	"slices"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	rows, err := util.ReadIntLists("8-2", false, "\n")
	if err != nil {
		panic(err)
	}

	row := slices.Collect(rows)[0]

	total := 0
	nailCount := 256
	nodes := make(map[int][]int, len(row))
	for i := 0; i < len(row)-1; i++ {
		curr := row[i] - 1
		next := row[i+1] - 1

		between := inBetweenRange(curr, next)
		for i := between[0] + 1; i < between[1]; i++ {
			for _, conn := range nodes[i%nailCount] {
				if conn != curr && conn != next && outsideRange(conn, between[0], between[1]) {
					total++
				}
			}
		}

		nodes[curr] = append(nodes[curr], next)
		nodes[next] = append(nodes[next], curr)
	}

	fmt.Println(total)
}

func outsideRange(val, low, high int) bool {
	return val < low || val > high
}

func inBetweenRange(a, b int) [2]int {
	return [2]int{min(a, b), max(a, b)}
}
