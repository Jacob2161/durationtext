# Changelog

This module follows semantic versioning. Before v1, a minor version may
change the API; each such change says what callers must do.

## v0.3.0 (2026-11-28)

- Requires Go 1.27, the newest Go release; Go 1.26 is no longer
  supported.
- CI runs staticcheck and govulncheck at the versions pinned in
  `tools/go.mod`, a module of its own that importers never see.

## v0.2.0 (2026-11-27)

This is a breaking change and ships as v0.2.0.

- The flag type `Value` is now `Duration`, and `NewValue` is now
  `NewDuration`. Its methods are unchanged. Replace `durationtext.Value`
  with `durationtext.Duration` and `durationtext.NewValue` with
  `durationtext.NewDuration`.
- CI also checks formatting, `go mod tidy -diff`, and `govulncheck`.

## v0.1.0 (2026-11-27)

First release.

- `Parse` reads a Go duration (90m, 1h30m) or a whole number of days (30d).
- `Format` writes positive whole days as Nd and drops trailing zero units,
  so `Parse(Format(d)) == d` for every duration.
- `Value` adapts both to the flag and pflag packages without depending on
  either.
