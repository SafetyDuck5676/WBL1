package main

import "fmt"

func quicksort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	pivot := arr[len(arr)-1] // Выбираем опорный элемент в качестве последнего элемента массива
	left := 0                // Индекс левой границы
	right := len(arr) - 2    // Индекс правой границы

	for left <= right {
		for arr[left] < pivot {
			left++ // Увеличиваем индекс левой границы, пока элементы меньше опорного
		}
		for right >= left && arr[right] > pivot {
			right-- // Уменьшаем индекс правой границы, пока элементы больше опорного
		}
		if right < left {
			break // Если индексы пересеклись, завершаем цикл
		}
		arr[left], arr[right] = arr[right], arr[left] // Меняем элементы местами
		left++
		right--
	}

	arr[left], arr[len(arr)-1] = arr[len(arr)-1], arr[left] // Перемещаем опорный элемент в правильную позицию

	quicksort(arr[:left])   // Рекурсивно сортируем элементы до опорного элемента
	quicksort(arr[left+1:]) // Рекурсивно сортируем элементы после опорного элемента
}

func main() {
	arr := []int{9, 4, 2, 7, 1, 5, 3, 6, 8} // Тестируемый массив

	fmt.Println("Массив до сортировки:", arr)
	quicksort(arr)
	fmt.Println("Массив после сортировки:", arr)
}
