package main

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
	"strings"
)

func main() {
	// Define a template
	// tmpl := template.New("example")

	// tmpl, err := tmpl.Parse("Hello, {{.Name}}! You are {{.Age}} years old.\n")
	// if err != nil {
	// 	panic(err)
	// }

	// Define a template with MustParse which panics on error, no need to declare an error variable
	tmpl := template.Must(template.New("example").Parse("Hello, {{.Name}}! You are {{.Age}} years old.\n"))

	// Create a data structure to hold the values
	data := map[string]interface{}{
		"Name": "Alice",
		"Age":  30,
	}

	// Execute the template with the data
	err := tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

	// Now let's create multiple templates and execute them based on user input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the employee name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	// Define multiple templates in a map
	templates := map[string]string{
		// name: tmpl
		"greeting": "Hello, {{.Name}}! Welcome to the company.\n",
		"salary":   "Dear {{.Name}}, your salary is ${{.Salary}} per month.\n",
		"designation": "Hi {{.Name}}, your designation is {{.Designation}}.\n",
	}

	// Parse the templates and store them in a map
	parsedTemplates := make(map[string]*template.Template)
	for name, tmpl := range templates {
		parsedTemplates[name] = template.Must(template.New(name).Parse(tmpl))
	}

	for {
		// show Menu
		fmt.Println("Select a template to execute:")
		fmt.Println("1. Greeting")
		fmt.Println("2. Salary")
		fmt.Println("3. Designation")
		fmt.Println("4. Exit")

		// Read user choice and trim whitespace
		// Note that the newline is in single quotes and not double quotes, 
		// this is to identify the end of the input
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		// data with key as string and value as interface, 
		// this allows us to store any type of value in the map
		var data map[string]interface{}

		// tmpl is a pointer to a template.Template struct, 
		// this will hold the template that we want to execute based on user choice
		var tmpl *template.Template

		// switch case to determine which template to execute based on user choice
		switch choice {	
		case "1":
			data = map[string]interface{}{
				"Name": name,
			}
			tmpl = parsedTemplates["greeting"]
		case "2":
			fmt.Print("Enter the salary: ")
			salaryStr, _ := reader.ReadString('\n')
			salaryStr = strings.TrimSpace(salaryStr)
			data = map[string]interface{}{
				"Name":   name,
				"Salary": salaryStr,
			}
			tmpl = parsedTemplates["salary"]
		case "3":
			fmt.Print("Enter the designation: ")
			designation, _ := reader.ReadString('\n')
			designation = strings.TrimSpace(designation)
			data = map[string]interface{}{
				"Name":        name,
				"Designation": designation,
			}
			tmpl = parsedTemplates["designation"]
		case "4":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
			continue
		}

		err := tmpl.Execute(os.Stdout, data)
		if err != nil {
			panic(err)
		}
	}

}
