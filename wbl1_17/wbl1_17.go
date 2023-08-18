package main

import (
	"fmt"
	"sort"
)

func main() {
	// Упорядоченный список чисел
	numbers := []int{1, 3, 5, 7, 9, 11, 13}

	// Искомое число
	target := 7

	// Используем функцию sort.Search для выполнения бинарного поиска
	// В функцию передаем длину списка и функцию-компаратор, которая сравнивает элемент из списка с целью
	index := sort.Search(len(numbers), func(i int) bool {
		return numbers[i] >= target
	})

	if index < len(numbers) && numbers[index] == target {
		// Число найдено
		fmt.Printf("Число %v найдено на позиции %v\n", target, index)
	} else {
		// Число не найдено
		fmt.Printf("Число %v не найдено\n", target)
	}
}
