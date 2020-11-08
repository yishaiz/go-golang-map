package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "40ff05",
		"white": "ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, " is ", hex)
	}
}

// colors := map[string]string{
// 	"red":   "#ff0000",
// 	"green": "40ff05",
// 	"blue":  "0000ff",
// }

// colors["white"] = "#ffffff"

// // colors2 := make(map[string]string)

// fmt.Println(colors)
// delete(colors, "white")
// fmt.Println(colors)

// colors2 := map[int]string{
// 	11: "#ff0000",
// 	12: "40ff05",
// 	13: "0000ff",
// }

// colors2[10] = "#f0f0f0"

// fmt.Println(colors2)
// delete(colors2, 10)
// fmt.Println(colors2)
