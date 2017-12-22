package main

import (
	"fmt"
)

const fahrenheit = 212.0

func main() {
	boilingPoint()
}

func boilingPoint() {
	var f = fahrenheit

	var c = (fahrenheit - 32) * 5 / 8

	fmt.Printf("boiling point = %gºF or %gºC\n", f, c)
}
