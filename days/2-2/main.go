package main

import (
	"fmt"
	"strings"

	"github.com/robryanx/everybodycodes2024/util"
)

func main() {
	strs, err := util.ReadStrings("2-2", false, "\n")
	if err != nil {
		panic(err)
	}

	wordsStr := strs[0][strings.Index(strs[0], ":")+1:]

	count := 0
	for _, str := range strs[2:] {
		allIndexes := make(map[int]struct{}, len(str))
		for _, word := range strings.Split(wordsStr, ",") {
			indexes := strIndexes(str, word)
			for k := range indexes {
				allIndexes[k] = struct{}{}
			}

			indexes = strIndexes(reverse(str), word)
			for k := range indexes {
				allIndexes[len(str)-k-1] = struct{}{}
			}
		}

		count += len(allIndexes)
	}

	fmt.Println(count)
}

func strIndexes(str, word string) map[int]struct{} {
	indexes := make(map[int]struct{}, len(str))

	for i := 0; i <= len(str)-len(word); i++ {
		if str[i] == word[0] {
			match := true
			for j := i + 1; j < i+len(word); j++ {
				if str[j] != word[j-i] {
					match = false
					break
				}
			}

			if match {
				for j := i; j < i+len(word); j++ {
					indexes[j] = struct{}{}
				}
			}
		}
	}

	return indexes
}

func reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
