package main

import (
	"fmt"
	"strings"
)

// Функция, которая переворачивает порядок слов в строке
func reverseWords(sentence string) string {
	// Разделяем строку на слова
	words := strings.Split(sentence, " ")

	// Создаем новый слайс для перевернутых слов
	reversedWords := make([]string, 0, len(words))

	// Перебираем слова в обратном порядке
	for i := len(words) - 1; i >= 0; i-- {
		// Добавляем каждое слово в новый слайс
		reversedWords = append(reversedWords, words[i])
	}

	// Соединяем перевернутые слова в одну строку
	reversedSentence := strings.Join(reversedWords, " ")

	// Возвращаем перевернутую строку
	return reversedSentence
}

func main() {
	// Тестовая строка
	sentence := "snow dog sun"

	// Вызываем функцию для переворачивания порядка слов
	reversed := reverseWords(sentence)

	// Выводим результат
	fmt.Println(reversed)
}
