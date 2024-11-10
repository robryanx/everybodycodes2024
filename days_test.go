package test

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var expectations = map[string]string{
	"1-1": "1368",
	"1-2": "5637",
	"1-3": "27685",
	"2-1": "25",
	"2-2": "5286",
	"2-3": "11701",
}

func TestDays(t *testing.T) {
	for day, expect := range expectations {
		expect := expect
		day := day
		t.Run(day, func(t *testing.T) {
			t.Parallel()
			cmd := exec.Command(fmt.Sprintf("bin/%s", day))
			out, err := cmd.CombinedOutput()

			assert.NoError(t, err)
			assert.Equal(t, expect, strings.TrimRight(string(out), "\n"), fmt.Sprintf("Day %s", day))
		})
	}
}
