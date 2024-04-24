package main

import (
	"fmt"
	"strconv"
	"strings"
)

func reverseNumber(num int) int {

	numStr := strconv.Itoa(num)
	characters := strings.Split(numStr, "")

	reversedCharacters := make([]string, len(characters))
	for i, j := 0, len(characters)-1; i < len(characters); i, j = i+1, j-1 {
		reversedCharacters[i] = characters[j]
	}

	reversedStr := strings.Join(reversedCharacters, "")
	reversedNum, _ := strconv.Atoi(reversedStr)

	return reversedNum
}

func main() {
	var number int
	fmt.Print("Введите трехзначное число: ")
	fmt.Scan(&number)

	if number >= 100 && number <= 999 && number%10 != 0 {
		fmt.Println("Исходное число:", number)
	} else {
		fmt.Println("Ошибка! Попробуйте ввести трехзначное число, не оканчивающееся на ноль!")
		return
	}

	reversed := reverseNumber(number)
	fmt.Println("Перевернутое число:", reversed)
}
