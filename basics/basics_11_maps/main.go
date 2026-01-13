package main

import "fmt"

func main() {

	var personAgeMap = map[string]int{
		"Alice": 30,
		"Bob":   25,
		"Eve":   35,
	}

	fmt.Println("personAgeMap:", personAgeMap)

	// Accessing map values
	fmt.Println("Age of Alice:", personAgeMap["Alice"])

	// Adding a new key-value pair
	personAgeMap["Charlie"] = 28
	fmt.Println("After adding Charlie:", personAgeMap)

	// Modifying an existing value
	personAgeMap["Bob"] = 26
	fmt.Println("After modifying Bob's age:", personAgeMap)

	// Deleting a key-value pair
	delete(personAgeMap, "Eve")
	fmt.Println("After deleting Eve:", personAgeMap)

	// Checking if a key exists
	age, exists := personAgeMap["Eve"]
	if exists {
		fmt.Println("Eve's age is:", age)
	} else {
		fmt.Println("Eve not found in the map")
	}

	// Iterating over a map
	fmt.Println("Iterating over personAgeMap:")
	for name, age := range personAgeMap {
		fmt.Printf("%s: %d\n", name, age)
	}

	// Length of the map
	fmt.Println("Number of entries in personAgeMap:", len(personAgeMap))	

	// Creating an empty map and adding entries
	countryCapitalMap := make(map[string]string)
	countryCapitalMap["USA"] = "Washington, D.C."
	countryCapitalMap["France"] = "Paris"
	fmt.Println("countryCapitalMap:", countryCapitalMap)

	// Copying a map
	copiedMap := make(map[string]int)
	for k, v := range personAgeMap {
		copiedMap[k] = v
	}
	fmt.Println("copiedMap:", copiedMap)

	// Modifying copiedMap will not affect personAgeMap
	copiedMap["Alice"] = 31
	fmt.Println("After modifying copiedMap:")
	fmt.Println("personAgeMap:", personAgeMap)
	fmt.Println("copiedMap:", copiedMap)

	// Note: Maps in Go are reference types, so assigning one map to another
	// will create a reference, not a copy.
	referenceMap := personAgeMap
	referenceMap["Bob"] = 27
	fmt.Println("After modifying referenceMap:")
	fmt.Println("personAgeMap:", personAgeMap)
	fmt.Println("referenceMap:", referenceMap)

	// Nested maps
	nestedMap := map[string]map[string]string{
		"USA": {
			"Capital": "Washington, D.C.",
			"Currency": "USD",
		},
		"France": {
			"Capital": "Paris",
			"Currency": "Euro",
		},
	}
	fmt.Println("nestedMap:", nestedMap)
	fmt.Println("Capital of USA:", nestedMap["USA"]["Capital"])	


}
