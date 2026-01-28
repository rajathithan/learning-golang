package main

import (
	"fmt"
)



// interface definition
// after "type" use uppercase if you want to make it public
type describable interface {
	describe() string
	// one more method can be added here
	
}

// product struct implementing the interface
type product struct {
	Name  string
	Price float64
}

// method to satisfy the Describable interface
func (p product) describe() string {
	return fmt.Sprintf("Product Name: %s, Price: %.2f", p.Name, p.Price)
}

// service struct implementing the interface
type service struct {
	Title       string
	HourlyRate  float64
}

// method to satisfy the Describable interface
func (s service) describe() string {
	return fmt.Sprintf("Service Title: %s, Hourly Rate: %.2f", s.Title, s.HourlyRate)
}

// service struct can have additional methods
func (s service) TotalCost(hours int) float64 {
	return s.HourlyRate * float64(hours)
}

// function that takes the Describable interface as a parameter
func PrintDescription(d describable) {
	fmt.Println(d.describe())
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
	fmt.Printf("Total cost for %d hours of service: %.2f\n", hours, service.TotalCost(hours))
}