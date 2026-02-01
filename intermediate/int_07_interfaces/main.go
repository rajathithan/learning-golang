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

// describe method to satisfy the details interface
func (p product) describe() string {
	return fmt.Sprintf("Product Name: %s, Price: %.2f", p.Name, p.Price)
}


// monthlyprice method to satisfy the details interface
func (p product) monthlyprice() float64 {
	return p.Price 
}



// service struct implementing the interface
type service struct {
	Title       string
	HourlyRate  float64
}

// describe method to satisfy the details interface
func (s service) describe() string {
	return fmt.Sprintf("Service Title: %s, Hourly Rate: %.2f", s.Title, s.HourlyRate)
}

// monthlyprice method to satisfy the details interface
func (s service) monthlyprice() float64 { 
	return s.HourlyRate * 160 // assuming 160 working hours in a month
}


// service struct can have additional methods
func (s service) hourlycost(hours int) float64 {
	return s.HourlyRate * float64(hours)
}

// manufacturer struct implementing the interface
type manufacturer struct {
	CompanyName string
	Country     string
}

// describe method to satisfy the details interface
func (m manufacturer) describe() string {
	return fmt.Sprintf("Manufacturer: %s, Country: %s", m.CompanyName, m.Country)
}

// monthlyprice method to satisfy the details interface
//func (m manufacturer) monthlyprice() float64 {
//	return 0 // Manufacturer might not have a monthly price, so returning 0
//}

// function that takes the details interface as a parameter
func PrintDescription(d details) {
	fmt.Println(d.describe())
	fmt.Printf("Monthly Price: %.2f\n", d.monthlyprice())
}



func main() {
	// Creating instances of Product and Service
	product := product{Name: "Laptop", Price: 1200.00}
	service := service{Title: "Consulting", HourlyRate: 150.00}
	//manufacturer := manufacturer{CompanyName: "TechCorp", Country: "USA"}

	// Using the PrintDescription function
	PrintDescription(product)
	PrintDescription(service)

	// cannot use manufacturer (variable of struct type manufacturer) as details value 
	// in argument to PrintDescription: manufacturer does not implement details 
	// (missing method monthlyprice)

	// PrintDescription(manufacturer)


	// Using the additional method of Service struct
	hours := 10
	fmt.Printf("Total cost for %d hours of service: %.2f\n", hours, service.hourlycost(hours))
}