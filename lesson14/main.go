package main

import "fmt"

func main() {
	cipheredText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"
	for i := 0; i < len(cipheredText); i++ {
		c := cipheredText[i]
		k := keyword[i%len(keyword)]
		decipher := c - 'A'
		key := k - 'A'
		diff := decipher - key
		wrapped := (diff+26)%26 + 'A'
		fmt.Printf("%c", wrapped)
	}
}
