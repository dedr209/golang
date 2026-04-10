# Lab 4 (GUI Skeleton)

Minimal starter for the 4th lab assignment using `github.com/andlabs/ui`.

## What is included

- `package main`
- Required import for `ui`
- `main()` with `ui.Main(...)`
- `initGUI()` that creates an empty window
- `OnClosing` handler that quits the app

## Run

```bash
go mod tidy
go run .
```

If your Linux system is missing GUI build dependencies for `libui`, install GTK3/WebKit2GTK development packages first.

