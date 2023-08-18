package main

import (
	"fmt"
)

func main() {
	// Создаем каналы для передачи данных
	in := make(chan int)  // Канал для входных чисел
	out := make(chan int) // Канал для результатов операции x*2

	// Горутина для чтения чисел из массива и записи их в канал in
	go func() {
		numbers := []int{1, 2, 3, 4, 5} // Входные числа из массива
		for _, number := range numbers {
			in <- number
		}
		close(in) // Закрываем канал in после завершения передачи чисел
	}()

	// Горутина для умножения чисел на 2 и записи результатов в канал out
	go func() {
		for number := range in {
			result := number * 2
			out <- result
		}
		close(out) // Закрываем канал out после завершения умножения
	}()

	// Чтение результатов из канала out и вывод их в stdout
	for result := range out {
		fmt.Println(result)
	}
}
