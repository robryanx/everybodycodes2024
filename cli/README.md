# Everybody Codes 2024 CLI

This command-line helper automates the tedious steps for the Everybody Codes 2024 event:

- fetch encrypted inputs, descriptions, and sample notes for one or more days,
- decrypt them using the published `key1`, `key2`, `key3` values,
- persist the decrypted strings under `inputs/<day>-<part>.txt` and `samples/<day>-<part>.txt`,
- optionally run your local Go solutions and submit their answers back to the site.

The binary lives in `cli/main.go` and is intended to be run via `go run ./cli [...]`.

## Prerequisites

1. Export your Everybody Codes login cookie once in your shell (or add it to `~/.bash_profile` as shown in this repo):

   ```bash
   export EVERYBODY_CODES_COOKIE="your-cookie-value"
   ```

2. The day/part solver programs are expected under `./days/<day>-<part>/main.go`. Each solver must print the final answer on stdout.

3. Use a writable Go build cache if your environment is sandboxed (e.g. `GOCACHE=$(pwd)/.gocache`).

## Usage

```bash
go run ./cli [flags]
```

### Flags

| Flag          | Description                                                                                 | Default |
| ------------- | ------------------------------------------------------------------------------------------- | ------- |
| `-day`        | Day number or inclusive range (e.g. `3`, `5-8`). Values must fall between 1 and 25.        | `1`     |
| `-part`       | Part number or inclusive range (e.g. `2`, `1-3`). Values must be within 1–3.               | `1-3`   |
| `-submit`     | When present, runs `go run ./days/<day>-<part>` to capture the answer, then POSTs it back. | `false` |

### Examples

Fetch day 9 inputs/descriptions for all three parts:

```bash
go run ./cli -day 9
```

Fetch days 7 through 10, only part 2:

```bash
go run ./cli -day 7-10 -part 2
```

Fetch day 12, part 3, and submit the locally-computed answer:

```bash
go run ./cli -day 12 -part 3 -submit
```

If a submission fails with HTTP 409, the CLI prints `answer already submitted` using the response body for context.

## Output Files

- `inputs/<day>-<part>.txt` – decrypted puzzle input.
- `samples/<day>-<part>.txt` – decrypted example notes (only written when found).

Existing files are overwritten so reruns keep the latest data.

## Error Handling

- Missing cookie or network failures cause the program to exit with a descriptive error.
- Decryption failures (bad key/ciphertext) surface immediately.
- Solver errors bubble up with combined stdout/stderr for easy debugging.
- Submission responses are parsed and reported; non-2xx responses include the API message when available.
