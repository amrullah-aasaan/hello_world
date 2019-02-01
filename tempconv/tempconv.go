package tempconv

import "fmt"

// Celsius Sample comment
type Celsius float64

// Fahrenheit Sample comment
type Fahrenheit float64

const (
	// AbsoluteZeroC sample comment
	AbsoluteZeroC Celsius = -273.15
	// Freezing sample comment
	Freezing Celsius = 0
	// BoilingC sample comment
	BoilingC Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}
