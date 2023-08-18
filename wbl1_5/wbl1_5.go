package main

import (
	"fmt"
	"time"
)

func main() {
	N := 5 // Количество секунд
	ch := make(chan int)

	// Функция для отправки значений в канал
	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
			time.Sleep(time.Second) // Подождать 1 секунду перед отправкой следующего значения
		}
		close(ch) // Закрыть канал после отправки всех значений
	}()

	// Функция для чтения значений из канала
	go func() {
		for num := range ch {
			fmt.Println("Принято значение:", num)
		}
	}()

	time.Sleep(time.Duration(N) * time.Second) // Подождать N секунд

	fmt.Println("Программа завершена")
}
