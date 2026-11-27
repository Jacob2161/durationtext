// Package durationtext reads and writes durations as people type them in
// flags, settings, and configuration files.
//
// Parse takes a Go duration (90m, 1h30m) or a whole number of days (30d).
// Format writes positive whole days as Nd and anything else as a Go duration
// without trailing zero units (10m, not 10m0s), so Parse(Format(d)) == d for
// every duration. Duration adapts both to the flag and pflag packages.
package durationtext

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

// Day is 24 hours, the unit Nd counts.
const Day = 24 * time.Hour

// Parse reads a Go duration or a whole, non-negative number of days, ignoring
// surrounding space.
func Parse(text string) (time.Duration, error) {
	text = strings.TrimSpace(text)
	days, inDays := strings.CutSuffix(text, "d")
	if inDays {
		count, err := strconv.ParseInt(days, 10, 64)
		if err != nil || count < 0 || count > math.MaxInt64/int64(Day) {
			return 0, fmt.Errorf("%q is not a whole number of days", text)
		}
		return time.Duration(count) * Day, nil
	}
	return time.ParseDuration(text)
}

// Format writes a positive whole number of days as Nd, and any other duration
// as a Go duration without trailing zero minutes or seconds.
func Format(duration time.Duration) string {
	if duration == 0 {
		return "0"
	}
	if duration > 0 && duration%Day == 0 {
		return strconv.FormatInt(int64(duration/Day), 10) + "d"
	}
	text := duration.String()
	if strings.HasSuffix(text, "m0s") {
		text = strings.TrimSuffix(text, "0s")
	}
	if strings.HasSuffix(text, "h0m") {
		text = strings.TrimSuffix(text, "0m")
	}
	return text
}

// Duration is a command line flag over a time.Duration that Parse reads and
// Format writes. It satisfies flag.Value and pflag.Value:
//
//	flags.Var(
//		durationtext.NewDuration(&timeout, 10*time.Minute),
//		"timeout",
//		"how long to wait",
//	)
//
// A *time.Duration converts to a *Duration directly, for a flag whose
// default is already in place: (*durationtext.Duration)(&timeout).
type Duration time.Duration

// NewDuration sets *target to defaultDuration and returns a Duration that
// sets *target when the flag is parsed.
func NewDuration(target *time.Duration, defaultDuration time.Duration) *Duration {
	*target = defaultDuration
	return (*Duration)(target)
}

// String writes the duration as Format does.
func (d *Duration) String() string { return Format(time.Duration(*d)) }

// Set reads the duration as Parse does.
func (d *Duration) Set(text string) error {
	parsed, err := Parse(text)
	if err != nil {
		return err
	}
	*d = Duration(parsed)
	return nil
}

// Type names the flag's type in pflag's help output.
func (d *Duration) Type() string { return "duration" }
