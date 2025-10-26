package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"html"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type User struct {
	Seed int `json:"seed"`
}

func main() {
	loginKey := os.Getenv("EVERYBODY_CODES_COOKIE")
	if loginKey == "" {
		panic("EVERYBODY_CODES_COOKIE environment variable is not set")
	}

	dayFlag := flag.String("day", "1", "Quest day number or inclusive range (e.g. 9 or 1-9)")
	partFlag := flag.String("part", "1-3", "Quest part number or inclusive range (e.g. 2 or 1-3)")
	debugFlag := flag.Bool("debug", false, "Print puzzle input and description to stdout")
	submitFlag := flag.Bool("submit", false, "Run local solver and submit answers")
	flag.Parse()

	days, err := parseSelection(*dayFlag, 1, 25) // hard cap at 25 days for safety
	if err != nil {
		panic(err)
	}

	parts, err := parseSelection(*partFlag, 1, 3)
	if err != nil {
		panic(err)
	}

	userURL := "https://everybody.codes/api/user/me"

	body, err := fetchWithCookie(userURL, loginKey)
	if err != nil {
		panic(err)
	}

	var user User
	if err := json.Unmarshal(body, &user); err != nil {
		panic(err)
	}

	for _, day := range days {
		if err := processDay(day, user.Seed, loginKey, *submitFlag, *debugFlag, parts); err != nil {
			panic(err)
		}
	}
}

func processDay(day, seed int, loginKey string, submit bool, debug bool, parts []int) error {
	questURL := fmt.Sprintf("https://everybody.codes/api/event/2024/quest/%d", day)
	body, err := fetchWithCookie(questURL, loginKey)
	if err != nil {
		return fmt.Errorf("fetch quest %d: %w", day, err)
	}

	var keys map[string]string
	if err := json.Unmarshal(body, &keys); err != nil {
		return fmt.Errorf("decode quest %d keys: %w", day, err)
	}

	puzzleURL := fmt.Sprintf("https://everybody-codes.b-cdn.net/assets/2024/%d/input/%d.json", day, seed)
	body, err = fetchWithCookie(puzzleURL, loginKey)
	if err != nil {
		return fmt.Errorf("fetch puzzle %d: %w", day, err)
	}

	var puzzles map[string]string
	if err := json.Unmarshal(body, &puzzles); err != nil {
		return fmt.Errorf("decode puzzle %d: %w", day, err)
	}

	descriptionURL := fmt.Sprintf("https://everybody-codes.b-cdn.net/assets/2024/%d/description.json", day)
	body, err = fetchWithCookie(descriptionURL, loginKey)
	if err != nil {
		return fmt.Errorf("fetch description %d: %w", day, err)
	}

	var descriptions map[string]string
	if err := json.Unmarshal(body, &descriptions); err != nil {
		return fmt.Errorf("decode description %d: %w", day, err)
	}

	for _, part := range parts {
		key := keys[fmt.Sprintf("key%d", part)]
		if key == "" {
			continue
		}

		puzzleCipher := puzzles[strconv.Itoa(part)]
		if puzzleCipher == "" {
			return fmt.Errorf("missing puzzle data for day %d part %d", day, part)
		}

		input, err := decrypt(puzzleCipher, key)
		if err != nil {
			return err
		}
		if debug {
			fmt.Printf("Day %d part %d input: %s\n\n", day, part, input)
		}

		if err := writeTextFile("inputs", day, part, input); err != nil {
			return err
		}

		descriptionCipher := descriptions[strconv.Itoa(part)]
		if descriptionCipher == "" {
			continue
		}

		description, err := decrypt(descriptionCipher, key)
		if err != nil {
			return err
		}
		if debug {
			fmt.Printf("Day %d part %d description: %s\n\n", day, part, description)
		}

		if sample := extractSampleNote(description); sample != "" {
			if debug {
				fmt.Printf("Day %d part %d sample: %s\n\n", day, part, sample)
			}
			if err := writeTextFile("samples", day, part, sample); err != nil {
				return err
			}
		}

		if submit {
			answer, err := runSolution(day, part)
			if err != nil {
				return err
			}
			fmt.Printf("Day %d part %d answer: %s\n", day, part, answer)

			resp, err := submitAnswer(day, part, answer, loginKey)
			if err != nil {
				fmt.Printf("Day %d part %d submission failed: %v\n\n", day, part, err)
			} else {
				fmt.Printf("Day %d part %d submission: correct=%v lengthCorrect=%v firstCorrect=%v\n\n", day, part, resp.Correct, resp.LengthCorrect, resp.FirstCorrect)
			}
		}
	}

	return nil
}

func decrypt(input, key string) (string, error) {
	ciphertext, err := hex.DecodeString(input)
	if err != nil {
		return "", fmt.Errorf("decode hex: %w", err)
	}

	keyBytes := []byte(key)
	switch len(keyBytes) {
	case 16, 24, 32:
	default:
		return "", fmt.Errorf("invalid AES key length %d", len(keyBytes))
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		return "", fmt.Errorf("ciphertext not a multiple of block size")
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", fmt.Errorf("init cipher: %w", err)
	}

	iv := keyBytes[:aes.BlockSize]
	plaintext := make([]byte, len(ciphertext))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintext, ciphertext)

	plaintext, err = pkcs7Unpad(plaintext, aes.BlockSize)
	if err != nil {
		return "", fmt.Errorf("pkcs7 unpad: %w", err)
	}

	return string(plaintext), nil
}

func pkcs7Unpad(data []byte, blockSize int) ([]byte, error) {
	if len(data) == 0 || len(data)%blockSize != 0 {
		return nil, fmt.Errorf("invalid padded data length %d", len(data))
	}
	padLen := int(data[len(data)-1])
	if padLen == 0 || padLen > blockSize || padLen > len(data) {
		return nil, fmt.Errorf("invalid padding length %d", padLen)
	}
	for _, b := range data[len(data)-padLen:] {
		if int(b) != padLen {
			return nil, fmt.Errorf("invalid padding byte %x", b)
		}
	}
	return data[:len(data)-padLen], nil
}

func fetchWithCookie(url, key string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "everybodycodes-cli/0.1 (+github.com/robryanx/everybodycodes2024)")
	req.AddCookie(&http.Cookie{Name: "everybody-codes", Value: key, Path: "/"})

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status %d: %s", resp.StatusCode, string(body))
	}

	return body, nil
}

var sampleNoteRE = regexp.MustCompile(`(?is)<pre[^>]*class="[^"]*note[^"]*"[^>]*>(.*?)</pre>`)

func extractSampleNote(descriptionHTML string) string {
	if descriptionHTML == "" {
		return ""
	}
	matches := sampleNoteRE.FindStringSubmatch(descriptionHTML)
	if len(matches) < 2 {
		return ""
	}
	note := strings.TrimSpace(matches[1])
	note = html.UnescapeString(note)
	note = strings.ReplaceAll(note, "\r\n", "\n")
	return strings.TrimSpace(note)
}

func writeTextFile(dir string, day, part int, content string) error {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("ensure dir %s: %w", dir, err)
	}
	filename := fmt.Sprintf("%d-%d.txt", day, part)
	path := filepath.Join(dir, filename)
	return os.WriteFile(path, []byte(content), 0o644)
}

func parseSelection(value string, min, max int) ([]int, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return nil, fmt.Errorf("selection is empty")
	}

	if strings.Contains(value, "-") {
		parts := strings.SplitN(value, "-", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid range %q", value)
		}
		start, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			return nil, fmt.Errorf("invalid number %q: %w", parts[0], err)
		}
		end, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			return nil, fmt.Errorf("invalid number %q: %w", parts[1], err)
		}
		if start < min || end < min || end > max || end < start {
			return nil, fmt.Errorf("range %q outside supported bounds %d-%d", value, min, max)
		}
		days := make([]int, 0, end-start+1)
		for d := start; d <= end; d++ {
			days = append(days, d)
		}
		return days, nil
	}

	n, err := strconv.Atoi(value)
	if err != nil {
		return nil, fmt.Errorf("invalid number %q: %w", value, err)
	}
	if n < min || n > max {
		return nil, fmt.Errorf("value %d outside supported bounds %d-%d", n, min, max)
	}
	return []int{n}, nil
}

func runSolution(day, part int) (string, error) {
	target := fmt.Sprintf("./days/%d-%d", day, part)
	cmd := exec.Command("go", "run", target)
	cmd.Env = os.Environ()
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("run solver for day %d part %d: %w\n%s", day, part, err, string(output))
	}
	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	for i := len(lines) - 1; i >= 0; i-- {
		line := strings.TrimSpace(lines[i])
		if line != "" {
			return line, nil
		}
	}
	return "", fmt.Errorf("solver for day %d part %d produced no output", day, part)
}

type answerPayload struct {
	Answer string `json:"answer"`
}

type submitResponse struct {
	Correct       bool   `json:"correct"`
	LengthCorrect bool   `json:"lengthCorrect"`
	FirstCorrect  bool   `json:"firstCorrect"`
	Time          int64  `json:"time"`
	LocalTime     int64  `json:"localTime"`
	GlobalTime    int64  `json:"globalTime"`
	GlobalPlace   int64  `json:"globalPlace"`
	GlobalScore   int64  `json:"globalScore"`
	Message       string `json:"message,omitempty"`
}

func submitAnswer(day, part int, answer, loginKey string) (*submitResponse, error) {
	url := fmt.Sprintf("https://everybody.codes/api/event/2024/quest/%d/part/%d/answer", day, part)
	payloadBytes, err := json.Marshal(answerPayload{Answer: answer})
	if err != nil {
		return nil, fmt.Errorf("encode answer payload: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(payloadBytes))
	if err != nil {
		return nil, fmt.Errorf("build answer request: %w", err)
	}
	req.Header.Set("User-Agent", "everybodycodes-cli/0.1 (+github.com/robryanx/everybodycodes2024)")
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(&http.Cookie{Name: "everybody-codes", Value: loginKey, Path: "/"})

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("submit answer: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read answer response: %w", err)
	}

	var sr submitResponse
	if len(body) > 0 {
		if err := json.Unmarshal(body, &sr); err != nil {
			// keep raw message on decode failure
			sr.Message = strings.TrimSpace(string(body))
		}
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		if sr.Message == "" && resp.StatusCode == http.StatusConflict {
			sr.Message = "answer already submitted"
		}
		if sr.Message == "" {
			sr.Message = strings.TrimSpace(string(body))
		}
		if sr.Message == "" {
			sr.Message = fmt.Sprintf("submission rejected with status %d", resp.StatusCode)
		}
		return &sr, fmt.Errorf("submission rejected (%d): %s", resp.StatusCode, sr.Message)
	}

	return &sr, nil
}
