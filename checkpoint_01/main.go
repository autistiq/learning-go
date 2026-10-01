package main

import (
	"fmt"
	"math/rand"
)

type (
	kelvin  float64
	celsius float64
)

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}

func main() {
	var min kelvin
	var max kelvin
	var iteraitonAmount celsius
	var sum celsius

	min = fakeSensor()
	max = fakeSensor()
	iteraitonAmount = 1

	for i := 0; i < 9; i++ {
		k := fakeSensor()
		c := kelvinToCelsius(k)
		fmt.Printf("%v K is %.0f °C\n", k, c)
		if k < min {
			min = k
		}
		if max < k {
			max = k
		}
		sum += c
		iteraitonAmount++
	}
	fmt.Printf("%vK is max %vK is min\n", max, min)
	meanTemp := sum / iteraitonAmount
	fmt.Printf("%.0f°C is the average temperature", meanTemp)
}
