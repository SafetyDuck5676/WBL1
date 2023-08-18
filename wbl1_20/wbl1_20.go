package main

import (
	"fmt"
	"strings"
)

// Функция, которая переворачивает слова в строке
func reverseWords(sentence string) string {
	// Разделяем строку на слова
	words := strings.Split(sentence, " ")

	// Создаем новый слайс для перевернутых слов
	reversedWords := make([]string, 0, len(words))

	// Перебираем слова в обратном порядке
	for i := len(words) - 1; i >= 0; i-- {
		// Добавляем каждое перевернутое слово в новый слайс
		reversedWords = append(reversedWords, reverseString(words[i]))
	}

	// Соединяем перевернутые слова в одну строку
	reversedSentence := strings.Join(reversedWords, " ")

	// Возвращаем перевернутую строку
	return reversedSentence
}

// Функция, которая переворачивает строку
func reverseString(s string) string {
	// Преобразуем строку в слайс символов
	runes := []rune(s)

	// Перебираем символы в слайсе и меняем их местами
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Преобразуем слайс символов обратно в строку
	return string(runes)
}

func main() {
	// Тестовая строка
	sentence := "snow dog sun — sun dog snow"

	// Вызываем функцию для переворачивания слов
	reversed := reverseWords(sentence)

	// Выводим результат
	fmt.Println(reversed)
}
