package firstTasks

import "fmt"

type IEmployee interface {
	GetSalary() int
	GetEmpInfo() string
}

type Employee struct {
	DailyWage  int
	Type       string
	DaysWorked int
	Name       string
}

type FreeLancer struct {
	HourlyPay   int
	HoursWorked int
	Name        string
	Type        string
}

func (e *Employee) GetSalary() int {
	return e.DailyWage * e.DaysWorked * 12
}

func (e *Employee) GetEmpInfo() string {
	return fmt.Sprintf("Name: %v - %v", e.Name, e.Type)
}

func (f *FreeLancer) GetSalary() int {
	return f.HourlyPay * f.HoursWorked
}

func (f *FreeLancer) GetEmpInfo() string {
	return fmt.Sprintf("Name: %v - %v", f.Name, f.Type)
}

func PrintSalary(emp IEmployee) {
	fmt.Printf("Salary is %v\n", emp.GetSalary())
}

func PrintEmpInfo(emp IEmployee) {
	fmt.Printf("Employee %v\n", emp.GetEmpInfo())
}
