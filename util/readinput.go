package util

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

type Iterator = func(s []byte)

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

	err := read(filename(dayPart, isSample), delim, func(b []byte) {
		vals = append(vals, string(b))
	})
	if err != nil {
		return nil, err
	}

	return vals, nil
}

func read(file string, delim string, iterator Iterator) error {
	b, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	for _, row := range bytes.Split(b, []byte(delim)) {
		iterator(row)
	}

	return nil
}
