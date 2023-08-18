package main

import (
	"fmt"
)

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) int {
	return a / b
}

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func main() {
	a := 2_097_152 // Значение a должно быть больше 2^20
	b := 1_048_576 // Значение b должно быть больше 2^20

	// Выполняем перемножение чисел a и b
	mul := multiply(a, b)
	fmt.Println("Результат перемножения:", mul)

	// Выполняем деление числа a на число b
	div := divide(a, b)
	fmt.Println("Результат деления:", div)

	// Выполняем сложение чисел a и b
	sum := add(a, b)
	fmt.Println("Результат сложения:", sum)

	// Выполняем вычитание числа b из числа a
	sub := subtract(a, b)
	fmt.Println("Результат вычитания:", sub)
}
