package main

import (
	"fmt"

	"github.com/robryanx/everybodycodes2024/util"
)

func main() {
	rows, err := util.ReadInts("9-2", false, "\n")
	if err != nil {
		panic(err)
	}

	stamps := []int{30, 25, 24, 20, 16, 15, 10, 5, 3, 1}

	total := 0
	shortestCache := make(map[int]int, 1028)
	for sparkball := range rows {
		shortest := sparkball
		path := 0
		base := sparkball / stamps[0] * 9 / 10
		sparkball -= base * stamps[0]
		rec(stamps, 0, sparkball, sparkball, path, &shortest, shortestCache)
		total += shortest + base
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
