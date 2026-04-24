package main

import "fmt"

func main() {
	worker := Worker{
		Name:      "Vitalii B.O.",
		Year:      2020,
		Month:     9,
		WorkPlace: "Dnipro Office",
	}

	fmt.Printf("Worker: %s, start date: %02d/%d, workplace: %s\n",
		worker.Name, worker.Month, worker.Year, worker.WorkPlace)
}
