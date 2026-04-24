package main

type Worker struct {
	Name      string
	Year      int
	Month     int
	WorkPlace string
}

func NewWorker(name string, year int, month int, workPlace string) Worker {
	return Worker{
		Name:      name,
		Year:      year,
		Month:     month,
		WorkPlace: workPlace,
	}
}

type Company struct {
	Name     string
	Position string
	Salary   int
}

func NewCompany(name string, position string, salary int) Company {
	return Company{
		Name:     name,
		Position: position,
		Salary:   salary,
	}
}
