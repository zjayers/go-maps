package main

import "fmt"

func main() {

	// Map Declarations

	// Method #1 - Literal Syntax
	// map[key type] value type
	colorMap := map[string]string{
		"red": "#FF0000",
		"blue": "#00FF00",
		"yellow": "#0000FF",
	}

	// Method #2 - Var syntax
	var colorMap2 map[string]string


	// Method #3 - Make Syntax
	colorMap3 := make(map[string]string)
	colorMap["white"] = "#FFFFFF"

	// Deletion of items inside of map
	delete(colorMap, "white")

	fmt.Println(colorMap, colorMap2, colorMap3)
	printMap(colorMap)
}

func printMap(c map[string]string)  {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
