package main

import (
	"fmt"
	"reflect"
)

func main() {
	val := "Hello World!" // Здесь можно изменить значение переменной val на любое другое значение

	// Используем рефлексию для определения типа переменной
	t := reflect.TypeOf(val)

	// Проверяем тип переменной
	switch t.Kind() {
	case reflect.Int:
		fmt.Println("Тип: int")
	case reflect.String:
		fmt.Println("Тип: string")
	case reflect.Bool:
		fmt.Println("Тип: bool")
	case reflect.Chan: // Пример определения типа channel
		fmt.Println("Тип: chan")
	default:
		fmt.Println("Неизвестный тип")
	}

	// Вывод полного типа переменной
	fmt.Println("Полный тип:", t)
}
