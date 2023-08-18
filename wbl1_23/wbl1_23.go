package main

import "fmt"

// Функция для удаления i-го элемента из слайса
func removeElement(slice []int, index int) []int {
	// Проверка на корректность индекса
	if index < 0 || index >= len(slice) {
		// Индекс вне допустимого диапазона, возвращаем исходный слайс
		return slice
	}

	// Копируем часть слайса после удаляемого элемента на место этого элемента
	copy(slice[index:], slice[index+1:])
	// Удаляем последний элемент
	slice = slice[:len(slice)-1]

	// Возвращаем измененный слайс
	return slice
}

func main() {
	// Создаем слайс с элементами
	numbers := []int{1, 2, 3, 4, 5}
	// Индекс элемента, который нужно удалить
	index := 2

	// Выводим исходный слайс
	fmt.Println("Исходный слайс:", numbers)

	// Удаляем элемент с заданным индексом
	numbers = removeElement(numbers, index)

	// Выводим измененный слайс
	fmt.Println("Измененный слайс:", numbers)
}
