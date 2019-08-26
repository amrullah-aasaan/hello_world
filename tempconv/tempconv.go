package tempconv

import (
	"flag"
	"fmt"
)

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

type celsiusFlag struct {
	Celsius
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64

	fmt.Scanf(s, "%f%s", &value, &unit)

	fmt.Println(s)

	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

// CelsiusFlag some comment
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
