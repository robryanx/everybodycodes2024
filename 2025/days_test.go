package test

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var expectations = map[string]string{
	"1-1": "Lorther",
	"1-2": "Fyrardith",
	"1-3": "Sorbel",
	"2-1": "[220380,826546]",
	"2-2": "1316",
	"2-3": "129213",
}

func TestDays2025(t *testing.T) {
	for day, expect := range expectations {
		day := day
		expect := expect
		t.Run(day, func(t *testing.T) {
			t.Parallel()
			runCmd := exec.Command("go", "run", ".")
			runCmd.Dir = filepath.Join("days", day)
			output, err := runCmd.CombinedOutput()
			if err != nil {
				fmt.Println(string(output))
			}

			assert.NoError(t, err)
			assert.Equal(t, expect, strings.TrimRight(string(output), "\n"), fmt.Sprintf("Day %s", day))
		})
	}
}
