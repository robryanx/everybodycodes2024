package main

import (
	"fmt"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	rows, err := util.ReadInts("9-3", false, "\n")
	if err != nil {
		panic(err)
	}

	stamps := []int{101, 100, 75, 74, 50, 49, 38, 37, 30, 25, 24, 20, 16, 15, 10, 5, 3, 1}

	shortestCache := make(map[int]int, 1024)

	total := 0
	for sparkball := range rows {
		smallest := sparkball
		sparkballMin := sparkball/2 - 50
		sparkballMax := sparkball/2 + 50
		if sparkball%2 == 1 {
			sparkballMin++
		}

		counts := make(map[int]int, 100)
		for i := sparkballMin; i <= sparkballMax; i++ {
			path := 0
			shortest := i
			base := i/stamps[0] - 12
			adj := i - base*stamps[0]
			rec(stamps, 0, adj, adj, path, &shortest, shortestCache)
			counts[i] = shortest + base
		}

		for i := sparkballMin; i <= sparkball/2; i++ {
			if counts[i]+counts[sparkball-i] < smallest {
				smallest = counts[i] + counts[sparkball-i]
			}
		}

		total += smallest
	}

	fmt.Println(total)
}

func rec(stamps []int, stampOffset, startingSparkball, sparkball int, path int, shortest *int, shortestCache map[int]int) {
	if stampOffset == len(stamps) || sparkball < 0 {
		return
	}

	// should we skip this branch as being definitely longer
	if stampOffset > 0 {
		for i := stampOffset - 1; i >= 0; i-- {
			if stamps[i]%stamps[stampOffset] == 0 && sparkball-stamps[i] > 0 {
				return
			}
		}
	}

	key := ((startingSparkball - sparkball) * 1000) + stamps[stampOffset]

	if sparkball > 0 {
		if shortestPath, ok := shortestCache[key]; ok {
			if path > shortestPath {
				return
			}
		}
		if path > 0 {
			shortestCache[key] = path
		}
		if path+1 >= *shortest {
			return
		}
		rec(stamps, stampOffset, startingSparkball, sparkball-stamps[stampOffset], path+1, shortest, shortestCache)
		rec(stamps, stampOffset+1, startingSparkball, sparkball, path, shortest, shortestCache)
	} else if path < *shortest {
		*shortest = path
		shortestCache[key] = path
	}
}
