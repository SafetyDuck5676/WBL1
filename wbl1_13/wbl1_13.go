package main

import "fmt"

func main() {
	// Объявляем и инициализируем переменные a и b
	a := 5
	b := 10

	fmt.Println("До замены:")
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	// Замена значений a и b без использования временной переменной
	a = a + b
	b = a - b
	a = a - b

	fmt.Println("После замены:")
	fmt.Println("a =", a)
	fmt.Println("b =", b)
}
