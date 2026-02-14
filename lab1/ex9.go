package main

import "fmt"

func main() {
	var first, second bool
	var third bool = true
	fourth := !third
	fifth := true

	fmt.Println("first  = ", first)       // false
	fmt.Println("second = ", second)      // false
	fmt.Println("third  = ", third)       // true
	fmt.Println("fourth = ", fourth)      // false
	fmt.Println("fifth  = ", fifth, "\n") // true

	fmt.Println("!true  = ", !true)        // false
	fmt.Println("!false = ", !false, "\n") // true

	fmt.Println("true && true   = ", true && true)         // true
	fmt.Println("true && false  = ", true && false)        // false
	fmt.Println("false && false = ", false && false, "\n") // false

	fmt.Println("true || true   = ", true || true)         // true
	fmt.Println("true || false  = ", true || false)        // true
	fmt.Println("false || false = ", false || false, "\n") // false

	fmt.Println("2 < 3  = ", 2 < 3)        // true
	fmt.Println("2 > 3  = ", 2 > 3)        // false
	fmt.Println("3 < 3  = ", 3 < 3)        // false
	fmt.Println("3 <= 3 = ", 3 <= 3)       // true
	fmt.Println("3 > 3  = ", 3 > 3)        // false
	fmt.Println("3 >= 3 = ", 3 >= 3)       // true
	fmt.Println("2 == 3 = ", 2 == 3)       // false
	fmt.Println("3 == 3 = ", 3 == 3)       // true
	fmt.Println("2 != 3 = ", 2 != 3)       // true
	fmt.Println("3 != 3 = ", 3 != 3, "\n") // false

	// Завдання.
	// 1. Пояснити результати операцій
	// Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
}
