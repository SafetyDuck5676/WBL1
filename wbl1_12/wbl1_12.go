package main

import (
	"fmt"
)

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]bool) // Создаем пустое множество

	for _, str := range sequence {
		set[str] = true // Добавляем каждую строку в множество
	}

	// Выводим все элементы множества
	for key := range set {
		fmt.Println(key)
	}
}
