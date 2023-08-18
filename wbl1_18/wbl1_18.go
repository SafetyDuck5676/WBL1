package main

import (
	"fmt"
	"sync"
)

// Counter - структура счетчика
type Counter struct {
	count int
	mux   sync.Mutex
}

// Increment - метод для инкрементирования счетчика
func (c *Counter) Increment() {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.count++
}

// GetValue - метод для получения текущего значения счетчика
func (c *Counter) GetValue() int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.count
}

func main() {
	// Создаем экземпляр счетчика
	counter := Counter{count: 0}

	// Создаем WaitGroup для контроля горутин
	var wg sync.WaitGroup

	// Определяем количество горутин, которые будут инкрементировать счетчик
	numGoroutines := 10

	// Увеличиваем счетчик WaitGroup на количество горутин
	wg.Add(numGoroutines)

	// Запускаем горутины
	for i := 0; i < numGoroutines; i++ {
		go func() {
			// Инкрементируем счетчик
			counter.Increment()

			// Горутина завершила работу, уменьшаем счетчик WaitGroup
			wg.Done()
		}()
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	// Выводим итоговое значение счетчика
	fmt.Println("Итоговое значение счетчика:", counter.GetValue())
}
