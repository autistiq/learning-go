package main

import (
	"fmt"
	"math/rand"
)

func main() {
	guessedNumber := 71
GameLoop:
	for {
		var n = rand.Intn(100) + 1
		switch {
		case n < guessedNumber:
			fmt.Printf("%v is less than guessed number\n", n)
		case n > guessedNumber:
			fmt.Printf("%v is more than guessed number\n", n)
		case n == guessedNumber:
			fmt.Printf("%v is guessed number!\n", n)
			break GameLoop
		}

	}
}
