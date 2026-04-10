package main

import "github.com/andlabs/ui"

func main() {
	ui.Main(func() {
		initGUI()
	})
}

func initGUI() {
	window := ui.NewWindow("Lab 4", 800, 600, false)

	// Prepare basic controls now; layout wiring will be added later.
	entry := ui.NewEntry()
	combobox := ui.NewCombobox()
	combobox.Append("Option 1")
	combobox.Append("Option 2")
	checkbox := ui.NewCheckbox("Enable option")
	button := ui.NewButton("Submit")
	label := ui.NewLabel("Sample label")

	_ = entry
	_ = combobox
	_ = checkbox
	_ = button
	_ = label

	window.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	window.Show()
}
