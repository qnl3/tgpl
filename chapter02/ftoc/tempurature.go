package main

import "fmt"

func tempurature() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%gºF fahrenheit is equal to %gºC celsius.\n", freezingF, fToC(freezingF))
	fmt.Printf("%gºF fahrenheit is equal to %gºC celsius.\n", boilingF, fToC(boilingF))
}
