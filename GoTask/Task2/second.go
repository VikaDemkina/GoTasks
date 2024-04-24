package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Println("Введите длины сторон вашего треугольника, следуя условию a<b<c:")
	fmt.Print("a: ")
	fmt.Scan(&a)
	fmt.Print("b: ")
	fmt.Scan(&b)
	fmt.Print("c: ")
	fmt.Scan(&c)
	if a > b && b > c {
		fmt.Println("Не соответствует поставленному условию!")
		return
	}

	if c*c == a*a+b*b {
		fmt.Println("Треугольник прямоугольный!")
	} else {
		fmt.Println("Треугольник непрямоугольный!")
	}
}
