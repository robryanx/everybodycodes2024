package main

import (
	"fmt"
	"slices"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	rows, err := util.ReadIntLists("8-3", false, "\n")
	if err != nil {
		panic(err)
	}

	row := slices.Collect(rows)[0]

	nailCount := 256
	nodes := make(map[int][]int, len(row))
	for i := 0; i < len(row)-1; i++ {
		nodes[row[i]-1] = append(nodes[row[i]-1], row[i+1]-1)
		nodes[row[i+1]-1] = append(nodes[row[i+1]-1], row[i]-1)
	}

	best := 0
	for i := 0; i < nailCount; i++ {
		for j := i + 1; j < nailCount; j++ {
			count := 0

			between := inBetweenRange(i, j)
			for i := between[0] + 1; i < between[1]; i++ {
				for _, conn := range nodes[i] {
					if outsideRange(conn, between[0], between[1]) {
						count++
					}
				}
			}
			if node, ok := nodes[between[0]]; ok {
				if slices.Contains(node, between[1]) {
					count++
				}
			}

			if count > best {
				best = count
			}
		}
	}

	fmt.Println(best)
}

func outsideRange(val, low, high int) bool {
	return val < low || val > high
}

func inBetweenRange(a, b int) [2]int {
	return [2]int{min(a, b), max(a, b)}
}
