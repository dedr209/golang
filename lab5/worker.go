package main

import "time"

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

func (w Worker) GetName() string {
	return w.Name
}

func (w *Worker) SetName(name string) {
	w.Name = name
}

func (w Worker) GetYear() int {
	return w.Year
}

func (w *Worker) SetYear(year int) {
	w.Year = year
}

func (w Worker) GetMonth() int {
	return w.Month
}

func (w *Worker) SetMonth(month int) {
	w.Month = month
}

func (w Worker) GetWorkPlace() string {
	return w.WorkPlace
}

func (w *Worker) SetWorkPlace(workPlace string) {
	w.WorkPlace = workPlace
}

func (w Worker) GetWorkerPosition(company Company) string {
	return company.GetPosition()
}

func (w Worker) GetWorkExperience() int {
	now := time.Now()
	startMonths := w.Year*12 + w.Month
	currentMonths := now.Year()*12 + int(now.Month())
	experience := currentMonths - startMonths
	if experience < 0 {
		return 0
	}
	return experience
}

func (w Worker) GetTotalMoney(company Company) int {
	return w.GetWorkExperience() * company.GetSalary()
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

func (c Company) GetName() string {
	return c.Name
}

func (c *Company) SetName(name string) {
	c.Name = name
}

func (c Company) GetPosition() string {
	return c.Position
}

func (c *Company) SetPosition(position string) {
	c.Position = position
}

func (c Company) GetSalary() int {
	return c.Salary
}

func (c *Company) SetSalary(salary int) {
	c.Salary = salary
}

func (c Company) GetAnnualSalary() int {
	return c.Salary * 12
}

func (c Company) HasPosition(position string) bool {
	return c.Position == position
}

func (c *Company) RaiseSalary(amount int) {
	if amount > 0 {
		c.Salary += amount
	}
}
