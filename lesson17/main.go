package main

import "fmt"

type (
	kelvin     float64
	fahrenheit float64
	celsius    float64
)

// Кельвины в цельсии +
func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

// Кельвины в фаренгейт +
func (k kelvin) fahrenheit() fahrenheit {
	return fahrenheit((k-273.15)*1.8 + 32)
}

// Фаренгейт в цельсии +
func (f fahrenheit) celsius() celsius {
	return celsius((f - 32) / 1.8)
}

// Фаренгейт в кельвины +
func (f fahrenheit) kelvin() kelvin {
	return kelvin((f-32)/1.8 + 273.15)
}

// Цельсии в кельвины +
func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

// Цельсии в фаренгейт
func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit(c*1.8 + 32)
}

func main() {
	var k kelvin = 100.0
	var c celsius = 100.0
	var f fahrenheit = 100.0
	fmt.Printf("%.2f °K is %.2f °C\n", k, k.celsius())
	fmt.Printf("%.2f °K is %.2f °F\n", k, k.fahrenheit())
	fmt.Printf("%.2f °F is %.2f °C\n", f, f.celsius())
	fmt.Printf("%.2f °F is %.2f °K\n", f, f.kelvin())
	fmt.Printf("%.2f °C is %.2f °K\n", c, c.kelvin())
	fmt.Printf("%.2f °C is %.2f °F\n", c, c.fahrenheit())
}
