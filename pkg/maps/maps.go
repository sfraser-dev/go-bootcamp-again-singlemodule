package pkg_maps

import (
	"fmt"
)

func CreateMaps() {

	// maps are like dicts (python) or hashes (perl)

	// Go maps are unordered so no push, pop, shift, unshift

	// short declaration (declare and assign all at once)
	colors := map[string]string{ // keys and values are of type string
		"red":    "#ff0000",
		"green":  "#fb1234",
		"blue":   "#457773",
		"yellow": "#765663",
		"black":  "#000000",
	}

	// long var declaration
	var vehicles = map[string]string{
		"01": "car",
		"02": "motorbike",
		"03": "boat",
		"04": "aeroplane",
	}

	// make creates a map an then assigns it to a variable
	computer := make(map[string]string) // creates an empty map
	computer["mac"] = "macOS"
	computer["pc"] = "windows"
	computer["server"] = "linux"

	fmt.Printf("colors = %+v\n", colors)
	fmt.Printf("vehicles = %+v\n", vehicles)
	fmt.Printf("computer = %+v\n", computer)

	fmt.Println("...Deleting red and green from colors...")
	delete(colors, "red")
	delete(colors, "green")
	fmt.Printf("colors = %+v\n", colors)
	fmt.Println("...Adding green back into colors...")
	colors["green"] = "#fb1234"
	fmt.Printf("colors = %+v\n", colors)

	// using for loop with range keyword to iterate over key value pairs in a map
	fmt.Println("...Iterating over maps...")
	for k, v := range colors {
		fmt.Printf("k=%-9v v=%-9v\n", k, v)
	}
}
