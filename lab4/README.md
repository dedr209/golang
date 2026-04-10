# Lab 4 (GUI Calculator)

This project now implements the required two calculation variants:

- Variant 1: pure Go functions.
- Variant 2: C implementation called from Go via cgo.

## Files

- `main.go` - GUI and variant selector.
- `calculator/variant1.go` - Go calculation logic.
- `calculator/variant2.go` + `calculator/price.c` - cgo wrapper and C logic.

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
