package main

import "fmt"

func main() {
	var chartype int8 = 'R'

	fmt.Printf("Code '%c' - %d\n", chartype, chartype)

	// Завдання.
	// 1. Вивести українську літеру 'Ї'
	// var ukrCharTest int8 = 'Ї' //overflowing
	var ukrChar rune = 'Ї'
	fmt.Printf("Code '%c' - %d\n", ukrChar, ukrChar)
	// 2. Пояснити призначення типу "rune" rune  →  до 1 114 111
}
