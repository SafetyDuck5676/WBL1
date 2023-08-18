package main

import (
	"fmt"
	"math"
)

func main() {
	// Исходные данные
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Создаем карту для хранения групп с шагом в 10 градусов
	temperatureGroups := make(map[float64][]float64)

	// Обходим все температуры
	for _, temp := range temperatures {
		// Округляем значение температуры до ближайшего меньшего числа, кратного 10
		groupKey := math.Floor(temp/10) * 10

		// Добавляем температуру в соответствующую группу
		temperatureGroups[groupKey] = append(temperatureGroups[groupKey], temp)
	}

	// Выводим результат
	for groupKey, groupTemps := range temperatureGroups {
		fmt.Printf("Группа %g: %v\n", groupKey, groupTemps)
	}
}
