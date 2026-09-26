package main

import (
	"fmt"
	"strings"
)

func main() {
	plainText := "text I want to be ciphered"
	noSpaces := strings.Replace(plainText, " ", "", -1)
	processedText := strings.ToUpper(noSpaces)
	keyword := "GOLANG"

	for i := 0; i < len(processedText); i++ {
		c := processedText[i]
		k := keyword[i%len(keyword)]
		cipher := c - 'A'
		key := k - 'A'
		diff := cipher + key
		wrapped := (diff+26)%26 + 'A'
		fmt.Printf("%c", wrapped)
	}
}
