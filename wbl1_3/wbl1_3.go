package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создаем слайс с последовательностью чисел
	numbers := []int{2, 4, 6, 8, 10}

	// Создаем канал, чтобы получить результаты суммы квадратов чисел
	resultChan := make(chan int)

	// Создаем группу ожидания для синхронизации горутин
	var wg sync.WaitGroup
	wg.Add(len(numbers))

	// Итерируемся по числам и запускаем горутину для каждого числа
	for _, num := range numbers {
		go func(n int) {
			// Вычисляем квадрат числа
			square := n * n

			// Отправляем результат в канал
			resultChan <- square

			// Сообщаем группе ожидания о завершении работы горутины
			wg.Done()
		}(num)
	}

	// В главной горутине ожидаем завершения всех горутин
	go func() {
		wg.Wait()

		// Закрываем канал после получения всех результатов
		close(resultChan)
	}()

	// Суммируем результаты из канала
	sum := 0
	for square := range resultChan {
		sum += square
	}

	// Выводим результат
	fmt.Println("Сумма квадратов:", sum)
}
