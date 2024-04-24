package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Print("Введите целове число: ")
	fmt.Scan(&input)

	for _, char := range input {

		if char >= '0' && char <= '9' {

			digit := int(char - '0')

			squared := digit * digit
			fmt.Print(squared)
		} else {
			fmt.Print("Вам символ не является цифрой! Попробуйте еще раз!")
		}
	}
	fmt.Println()
}
