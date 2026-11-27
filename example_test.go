package durationtext_test

import (
	"flag"
	"fmt"
	"time"

	"github.com/jacob2161/durationtext"
)

func ExampleParse() {
	retention, err := durationtext.Parse("30d")
	if err != nil {
		panic(err)
	}
	fmt.Println(retention, durationtext.Format(retention))
	// Output: 720h0m0s 30d
}

func ExampleNewValue() {
	flags := flag.NewFlagSet("backup", flag.ContinueOnError)
	var retention time.Duration
	flags.Var(durationtext.NewValue(&retention, 7*durationtext.Day), "retention", "how long to keep backups")
	if err := flags.Parse([]string{"-retention", "90m"}); err != nil {
		panic(err)
	}
	fmt.Println(retention, flags.Lookup("retention").DefValue)
	// Output: 1h30m0s 7d
}
