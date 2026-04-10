package main

import (
	"fmt"
	"strconv"

	"github.com/andlabs/ui"
	// "github.com/andlabs/ui/winmanifest" // Uncomment if compiling on Windows for native styling
)

func initGUI() {
	// 1. Create the main window
	window := ui.NewWindow("Калькулятор склопакета", 500, 300, false)
	window.SetMargined(true)

	// 2. Create UI components (inputs)
	widthEntry := ui.NewEntry()
	heightEntry := ui.NewEntry()

	materialCombo := ui.NewCombobox()
	materialCombo.Append("Дерево")
	materialCombo.Append("Метал")
	materialCombo.Append("Металопластик")
	materialCombo.SetSelected(0)

	glassCombo := ui.NewCombobox()
	glassCombo.Append("Однокамерний")
	glassCombo.Append("Двокамерний")
	glassCombo.SetSelected(0)

	windowsillCheck := ui.NewCheckbox("Підвіконня")

	calcButton := ui.NewButton("Розрахувати")
	resultLabel := ui.NewLabel("0.00 грн")

	// 3. Layout management
	leftBox := ui.NewVerticalBox()
	leftBox.SetPadded(true)
	leftBox.Append(ui.NewLabel("Розмір вікна"), false)
	leftBox.Append(ui.NewLabel("Ширина, см"), false)
	leftBox.Append(widthEntry, false)
	leftBox.Append(ui.NewLabel("Висота, см"), false)
	leftBox.Append(heightEntry, false)
	leftBox.Append(ui.NewLabel("Матеріал"), false)
	leftBox.Append(materialCombo, false)

	rightBox := ui.NewVerticalBox()
	rightBox.SetPadded(true)
	rightBox.Append(ui.NewLabel("Склопакет"), false)
	rightBox.Append(glassCombo, false)
	rightBox.Append(windowsillCheck, false)

	topBox := ui.NewHorizontalBox()
	topBox.SetPadded(true)
	topBox.Append(leftBox, true)
	topBox.Append(rightBox, true)

	mainBox := ui.NewVerticalBox()
	mainBox.SetPadded(true)
	mainBox.Append(topBox, false)
	mainBox.Append(calcButton, false)
	mainBox.Append(resultLabel, false)

	window.SetChild(mainBox)

	// 4. Event handling
	calcButton.OnClicked(func(*ui.Button) {
		width, errW := strconv.ParseFloat(widthEntry.Text(), 64)
		height, errH := strconv.ParseFloat(heightEntry.Text(), 64)

		if errW != nil || errH != nil {
			resultLabel.SetText("Помилка: введіть коректні числа")
			return
		}

		area := width * height
		matIdx := materialCombo.Selected()
		glassIdx := glassCombo.Selected()

		var pricePerCm float64
		if matIdx == 0 && glassIdx == 0 {
			pricePerCm = 2.5
		}
		if matIdx == 0 && glassIdx == 1 {
			pricePerCm = 3.0
		}
		if matIdx == 1 && glassIdx == 0 {
			pricePerCm = 0.5
		}
		if matIdx == 1 && glassIdx == 1 {
			pricePerCm = 1.0
		}
		if matIdx == 2 && glassIdx == 0 {
			pricePerCm = 1.5
		}
		if matIdx == 2 && glassIdx == 1 {
			pricePerCm = 2.0
		}

		total := area * pricePerCm
		if windowsillCheck.Checked() {
			total += 350.0
		}

		resultLabel.SetText(fmt.Sprintf("%.2f грн", total))
	})

	// 5. Window closure event
	window.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})

	window.Show()
}

func main() {
	err := ui.Main(initGUI)
	if err != nil {
		panic(err)
	}
}
