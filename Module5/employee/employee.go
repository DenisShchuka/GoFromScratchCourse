package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

func (e *Employee) UpdateName(name string) bool {
	if len(name) > 0 {
		e.Name = name
		return true
	}
	return false
}

func (e *Employee) String() string {
	return fmt.Sprintf("%v | %s - %v", e.ID, e.Name, e.Salary)
}

func main() {
	empl := &Employee{ID: 0, Name: "Mike", Salary: 10000}
	fmt.Println(empl)
}
