# Lab 4 (GUI Tour Calculator)

This project computes the total cost of a holiday tour order and implements two required calculation variants:

- Variant 1: pure Go functions.
- Variant 2: C implementation called from Go via cgo.

## Files

- `main.go` - GUI form and variant selector.
- `calculator/variant1.go` - Go calculation logic.
- `calculator/variant2.go` + `calculator/price.c` - cgo wrapper and C logic.

## Pricing rules

- Bulgaria: summer `$100/day`, winter `$150/day`
- Germany: summer `$160/day`, winter `$200/day`
- Poland: summer `$120/day`, winter `$180/day`
- Individual guide: `$50/day` for the whole order
- Lux room: `+20%` markup

## Run

```bash
go mod tidy
go run .
```

## Test calculation logic

```bash
go test ./calculator
```

## Build requirements

`github.com/andlabs/ui` wraps a C library, so native toolchains and GUI development packages are required.

- Linux: install GTK3/WebKit2GTK development packages and a C compiler.
- Windows: install `mingw-w64` (required by the assignment) so cgo and `andlabs/ui` can compile.
