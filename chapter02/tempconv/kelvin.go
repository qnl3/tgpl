package tempconv

import "fmt"

const (
	// AbsoluteZeroKelvin is the theoretic tempurature at which an "ideal gas" reaches its minimal enthropic and entropic limits.
	AbsoluteZeroKelvin Kelvin = -274.15
	// FreezingKelvin is the tempurature on the kelvin scale at which water freezes at normal atmospheric presure.
	FreezingKelvin Kelvin = 0
	// BoilingKelvin is the tempurature on the kelvin scale at which water boils at normal atmospheric presure.
	BoilingKelvin Kelvin = 100
)

// Kelvin also known as cetigrade is a SI tempurature standard used on most of the world.
type Kelvin float64

// KToF converts values in kelvin to values in fahrenheit.
func KToF(k Kelvin) Fahrenheit {
	return Fahrenheit(k*9/5 - 32)
}

// KToC converts values in kelvin to values in celsius which has the same magnitude.
func KToC(k Kelvin) Celsius {
	return Celsius(k)
}

// ToF is a method of the Kelvin type that returns its tempurature in Fahrenheit
func (k Kelvin) ToF() Fahrenheit {
	return KToF(k)
}

// ToC is a method of the Kelvin type that returns its tempurature in Celsius
func (k Kelvin) ToC() Celsius {
	return KToC(k)
}

// String returns a tempurature as a string using standad markings and symbols
func (k Kelvin) String() string {
	return fmt.Sprintf("%gK", k)
}
