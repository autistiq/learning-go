package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	era := "AD"
	daysInMonth := 31
	for i := 0; i < 10; i++ {
		year := rand.IntN(2100) + 1
		month := rand.IntN(12) + 1
		switch month {
		case 2:
			if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
				daysInMonth = 29
			} else {
				daysInMonth = 28
			}
		case 4, 6, 9, 11:
			daysInMonth = 30
		}
		day := rand.IntN(daysInMonth) + 1
		fmt.Println(era, year, month, day)
	}

}
