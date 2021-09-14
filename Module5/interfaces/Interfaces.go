package main

import "fmt"

type Person interface {
	UpdateName(name string) bool
	String() string
}

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

type Contractor struct {
	ID   int
	Name string
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

func (e *Contractor) UpdateName(name string) bool {
	if len(name) > 0 {
		e.Name = name
		return true
	}
	return false
}

func (e *Contractor) String() string {
	return fmt.Sprintf("%v | %s ", e.ID, e.Name)
}

func printer(persons []Person) {
	for _, v := range persons {
		fmt.Println(v)
	}
}

func main() {
	empl := &Employee{ID: 0, Name: "Mike", Salary: 10000}
	contr := &Contractor{ID: 0, Name: "John"}
	p := []Person{empl, contr}
	printer(p)
}
