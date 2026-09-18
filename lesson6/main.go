package main

import (
	"fmt"
	"math/rand"
)

func main() {
	balance := 0
	for balance < 2000 {
		switch coinType := rand.Intn(3); coinType {
		case 0:
			balance += 5
		case 1:
			balance += 10
		case 2:
			balance += 25
		}
		fmt.Println("Баланс в копилке - ", float64(balance)/100, "$!")
	}
}
