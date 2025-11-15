package main

import (
	"fmt"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	rows, err := util.ReadStrings("9-1", false, "\n")
	if err != nil {
		panic(err)
	}

	var chains [][]byte
	for row := range rows {
		chains = append(chains, []byte(row[2:]))
	}

	var child []byte
	var parentA []byte
	var parentB []byte
	if candidateCheck(chains[0], chains[1], chains[2]) {
		child = chains[0]
		parentA = chains[1]
		parentB = chains[2]
	} else if candidateCheck(chains[1], chains[0], chains[2]) {
		child = chains[1]
		parentA = chains[0]
		parentB = chains[2]
	} else if candidateCheck(chains[2], chains[0], chains[1]) {
		child = chains[2]
		parentA = chains[0]
		parentB = chains[1]
	} else {
		panic("bad")
	}

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

	fmt.Println(countA * countB)
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
