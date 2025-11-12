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
			memo := make(map[int]int, 2048)
			start := word[len(word)-1]

			p := pathsMemo(start, rules, memo, []byte(word), longestWord, skipWords)
			total += p
		}
	}

	fmt.Println(total)
}

func pathsMemo(curr byte, rules map[byte][]byte, memo map[int]int, word []byte, longestWord int, skipWords map[string]struct{}) int {
	if len(word) <= longestWord {
		delete(skipWords, string(word))
	}

	if len(word) == 11 {
		return 1
	}

	count := 0
	if letters, ok := rules[curr]; ok {
		for _, letter := range letters {
			key := int(letter)*100 + len(word) + 1
			if p, mOk := memo[key]; mOk {
				count += p
			} else {
				newWord := append(word, letter)
				p = pathsMemo(letter, rules, memo, newWord, longestWord, skipWords)
				memo[key] = p
				count += p
			}
		}
	}

	if len(word) >= 7 {
		count++
	}

	return count
}
