package main

import (
	"fmt"
	"strings"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	rows, err := util.ReadStrings("9-2", false, "\n")
	if err != nil {
		panic(err)
	}

	var chains [][]byte
	for row := range rows {
		chains = append(chains, []byte(row[strings.Index(row, ":")+1:]))
	}

	total := 0
	parents := make([]bool, len(chains))
	for i := 0; i < len(chains); i++ {
		if parents[i] {
			continue
		}

		for j := 0; j < len(chains); j++ {
			for k := j + 1; k < len(chains); k++ {
				if j != i && k != i {
					if candidateCheck(chains[i], chains[j], chains[k]) {
						parents[j] = true
						parents[k] = true

						total += similarity(chains[i], chains[j], chains[k])
					}
				}
			}
		}
	}

	fmt.Println(total)
}

func similarity(child, parentA, parentB []byte) int {
	countA := 0
	countB := 0

	for i := 0; i < len(child); i++ {
		if child[i] == parentA[i] {
			countA++
		}
		if child[i] == parentB[i] {
			countB++
		}
	}

	return countA * countB
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
