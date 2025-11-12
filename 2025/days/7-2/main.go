package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	rows, err := util.ReadStrings("7-2", false, "\n")
	if err != nil {
		panic(err)
	}

	var words []string
	rules := map[byte][]byte{}

	count := 0
	for row := range rows {
		if count == 0 {
			words = strings.Split(row, ",")
		} else if count > 1 {
			parts := strings.Split(row, " > ")
			var letters []byte
			for letter := range strings.SplitSeq(parts[1], ",") {
				letters = append(letters, letter[0])
			}
			rules[parts[0][0]] = letters
		}

		count++
	}

	total := 0
	for i, word := range words {
		valid := true
		for j := 0; j < len(word)-1; j++ {
			if !slices.Contains(rules[word[j]], word[j+1]) {
				valid = false
				break
			}
		}

		if valid {
			total += i + 1
		}
	}

	fmt.Println(total)
}
