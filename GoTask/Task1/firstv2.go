package main

import (
	"fmt"
)

func reverseNumber(num int) int {
	newInt := 0
	for num > 0 {
		remainder := num % 10
		newInt = newInt*10 + remainder
		num /= 10
	}
	return newInt
}

func main() {
	var number int
	fmt.Print("Введите трехзначное число: ")
	fmt.Scan(&number)

	if number >= 100 && number <= 999 && number%10 != 0 {
		fmt.Println("Исходное число:", number)
		reversed := reverseNumber(number)
		fmt.Println("Перевернутое число:", reversed)
	} else {
		fmt.Println("Ошибка! Попробуйте ввести трехзначное число, не оканчивающееся на ноль!")
	}
}
