package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"Red":   "#ff0000",
		"Blue":  "#4bf745",
		"White": "#ffffff",
	}

	printMap(colors)
}

func printMap(m map[string]string) {
	for k, v := range m {
		fmt.Println("Hex code for ", k, " is:", v)
	}
}
