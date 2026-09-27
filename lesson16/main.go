package main

import "fmt"

func kelvinToCelsius(k float64) float64 {
	return k - 273.15
}

func celsiusToFahrenheit(c float64) float64 {
	f := (c * 9.0 / 5.0) + 32.0
	return f
}

func kelvinToFahrenheit(k float64) float64 {
	f := (k-273.15)*1.8 + 32
	return f
}

func main() {
	kelvin := 233.0
	celsius := kelvinToCelsius(kelvin)
	fmt.Printf("%.2f °K is %.2f °C\n", kelvin, celsius)
	fahrenheit := celsiusToFahrenheit(celsius)
	fmt.Printf("%.2f °C is %.2f °F\n", celsius, fahrenheit)
	kelvin = 0.0
	fahrenheit = kelvinToFahrenheit(kelvin)
	fmt.Printf("%.2f °K is %.2f °F\n", kelvin, fahrenheit)
}
