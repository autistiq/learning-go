package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64
type sensor func() kelvin

func calibrate() sensor {
	return func() kelvin {
		return kelvin(rand.Intn(151) + 150)
	}
}

func main() {
	sensor := calibrate()
	fmt.Println(sensor())
	fmt.Println(sensor())
	fmt.Println(sensor())
}
