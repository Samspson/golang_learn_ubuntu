package main

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9.0/5.0 + 32.0) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32.0) * 5.0 / 9.0) }

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

type celsuisFlag struct{ Celsius }

func (f *celsuisFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Scanf(s, "%f%s", &value, &unit)
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

func CelsuisFlag(name string, value Celsius, usage string) *Celsius {
	f := celsuisFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

func main() {
	var temp = CelsuisFlag("temp", 20.0, "the temperature")
	flag.Parse()
	fmt.Println(*temp)
}
