package main

import (
	"fmt"
)

type Person struct {
	Name   string
	Age    int
	Emails []string
}

// anonymous fields struct
type Phone struct {
	home   string
	mobile string
	work   string
}

// Embedding a struct
type Employee struct {
	Emp Person   // Embedded struct
	Position string
	Salary   float64
	Phone    // Embedded struct with anonymous fields
}

// value receiver method
func (e Employee) DisplayInfo() {
	fmt.Printf("Name: %s, Position: %s, Salary: %.2f\n", e.Emp.Name, e.Position, e.Salary)
}

// pointer receiver method
func (e *Employee) GiveRaise(percent float64) {
	e.Salary += e.Salary * percent / 100
}

func main() {	// Creating an instance of Person struct
	p := Employee{
		Emp: Person{
			Name:   "Alice",
			Age:    30,
			Emails: []string{"alice@example.com", "alice.work@example.com"},
		},
		Position: "Software Engineer",
		Salary:   75000.50,
		Phone: Phone{
			home:   "555-1234",
			mobile: "555-5678",
			work:   "555-8765",
		},
	}

	p2 := Employee{}	
	p2.Emp.Name = "Bob"
	p2.Emp.Age = 28
	p2.Position = "Data Analyst"
	p2.Salary = 65000.00
	p2.Emp.Emails = append(p2.Emp.Emails, "bob@example.com")
	p2.home = "555-4321"
	p2.mobile = "555-8765"
	p2.work = "555-5678"

	fmt.Println("Employee-1", p.Emp.Name)
	fmt.Println("Employee-2", p2.Emp.Name)

	p.DisplayInfo()
	p.GiveRaise(10)
	fmt.Println("After 10% raise:")
	p.DisplayInfo()

}