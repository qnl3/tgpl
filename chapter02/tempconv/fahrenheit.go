package tempconv

import "fmt"

const (
	// AbsoluteZeroFahrenheit is the theoretic tempurature at with an "ideal gas" reaches its minimal enthropic and entropic limit.
	AbsoluteZeroFahrenheit Fahrenheit = -459.67
	// FreezingFahrenheit is the tempurature on the celsius scale at which water freezes at normal atmospheric presures.
	FreezingFahrenheit Fahrenheit = 32.0
	// BoilingFahrenheit is the tempurature on the celsius scale at which water boils at normal atmospheric presures.
	BoilingFahrenheit Fahrenheit = 212.0
)

// Fahrenheit is a tepurature scale invented by Daniel Gabriel Fahrenheit in the year 1724.
type Fahrenheit float64

// FToC convers values in fahrenheit to values in celsius.
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// FToK convers values in fahrenheit to values in Kelvin.
func FToK(f Fahrenheit) Kelvin {
	return Kelvin((f - 32) * 5 / 9)
}

// ToC is a method of the Fahrenhiet type that returns its tempurature in Celsius
func (f Fahrenheit) ToC() Celsius {
	return FToC(f)
}

// ToK is a method of the Fahrenhiet type that returns its tempurature in Kelvin.
func (f Fahrenheit) ToK() Kelvin {
	return FToK(f)
}

// String returns a tempurature as a string uising standad markings and symbols
func (f Fahrenheit) String() string {
	return fmt.Sprintf("%gºF", f)
}
