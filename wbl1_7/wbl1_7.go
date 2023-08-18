package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создаем sync.Map для безопасного доступа к данным
	dataMap := sync.Map{}

	// Создаем wait group, чтобы дождаться завершения всех горутин
	wg := sync.WaitGroup{}

	// Запускаем горутины для записи данных
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			// Генерируем уникальный ключ для записи
			key := fmt.Sprintf("key%d", index)

			// Создаем новый объект Data и записываем его в sync.Map
			data := Data{value: index}
			dataMap.Store(key, data)

			fmt.Printf("[%d] Записано значение %d для ключа %s\n", index, index, key)

			wg.Done()
		}(i)
	}

	// Ждем завершения всех горутин
	wg.Wait()

	// Выводим результаты
	fmt.Println("\nВывод данных из sync.Map:")
	dataMap.Range(func(key, value interface{}) bool {
		fmt.Printf("Ключ: %s, Значение: %v\n", key, value)
		return true
	})
}

type Data struct {
	value int
}
