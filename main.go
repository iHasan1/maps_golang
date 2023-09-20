package main

import "fmt"

func main () {
	// map declaration
	colors := map[string]string{
		"red": "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap (c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is =", hex );
	}
}