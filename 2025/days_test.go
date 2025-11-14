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
	"3-1": "2704",
	"3-2": "250",
	"3-3": "3672",
	"4-1": "11376",
	"4-2": "1523178807948",
	"4-3": "279543324177",
	"5-1": "6572742476",
	"5-2": "8823824244654",
	"5-3": "31293916",
	"6-1": "171",
	"6-2": "4249",
	"6-3": "1667589569",
	"7-1": "Nydravor",
	"7-2": "1825",
	"7-3": "4047997",
	"8-1": "55",
	"8-2": "2923890",
	"8-3": "2802",
}

func TestDays2025(t *testing.T) {
	for day, expect := range expectations {
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
