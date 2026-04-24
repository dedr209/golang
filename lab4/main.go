package main

import (
	"fmt"
	"strconv"

	"github.com/andlabs/ui"
	"lab4.com/calculator"
	// "github.com/andlabs/ui/winmanifest" // Uncomment if compiling on Windows for native styling
)

func createTourTab() ui.Control {
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
	calcButton := ui.NewButton("Розрахувати Вартість Туру")
	resultLabel := ui.NewLabel("0.00 $")
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
		var total float64
		var err error
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
	return mainBox
}
func createWindowTab() ui.Control {
	widthEntry := ui.NewEntry()
	heightEntry := ui.NewEntry()
	materialCombo := ui.NewCombobox()
	materialCombo.Append("Дерево")
	materialCombo.Append("Метал")
	materialCombo.Append("Металопластик")
	materialCombo.SetSelected(0)
	chamberCombo := ui.NewCombobox()
	chamberCombo.Append("1-камерний")
	chamberCombo.Append("2-камерний")
	chamberCombo.SetSelected(0)
	sillCheck := ui.NewCheckbox("Підвіконня (+350 грн)")
	variantCombo := ui.NewCombobox()
	variantCombo.Append("Варіант 1 (Go)")
	variantCombo.Append("Варіант 2 (C через cgo)")
	variantCombo.SetSelected(0)
	calcButton := ui.NewButton("Розрахувати Вартість Вікна")
	resultLabel := ui.NewLabel("0.00 грн")
	leftBox := ui.NewVerticalBox()
	leftBox.SetPadded(true)
	leftBox.Append(ui.NewLabel("Ширина (см):"), false)
	leftBox.Append(widthEntry, false)
	leftBox.Append(ui.NewLabel("Висота (см):"), false)
	leftBox.Append(heightEntry, false)
	leftBox.Append(ui.NewLabel("Спосіб обчислення:"), false)
	leftBox.Append(variantCombo, false)
	rightBox := ui.NewVerticalBox()
	rightBox.SetPadded(true)
	rightBox.Append(ui.NewLabel("Матеріал:"), false)
	rightBox.Append(materialCombo, false)
	rightBox.Append(ui.NewLabel("Кількість камер:"), false)
	rightBox.Append(chamberCombo, false)
	rightBox.Append(sillCheck, false)
	topBox := ui.NewHorizontalBox()
	topBox.SetPadded(true)
	topBox.Append(leftBox, true)
	topBox.Append(rightBox, true)
	mainBox := ui.NewVerticalBox()
	mainBox.SetPadded(true)
	mainBox.Append(topBox, false)
	mainBox.Append(calcButton, false)
	mainBox.Append(resultLabel, false)
	calcButton.OnClicked(func(*ui.Button) {
		width, errWidth := strconv.ParseFloat(widthEntry.Text(), 64)
		height, errHeight := strconv.ParseFloat(heightEntry.Text(), 64)
		if errWidth != nil || errHeight != nil {
			resultLabel.SetText("Помилка: введіть числа для ширини і висоти")
			return
		}
		material := materialCombo.Selected()
		chambers := chamberCombo.Selected() + 1
		hasSill := sillCheck.Checked()
		var total float64
		var err error
		if variantCombo.Selected() == 0 {
			total, err = calculator.CalculateWindowPriceGo(width, height, material, chambers, hasSill)
		} else {
			total, err = calculator.CalculateWindowPriceC(width, height, material, chambers, hasSill)
		}
		if err != nil {
			resultLabel.SetText("Помилка: " + err.Error())
			return
		}
		resultLabel.SetText(fmt.Sprintf("%.2f грн", total))
	})
	return mainBox
}
func initGUI() {
	window := ui.NewWindow("Лабораторна робота 4", 600, 400, false)
	window.SetMargined(true)
	tab := ui.NewTab()
	windowTab := createWindowTab()
	tourTab := createTourTab()
	tab.Append("Завдання 1: Калькулятор Вікна", windowTab)
	tab.SetMargined(0, true)
	tab.Append("Завдання 2: Калькулятор Туру", tourTab)
	tab.SetMargined(1, true)
	window.SetChild(tab)
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
