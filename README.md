# durationtext

[![Go Reference](https://pkg.go.dev/badge/github.com/jacob2161/durationtext.svg)](https://pkg.go.dev/github.com/jacob2161/durationtext)
[![CI](https://github.com/jacob2161/durationtext/actions/workflows/ci.yml/badge.svg)](https://github.com/jacob2161/durationtext/actions/workflows/ci.yml)

Durations as people type them: Go durations (90m, 1h30m) and whole days (30d). Format writes positive whole days as Nd and drops trailing zero units (10m, not 10m0s), and Parse(Format(d)) == d for every duration.

```bash
go get github.com/jacob2161/durationtext
```

## Usage

```go
retention, err := durationtext.Parse("30d") // 720h0m0s
fmt.Println(durationtext.Format(retention))  // 30d

var timeout time.Duration
flags.Var(durationtext.NewValue(&timeout, 10*time.Minute), "timeout", "how long to wait")
```

- `Value` satisfies both `flag.Value` and `pflag.Value` (cobra), with no dependency on either.
- The round trip is fuzz-tested over every int64 duration.

Tested on the two newest Go releases. Licensed under the MIT license.
