package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/robryanx/everybodycodes2024/util"
)

func main() {
	strs, err := util.ReadStrings("2-1", false, "\n")
	if err != nil {
		panic(err)
	}

	allStrs := slices.Collect(strs)

	wordsStr := allStrs[0][strings.Index(allStrs[0], ":")+1:]

	count := 0
	for _, word := range strings.Split(wordsStr, ",") {
		count += strings.Count(allStrs[2], word)
	}

	fmt.Println(count)
}
