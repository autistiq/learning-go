package main

import "fmt"

func main() {
	prices := []int{1500, 800, 3000, 450}
	total := 0

	for _, price := range prices {
		total += price
	}
	fmt.Println("Итоговая сумма товаров в иенах:", total)
}
