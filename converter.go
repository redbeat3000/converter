package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run converter.go <value> <from_unit> <to_unit>")
		fmt.Println("Example: go run converter.go 100 cm m")
		return
	}

	value, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Invalid number:", os.Args[1])
		return
	}

	fromUnit := strings.ToLower(os.Args[2])
	toUnit := strings.ToLower(os.Args[3])

	result, err := convert(value, fromUnit, toUnit)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("%.4f %s = %.4f %s\n", value, fromUnit, result, toUnit)
	}
}

func convert(value float64, from, to string) (float64, error) {
	conversions := map[string]float64{
		"cm:m": 0.01,
		"m:cm": 100,
		"km:m": 1000,
		"m:km": 0.001,
		"kg:g": 1000,
		"g:kg": 0.001,
		"c:f":  0, // special case
		"f:c":  0, // special case
		"c:k":  0, // special case
		"k:c":  0, // special case
	}

	key := from + ":" + to

	switch key {
	case "c:f":
		return (value*9/5 + 32), nil
	case "f:c":
		return ((value - 32) * 5 / 9), nil
	case "c:k":
		return value + 273.15, nil
	case "k:c":
		return value - 273.15, nil
	}

	factor, ok := conversions[key]
	if !ok {
		return 0, fmt.Errorf("conversion from %s to %s not supported", from, to)
	}

	return value * factor, nil
}
