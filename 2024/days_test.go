package test

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var expectations = map[string]string{
	"1-1":  "1368",
	"1-2":  "5637",
	"1-3":  "27685",
	"2-1":  "25",
	"2-2":  "5286",
	"2-3":  "11701",
	"3-1":  "125",
	"3-2":  "2754",
	"3-3":  "10369",
	"4-1":  "89",
	"4-2":  "990436",
	"4-3":  "120796719",
	"5-1":  "3222",
	"5-2":  "11247667667709",
	"5-3":  "7700100110001000",
	"6-1":  "RRMKKFQTFGBM@",
	"6-2":  "RWQDLPWFVH@",
	"6-3":  "RSBGXNMFTVDR@",
	"7-1":  "IDBKJFGEC",
	"7-2":  "CFDGEJBKA",
	"7-3":  "6201",
	"8-1":  "10547645",
	"8-2":  "146862375",
	"8-3":  "41082",
	"9-1":  "13317",
	"9-2":  "4945",
	"9-3":  "153680",
	"10-1": "QSCHGTMVLRJXWDNP",
	"10-2": "193169",
	"10-3": "214674",
}

func TestDays2024(t *testing.T) {
	for day, expect := range expectations {
		t.Run(day, func(t *testing.T) {
			t.Parallel()
			runCmd := exec.Command("go", "run", fmt.Sprintf("days/%s/main.go", day))
			output, err := runCmd.CombinedOutput()
			if err != nil {
				fmt.Println(string(output))
			}

			assert.NoError(t, err)
			assert.Equal(t, expect, strings.TrimRight(string(output), "\n"), fmt.Sprintf("Day %s", day))
		})
	}
}
