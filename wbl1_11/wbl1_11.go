package main

import "fmt"

func intersection(set1, set2 []int) []int {
	set1Map := make(map[int]bool)
	intersectionSet := []int{}

	// Заполняем map для первого множества
	for _, val := range set1 {
		set1Map[val] = true
	}

	// Проверяем каждый элемент второго множества на наличие в первом множестве
	for _, val := range set2 {
		if set1Map[val] {
			intersectionSet = append(intersectionSet, val)
		}
	}

	return intersectionSet
}

func main() {
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{4, 5, 6, 7, 8}
	result := intersection(set1, set2)
	fmt.Println("Пересечение:", result)
}
