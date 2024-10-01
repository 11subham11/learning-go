package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Conversion constants
const (
	ftToCm    = 30.48
	inToCm    = 2.54
	kgToPound = 2.20462
)

// Function to convert height from ft and in to cm
func convertHeight(feet, inches int) float64 {
	// Convert feet to cm and inches to cm, then add them
	return (float64(feet) * ftToCm) + (float64(inches) * inToCm)
}

// Function to convert weight from kg to pounds and pounds to kg
func convertWeight(value float64, toUnit string) float64 {
	if toUnit == "pounds" {
		return value * kgToPound
	} else if toUnit == "kg" {
		return value / kgToPound
	}
	return 0
}

// Parse height input in the format "5'11"
func parseHeightInput(input string) (int, int, error) {
	// Split input into feet and inches parts
	parts := strings.Split(input, "'")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid height format")
	}

	feet, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid feet value")
	}

	inchesStr := strings.Trim(parts[1], `"`) // Remove double quote if it's present
	inches, err := strconv.Atoi(inchesStr)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid inches value")
	}

	return feet, inches, nil
}

func main() {
	var heightInput string
	var weightValue float64
	var weightUnit string

	// Height conversion
	fmt.Print("Enter height in feet and inches (e.g., 5'11\"): ")
	fmt.Scanln(&heightInput)

	feet, inches, err := parseHeightInput(heightInput)
	if err != nil {
		fmt.Println("Error parsing height:", err)
		return
	}
	heightInCm := convertHeight(feet, inches)
	fmt.Printf("Height in cm: %.2f cm\n", heightInCm)

	// Weight conversion
	fmt.Print("Enter weight value (in kg or pounds): ")
	fmt.Scanf("%f %s", &weightValue, &weightUnit)

	if weightUnit == "kg" {
		fmt.Printf("Weight in pounds: %.2f pounds\n", convertWeight(weightValue, "pounds"))
	} else if weightUnit == "pounds" {
		fmt.Printf("Weight in kg: %.2f kg\n", convertWeight(weightValue, "kg"))
	} else {
		fmt.Println("Invalid weight unit. Use 'kg' or 'pounds'.")
	}
}
