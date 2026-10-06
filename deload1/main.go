package main

import (
	"fmt"
)

func main() {
	distance := 56000000
	days := 28
	hours := days * 24
	speed := distance / hours
	fmt.Printf("%v km/h", speed)
}
