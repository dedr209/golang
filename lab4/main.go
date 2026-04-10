package main

import "github.com/andlabs/ui"

func main() {
	ui.Main(func() {
		initGUI()
	})
}

func initGUI() {
	window := ui.NewWindow("Lab 4", 800, 600, false)
	window.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	window.Show()
}
