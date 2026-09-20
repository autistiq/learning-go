package main

import "fmt"

func main() {
	message := "L fdph, L vdz, L frqtxhuhg"
	runes := []rune(message)
	for i := 0; i < len(runes); i++ {
		c := runes[i]
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			c = c - 3
		}
		fmt.Printf("%c", c)
	}
}
