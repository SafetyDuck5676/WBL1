package main

import (
	"fmt"
	"time"
)

// sleep - функция, которая приостанавливает выполнение программы на заданное количество секунд.
func sleep(seconds int) {
	// time.Sleep принимает аргумент типа time.Duration,
	// поэтому нужно преобразовать секунды в длительность.
	duration := time.Duration(seconds) * time.Second

	time.Sleep(duration) // Приостанавливаем выполнение программы на указанное количество секунд.
}

func main() {
	fmt.Println("Начало программы")

	sleep(3) // Приостанавливаем выполнение программы на 3 секунды.

	fmt.Println("Продолжение программы")
}
