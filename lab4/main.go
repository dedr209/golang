package main

import (
	"fmt"
	"strconv"

	"github.com/andlabs/ui"
	"lab4.com/calculator"
	// "github.com/andlabs/ui/winmanifest" // Uncomment if compiling on Windows for native styling
)

func initGUI() {
	// 1. Create the main window
	window := ui.NewWindow("Калькулятор вартості туру", 540, 360, false)
	window.SetMargined(true)

	// 2. Create UI components (inputs)
	daysEntry := ui.NewEntry()
	vouchersEntry := ui.NewEntry()

	countryCombo := ui.NewCombobox()
	countryCombo.Append("Болгарія")
	countryCombo.Append("Німеччина")
	countryCombo.Append("Польща")
	countryCombo.SetSelected(0)

	seasonCombo := ui.NewCombobox()
	seasonCombo.Append("Літо")
	seasonCombo.Append("Зима")
	seasonCombo.SetSelected(0)

	roomCombo := ui.NewCombobox()
	roomCombo.Append("Стандарт")
	roomCombo.Append("Люкс (+20%)")
	roomCombo.SetSelected(0)

	variantCombo := ui.NewCombobox()
	variantCombo.Append("Варіант 1 (Go)")
	variantCombo.Append("Варіант 2 (C через cgo)")
	variantCombo.SetSelected(0)

	guideCheck := ui.NewCheckbox("Індивідуальний гід ($50/день)")

	calcButton := ui.NewButton("Розрахувати")
	resultLabel := ui.NewLabel("0.00 $")

	// 3. Layout management
	leftBox := ui.NewVerticalBox()
	leftBox.SetPadded(true)
	leftBox.Append(ui.NewLabel("Параметри поїздки"), false)
	leftBox.Append(ui.NewLabel("Кількість днів"), false)
	leftBox.Append(daysEntry, false)
	leftBox.Append(ui.NewLabel("Кількість путівок"), false)
	leftBox.Append(vouchersEntry, false)
	leftBox.Append(ui.NewLabel("Країна"), false)
	leftBox.Append(countryCombo, false)

	rightBox := ui.NewVerticalBox()
	rightBox.SetPadded(true)
	rightBox.Append(ui.NewLabel("Сезон"), false)
	rightBox.Append(seasonCombo, false)
	rightBox.Append(ui.NewLabel("Тип номера"), false)
	rightBox.Append(roomCombo, false)
	rightBox.Append(ui.NewLabel("Спосіб обчислення"), false)
	rightBox.Append(variantCombo, false)
	rightBox.Append(guideCheck, false)

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
		days, errDays := strconv.Atoi(daysEntry.Text())
		vouchers, errVouchers := strconv.Atoi(vouchersEntry.Text())

		if errDays != nil || errVouchers != nil {
			resultLabel.SetText("Помилка: введіть цілі числа для днів і путівок")
			return
		}

		country := countryCombo.Selected()
		season := seasonCombo.Selected()
		roomType := roomCombo.Selected()
		withGuide := guideCheck.Checked()

		var (
			total float64
			err   error
		)

		if variantCombo.Selected() == 0 {
			total, err = calculator.CalculateTourPriceGo(days, country, season, vouchers, roomType, withGuide)
		} else {
			total, err = calculator.CalculateTourPriceC(days, country, season, vouchers, roomType, withGuide)
		}

		if err != nil {
			resultLabel.SetText("Помилка: " + err.Error())
			return
		}

		resultLabel.SetText(fmt.Sprintf("%.2f $", total))
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
