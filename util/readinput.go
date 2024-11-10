package util

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

type Iterator = func(s string)

func filename(dayPart string, isSample bool) string {
	folder := "inputs"
	if isSample {
		folder = "samples"
	}

	_, fileName, _, _ := runtime.Caller(0)
	prefixPath := filepath.Dir(fileName)

	return fmt.Sprintf("%s/../%s/%s.txt", prefixPath, folder, dayPart)
}

func ReadBytes(dayPart string, isSample bool) ([]byte, error) {
	return os.ReadFile(filename(dayPart, isSample))
}

func ReadStrings(dayPart string, isSample bool, delim string) ([]string, error) {
	var vals []string

	err := read(filename(dayPart, isSample), delim, func(s string) {
		vals = append(vals, s)
	})
	if err != nil {
		return nil, err
	}

	return vals, nil
}

func ReadInts(dayPart string, isSample bool, delim string) ([]int, error) {
	var vals []int

	err := read(filename(dayPart, isSample), delim, func(s string) {
		i, _ := strconv.Atoi(s)
		vals = append(vals, i)
	})
	if err != nil {
		return nil, err
	}

	return vals, nil
}

func ReadFloats(dayPart string, isSample bool, delim string) ([]float64, error) {
	var vals []float64

	err := read(filename(dayPart, isSample), delim, func(s string) {
		i, _ := strconv.ParseFloat(s, 64)
		vals = append(vals, i)
	})
	if err != nil {
		return nil, err
	}

	return vals, nil
}

func read(file string, delim string, iterator Iterator) error {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	for _, row := range strings.Split(string(bytes), delim) {
		iterator(row)
	}

	return nil
}
