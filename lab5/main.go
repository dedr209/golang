package main

import "fmt"

func main() {
	worker := NewWorker("Vitalii B.O.", 2020, 9, "Dnipro Office")
	company := NewCompany("EPAM", "Go Developer", 50000)

	fmt.Printf("Worker: %s, start date: %02d/%d, workplace: %s\n",
		worker.Name, worker.Month, worker.Year, worker.WorkPlace)
	fmt.Printf("Company: %s, position: %s, salary: %d\n",
		company.Name, company.Position, company.Salary)
}
