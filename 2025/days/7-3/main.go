package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	rows, err := util.ReadStrings("7-3", false, "\n")
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

	longestWord := 0
	skipWords := make(map[string]struct{}, len(words))
	for _, word := range words {
		skipWords[word] = struct{}{}
		if len(word) > longestWord {
			longestWord = len(word)
		}
	}

	for _, word := range words {
		if _, ok := skipWords[word]; !ok {
			continue
		}

		valid := true
		for j := 0; j < len(word)-1; j++ {
			if !slices.Contains(rules[word[j]], word[j+1]) {
				valid = false
				break
			}
		}

		if valid {
			start := word[len(word)-1]

			paths(start, rules, []byte(word), longestWord, skipWords, &total)
		}
	}

	fmt.Println(total)
}

func paths(curr byte, rules map[byte][]byte, word []byte, longestWord int, skipWords map[string]struct{}, total *int) {
	if len(word) <= longestWord {
		delete(skipWords, string(word))
	}

	if len(word) == 11 {
		*total++
		return
	}

	if letters, ok := rules[curr]; ok {
		for _, letter := range letters {
			newWord := append(word, letter)
			paths(letter, rules, newWord, longestWord, skipWords, total)
		}
	}

	if len(word) >= 7 {
		*total++
	}
}
