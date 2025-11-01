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
		rec(stamps, sparkball, sparkball, path, &shortest, shortestCache)
		total += shortest + base
	}

	fmt.Println(total)
}

func rec(stamps []int, startingSparkball, sparkball int, path int, shortest *int, shortestCache map[int]int) {
	if len(stamps) == 0 || sparkball < 0 {
		return
	}

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
		rec(stamps, startingSparkball, sparkball-stamps[0], path+1, shortest, shortestCache)
		rec(stamps[1:], startingSparkball, sparkball, path, shortest, shortestCache)
	} else if path < *shortest {
		*shortest = path
		shortestCache[key] = path
	}
}
