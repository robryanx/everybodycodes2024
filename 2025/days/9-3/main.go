package main

import (
	"fmt"
	"strings"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	rows, err := util.ReadStrings("9-3", false, "\n")
	if err != nil {
		panic(err)
	}

	var chains [][]byte
	for row := range rows {
		chains = append(chains, []byte(row[strings.Index(row, ":")+1:]))
	}

	ducks := make(map[int][]int, 1024)
	for i := 0; i < len(chains); i++ {
	check:
		for j := 0; j < len(chains); j++ {
			for k := j + 1; k < len(chains); k++ {
				if j != i && k != i {
					if candidateCheck(chains[i], chains[j], chains[k]) {
						ducks[i] = append(ducks[i], j, k)
						ducks[j] = append(ducks[j], i)
						ducks[k] = append(ducks[k], i)
						break check
					}
				}
			}
		}
	}

	best := 0
	ducksMapped := make(map[int]struct{}, 1024)
	for scale, d := range ducks {
		if _, ok := ducksMapped[scale]; ok {
			continue
		}

		count := recTree(scale, d, ducks, ducksMapped)
		if count > best {
			best = count
		}
	}

	fmt.Println(best)
}

func recTree(scale int, duck []int, ducks map[int][]int, tree map[int]struct{}) int {
	if _, ok := tree[scale]; ok {
		return 0
	}

	tree[scale] = struct{}{}
	count := scale + 1
	for _, c := range duck {
		count += recTree(c, ducks[c], ducks, tree)
	}
	return count
}

func candidateCheck(child, parentA, parentB []byte) bool {
	for i := 0; i < len(child); i++ {
		if parentA[i] == parentB[i] && child[i] != parentA[i] ||
			(child[i] != parentA[i] && child[i] != parentB[i]) {
			return false
		}
	}

	return true
}
