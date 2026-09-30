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
	for i := 0; i < 10; i++ {
		k := fakeSensor()
		c := kelvinToCelsius(k)
		fmt.Printf("%vK is %v°C\n", k, c)
		if k < min {
			min = k
		} else if max < k {
			max = k
		}
	}
	fmt.Printf("%vK is max %vK is min", max, min)
}
