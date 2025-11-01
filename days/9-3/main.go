package main

import (
	"fmt"

	"github.com/robryanx/everybodycodes2024/util"
)

func main() {
	rows, err := util.ReadInts("9-3", false, "\n")
	if err != nil {
		panic(err)
	}

	stamps := []int{101, 100, 75, 74, 50, 49, 38, 37, 30, 25, 24, 20, 16, 15, 10, 5, 3, 1}

	shortestCache := make(map[int]int, 1024)

	// TODO: Make this less painfully slow

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
			iterations := 0
			base := i/stamps[0] - 12
			adj := i - base*stamps[0]
			rec(stamps, adj, adj, path, &iterations, &shortest, shortestCache)
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

func rec(stamps []int, startingSparkball, sparkball int, path int, iterations, shortest *int, shortestCache map[int]int) {
	if len(stamps) == 0 || sparkball < 0 {
		return
	}

	*iterations++
	key := ((startingSparkball - sparkball) * 1000) + stamps[0]

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
		rec(stamps, startingSparkball, sparkball-stamps[0], path+1, iterations, shortest, shortestCache)
		rec(stamps[1:], startingSparkball, sparkball, path, iterations, shortest, shortestCache)
	} else if path < *shortest {
		*shortest = path
		shortestCache[key] = path
	}
}
