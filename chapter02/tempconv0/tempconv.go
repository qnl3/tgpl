package tempconv

import "fmt"

// Celsius also known as cetigrade is a SI tempurature standard used on most of the world.
type Celsius float64

// Fahrenheit is a tepurature scale invented by Daniel Gabriel Fahrenheit in the year 1724.
type Fahrenheit float64

const (
	// AbsoluteZeroCelsius is the theoretic tempurature at with an "ideal gas" reaches its minimal enthropic and entropic limit.
	AbsoluteZeroCelsius Celsius = -274.15
	// FreezingCelsius is the tempurature on the celsius scale at which water freezes at normal atmospheric presures.
	FreezingCelsius Celsius = 0
	// BoilingCelsius is the tempurature on the celsius scale at which water boils at normal atmospheric presures.
	BoilingCelsius Celsius = 100
)

// CToF converts values in celsius to values in fahrenheit.
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 - 32)
}

// FToC convers values in fahrenheit to values in celsius.
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// ToC is a method of the Fahrenhiet type that returns its tempurature in Celsius
func (f Fahrenheit) ToC() Celsius {
	return FToC(f)
}

// String returns a tempurature as a string uising standad markings and symbols
func (f Fahrenheit) String() string {
	return fmt.Sprintf("%gºF", f)
}

// ToF is a method of the Celsius type that returns its tempuratur in Fahrenheit
func (c Celsius) ToF() Fahrenheit {
	return CToF(c)
}

// String returns a tempurature as a string uising standad markings and symbols
func (c Celsius) String() string {
	return fmt.Sprintf("%gºC", c)
}
