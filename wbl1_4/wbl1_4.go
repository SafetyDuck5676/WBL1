package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	// Создаем канал для передачи данных
	dataChannel := make(chan string)

	// Вводим количество воркеров
	numWorkers := 5

	// Создаем контекст и отменяющую функцию для остановки всех воркеров
	ctx, cancel := context.WithCancel(context.Background())

	// Запускаем горутину для записи данных в канал
	go writeToChannel(dataChannel)

	// Создаем sync.WaitGroup для ожидания завершения работы всех воркеров
	var wg sync.WaitGroup

	// Запускаем воркеры
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, dataChannel, &wg, ctx)
	}

	// Ждем сигналов Ctrl+C
	waitForCtrlC(cancel)

	// Ждем завершения работы всех воркеров
	wg.Wait()

	// Закрываем канал
	close(dataChannel)

	fmt.Println("Все воркеры завершили работу")
}

func writeToChannel(ch chan<- string) {
	// В данном примере просто записываем произвольные данные в канал
	for i := 1; i <= 10; i++ {
		ch <- fmt.Sprintf("Данные %d", i)
	}
}

func worker(id int, ch <-chan string, wg *sync.WaitGroup, ctx context.Context) {
	// Передаем WaitGroup, чтобы сообщить, что воркер завершил работу
	defer wg.Done()

	// Читаем данные из канала и выводим на stdout
	for {
		select {
		case data := <-ch:
			fmt.Printf("Воркер %d получил данные: %s\n", id, data)
		case <-ctx.Done():
			fmt.Printf("Воркер %d завершил работу\n", id)
			return
		}
	}
}

func waitForCtrlC(cancel context.CancelFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cancel()
	}()
}
