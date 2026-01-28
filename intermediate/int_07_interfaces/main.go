package main

import (
	"fmt"
)



// interface definition
// after "type" use uppercase if you want to make it public
type details interface {
	describe() string
	monthlyprice() float64

}

// product struct implementing the interface
type product struct {
	Name  string
	Price float64
}

// method to satisfy the details interface
func (p product) describe() string {
	return fmt.Sprintf("Product Name: %s, Price: %.2f", p.Name, p.Price)
}

func (p product) monthlyprice() float64 {
	return p.Price * 31
}

// service struct implementing the interface
type service struct {
	Title       string
	HourlyRate  float64
}

// method to satisfy the details interface
func (s service) describe() string {
	return fmt.Sprintf("Service Title: %s, Hourly Rate: %.2f", s.Title, s.HourlyRate)
}

func (s service) monthlyprice() float64 { 
	return s.HourlyRate * 160 // assuming 160 working hours in a month
}


// service struct can have additional methods
func (s service) hourlycost(hours int) float64 {
	return s.HourlyRate * float64(hours)
}


// function that takes the details interface as a parameter
func PrintDescription(d details) {
	fmt.Println(d.describe())
	fmt.Printf("Monthly Price: %.2f\n", d.monthlyprice())
}

func main() {
	// Creating instances of Product and Service
	product := product{Name: "Laptop", Price: 1200.00}
	service := service{Title: "Consulting", HourlyRate: 150.00}

	// Using the PrintDescription function
	PrintDescription(product)
	PrintDescription(service)

	// Using the additional method of Service struct
	hours := 10
	fmt.Printf("Total cost for %d hours of service: %.2f\n", hours, service.hourlycost(hours))
}