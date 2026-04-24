package main

import "fmt"

func main() {
	workers := ReadWorkersArray()
	currentCompany = readCompany()
	PrintWorkers(workers)
	maxSalary, minSalary := GetWorkersInfo(workers)
	fmt.Printf("Max salary: %d\n", maxSalary)
	fmt.Printf("Min salary: %d\n", minSalary)
}
