package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Синоніми цілих типів\n")

	fmt.Println("byte    - int8")
	fmt.Println("rune    - int32")
	fmt.Println("int     - int32 або int64, в залежності від платформи")
	fmt.Println("uint    - uint32 або uint64, в залежності від платформи")

	// Завдання.
	// 1. Визначити розрядність платформи
	fmt.Println("розрядність платформи", strconv.IntSize)
}
