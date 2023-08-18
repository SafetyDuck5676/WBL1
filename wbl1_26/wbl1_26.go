package main

import (
	"fmt"
	"strings"
)

// Функция проверки уникальности символов в строке
func isUnique(str string) bool {
	// Приводим строку к нижнему регистру для регистронезависимой проверки
	str = strings.ToLower(str)

	// Создаем карту для отслеживания уникальных символов
	seen := make(map[rune]bool)

	// Итерируемся по каждому символу в строке
	for _, char := range str {
		// Если символ уже был встречен ранее, то строка не уникальна
		if seen[char] {
			return false
		}
		// Добавляем символ в карту
		seen[char] = true
	}

	// Если прошли по всем символам и не встретили повторений, то строка уникальна
	return true
}

func main() {
	// Примеры проверки уникальности строк
	str1 := "Hello, World!"
	str2 := "GoLang"
	str3 := "gophEr"
	str4 := "It is a sunny day."
	str5 := ""

	fmt.Printf("Уникальность строки \"%s\": %v\n", str1, isUnique(str1))
	fmt.Printf("Уникальность строки \"%s\": %v\n", str2, isUnique(str2))
	fmt.Printf("Уникальность строки \"%s\": %v\n", str3, isUnique(str3))
	fmt.Printf("Уникальность строки \"%s\": %v\n", str4, isUnique(str4))
	fmt.Printf("Уникальность строки \"%s\": %v\n", str5, isUnique(str5))
}
