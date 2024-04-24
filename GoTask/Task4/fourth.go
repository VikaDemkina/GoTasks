package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var input string
	fmt.Print("Введите строку: ")
	fmt.Scanln(&input)

	if len(input) > 1000 {
		fmt.Println("Ошибка! Длина строки не должна превышать 1000 символов.")
		return
	}

	for _, char := range input {

		if !unicode.IsLetter(char) || (char < 'A' || char > 'z') {
			fmt.Println("Ошибка! Вводите только английские буквы.")
			return
		}
	}

	letters := strings.Split(input, "")
	result := strings.Join(letters, "*")
	fmt.Println(result)
}
