package durationtext

import (
	"testing"
	"time"
)

// TestParseAndFormat checks days and Go durations both ways, and refuses
// days that are not a whole, representable number.
func TestParseAndFormat(t *testing.T) {
	for text, want := range map[string]time.Duration{
		"30d":   30 * Day,
		" 2d ":  2 * Day,
		"90m":   90 * time.Minute,
		"1h30m": 90 * time.Minute,
		"0":     0,
	} {
		parsed, err := Parse(text)
		if err != nil || parsed != want {
			t.Errorf("Parse(%q) = %s, %v; want %s", text, parsed, err, want)
		}
	}
	for _, text := range []string{"1.5d", "-1d", "d", "99999999999d", "soon"} {
		if _, err := Parse(text); err == nil {
			t.Errorf("Parse(%q) passed", text)
		}
	}
	for duration, want := range map[time.Duration]string{
		0:                "0",
		48 * time.Hour:   "2d",
		10 * time.Minute: "10m",
		2 * time.Hour:    "2h",
		90 * time.Second: "1m30s",
	} {
		if got := Format(duration); got != want {
			t.Errorf("Format(%s) = %q, want %q", duration, got, want)
		}
	}
}

// TestFormatNegative writes negative whole days as a Go duration, which Parse
// reads back, instead of as days, which it refuses.
func TestFormatNegative(t *testing.T) {
	if got := Format(-2 * Day); got != "-48h" {
		t.Fatalf("Format(-2d) = %q, want -48h", got)
	}
}

// FuzzRoundTrip checks that Format writes every duration so Parse reads it
// back unchanged, and that whatever Parse accepts formats to text it accepts.
func FuzzRoundTrip(f *testing.F) {
	for _, seed := range []string{"30d", "90m", "1h30m", "0", "-2h", "1.5h", "2562047h"} {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, text string) {
		parsed, err := Parse(text)
		if err != nil {
			return
		}
		formatted := Format(parsed)
		reparsed, err := Parse(formatted)
		if err != nil || reparsed != parsed {
			t.Fatalf("Parse(%q) = %s, formatted %q, reparsed %s, %v", text, parsed, formatted, reparsed, err)
		}
	})
}

// FuzzFormat checks the round trip from any duration.
func FuzzFormat(f *testing.F) {
	for _, seed := range []int64{0, 1, -1, int64(Day), -int64(Day), 1<<63 - 1, -1 << 63} {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, nanoseconds int64) {
		duration := time.Duration(nanoseconds)
		reparsed, err := Parse(Format(duration))
		if err != nil || reparsed != duration {
			t.Fatalf("Format(%d) = %q, reparsed %s, %v", nanoseconds, Format(duration), reparsed, err)
		}
	})
}
