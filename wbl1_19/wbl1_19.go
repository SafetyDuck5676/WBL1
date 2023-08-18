package main

import (
	"fmt"
	"unicode/utf8"
)

// Функция, которая переворачивает строку
func reverseString(s string) string {
	// Создаем слайс руин, чтобы хранить руны в обратном порядке
	runes := make([]rune, utf8.RuneCountInString(s))

	// Заполняем слайс рунами из оригинальной строки в обратном порядке
	n := len(runes)
	for _, r := range s {
		n--
		runes[n] = r
	}

	// Возвращаем перевернутую строку, преобразуя слайс рун обратно в строку
	return string(runes)
}

func main() {
	// Тестовая строка
	input := "главрыба"

	// Вызываем функцию для переворота строки
	reversed := reverseString(input)

	// Выводим результат
	fmt.Println(reversed)
}
