package main

import (
	"bufio"
	"fmt"
	"os"

	emp "example.com/employee"
)

func main() {
	var empls []emp.Employee

	empls = append(empls, emp.Employee{Name: "Mike", Salary: 10000},
		emp.Employee{Name: "John", Salary: 9000})

	for _, e := range empls {
		fmt.Printf("%s - %f\n", e.Name, e.Salary)
	}

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
}
