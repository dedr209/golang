package main

// Импорт нескольких пакетов
import (
	"fmt"
	"math"
)

func main() {
	var defaultFloat float32
	var defaultDouble float64 = 5.5

	fmt.Println("defaultfloat       = ", defaultFloat)
	fmt.Printf("defaultDouble (%T) = %f\n\n", defaultDouble, defaultDouble)

	fmt.Println("MAX float32        = ", math.MaxFloat32)
	fmt.Println("MIN float32        = ", math.SmallestNonzeroFloat32, "\n")

	fmt.Println("MAX float64        = ", math.MaxFloat64)
	fmt.Println("MIN float64        = ", math.SmallestNonzeroFloat64, "\n")

	// Завдання.
	// 1. Створіть змінні різних типів, використовуючи короткий запис та ініціалізацію за замовчуванням. Результат вивести в консоль
	name := "Віталій"
	age := 20
	height := 1.75
	isStudent := true

	var defaultInt int
	var defaultString string
	var defaultBool bool

	fmt.Printf("name=%s age=%d height=%.2f isStudent=%t\n", name, age, height, isStudent)
	fmt.Printf("defaultInt=%d defaultString=%q defaultBool=%t\n", defaultInt, defaultString, defaultBool)
}
