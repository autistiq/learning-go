package main

import "fmt"

func main() {
	message := "shalom"
	for i := 0; i < len(message); i++ {
		c := message[i]
		fmt.Printf("%c", c)
	}
}
