package main

import (
	"fmt"
)

func main() {
	var k int

	fmt.Println("Введите число, соответствующее условию 0<k<86399: ")
	fmt.Scan(&k)

	if k > 0 && k < 86399 {
		h := k / 3600
		m := (k % 3600) / 60
		fmt.Printf("It is %d hours %d minutes", h, m)

	} else {
		fmt.Print("Ваше число не соотвтствует поставленному условию. Попробуйте еще раз!")
	}
}
