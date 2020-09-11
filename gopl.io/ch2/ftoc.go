package main

import (
	"fmt"
)

func main() {
	const freezingF, bolingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToc(freezingF))
	fmt.Printf("%g°F = %g°C\n", bolingF, fToc(bolingF))
}

func fToc(f float64) float64 {
	return (f - 32) * 5 / 9
}
