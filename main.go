package main

import "fmt"

func main () {
	// map declaration
	// colors := map[string]string{
	// 	"red": "#ff0000",
	// 	"green": "#4bf745",
	// }

	// Others ways of kdeclaring maps
	// var colors map[string]string // initializes with zero value
	
	//another way declaring maps
	colors := make(map[string]string)

	colors["white"] = " #ffffff"

	delete(colors, "white")

	fmt.Println(colors)
}