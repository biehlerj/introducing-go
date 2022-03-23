package main

import "fmt"

func main() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)
	fToC()
	ftToM()
}

func fToC() {
	fmt.Print("Enter the temperature in Fahrenheit (without the F): ")
	var input float64
	fmt.Scanf("%f", &input)

	celsius := (input - 32) * 5 / 9

	fmt.Println(celsius)
}

func ftToM() {
	fmt.Print("Enter the length in feet (without the ft): ")
	var input float64
	fmt.Scanf("%f", &input)

	meters := input * 0.3048

	fmt.Println(meters)
}
