package main

import "fmt"

func main() {
	isActive := "1"
	var b bool
	switch isActive {
	case "true", "1", "yes":
		b = true
	case "false", "no", "0":
		b = false
	default:
		fmt.Println("Ошибка ввода")
		return
	}
	fmt.Println(b)
}
